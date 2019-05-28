// Package sync ...
// golang 1.11.1
package sync

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

// Map 结构类似于go中map[interface{}]interface{}, 但ta是并发安全的.
// Map结构比较特殊, 大多数情况应该使用原生的map类型, 额外加锁即可.
// "map+锁"可以确保类型安全(Map结构中存储的类型都是interface{}),
// 而且存取操作更加方便(Map结构取时需要使用类型断言).
//
// 优先使用Map的两个场景:
// 1. 给定键值只写一次, 但是要读很多次时(就像在caches中, 键值只增不减的情况).
// 2. 当多个协程同时操作, 但是读, 写, 修改的目标各不相交时.
// 在这两种情况下, 使用Map类型相比与"map+(读写)锁"能显著减少锁竞争.
//
// Map一经初始化即可使用, 第一次存/取后就禁止再被拷贝.
type Map struct {
	mu sync.Mutex // 用于对dirty map并发操作

	// read contains the portion of the map's contents that are safe for
	// concurrent access (with or without mu held).
	//
	// The read field itself is always safe to load, but must only be stored with
	// mu held.
	//
	// Entries stored in read may be updated concurrently without mu, but updating
	// a previously-expunged entry requires that the entry be copied to the dirty
	// map and unexpunged with mu held.
	// read map 优先读的map, 支持原子操作(不需要得到mu锁即可读)
	read atomic.Value // 调用Store存储readOnly对象

	// dirty contains the portion of the map's contents that require mu to be
	// held. To ensure that the dirty map can be promoted to the read map quickly,
	// it also includes all of the non-expunged entries in the read map.
	//
	// Expunged entries are not stored in the dirty map. An expunged entry in the
	// clean map must be unexpunged and added to the dirty map before a new value
	// can be stored to it.
	//
	// If the dirty map is nil, the next write to the map will initialize it by
	// making a shallow copy of the clean map, omitting stale entries.
	// dirty存储map中非线程安全的内容(需要先获取mu锁才能写)
	dirty map[interface{}]*entry

	// misses 记录从read map读取不到数据, 从而需要加锁读取read map及dirty map的次数.
	// 当misses的值增长到大于等于dirty map的长度时, 就需要将dirty map"提升"为read map.
	// 因为此时拷贝dirty map的开销已经低于miss记录的值, 提升为read map是更好的选择.
	// 提升为read map时, amended成员为false.
	// 下次存储新的键值时会创建一个新的dirty map.
	misses int
}

// readOnly is an immutable struct stored atomically in the Map.read field.
type readOnly struct {
	m map[interface{}]*entry
	// 当misses值大于等于dirty map长度而将dirty map提升为read map时, 新read map的的amended值为false.
	// 当dirty map中包含一个m(即read map)中没有的key时, amended为true
	amended bool
}

// expunged是一个任意指针(或者叫野指针?), 标记entry对象已经从dirty map中删除.
// 在unexpungeLocked()和tryExpungeLocked()使用atomic.CompareAndSwapPointer()函数做过修改,
// 其余的地方都是对`p == expunged`进行判断.
var expunged = unsafe.Pointer(new(interface{}))

// An entry is a slot in the map corresponding to a particular key.
type entry struct {
	// p points to the interface{} value stored for the entry.
	//
	// If p == nil, the entry has been deleted and m.dirty == nil.
	// 如果p == nil 表示当前entry已经被删除, 且m.dirty = nil
	// If p == expunged, the entry has been deleted, m.dirty != nil, and the entry
	// is missing from m.dirty.
	// 如果p == expunged 表示当前entry已经被删除, 且m.dirty != nil,
	// Otherwise, the entry is valid and recorded in m.read.m[key] and, if m.dirty
	// != nil, in m.dirty[key].
	//
	// An entry can be deleted by atomic replacement with nil: when m.dirty is
	// next created, it will atomically replace nil with expunged and leave
	// m.dirty[key] unset.
	//
	// An entry's associated value can be updated by atomic replacement, provided
	// p != expunged. If p == expunged, an entry's associated value can be updated
	// only after first setting m.dirty[key] = e so that lookups using the dirty
	// map find the entry.
	p unsafe.Pointer // *interface{}
}

// newEntry ...
// caller: m.Store()
func newEntry(i interface{}) *entry {
	return &entry{p: unsafe.Pointer(&i)}
}

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (m *Map) Load(key interface{}) (value interface{}, ok bool) {
	read, _ := m.read.Load().(readOnly)
	e, ok := read.m[key]
	if !ok && read.amended {
		m.mu.Lock()
		// Avoid reporting a spurious miss if m.dirty got promoted while we were
		// blocked on m.mu. (If further loads of the same key will not miss, it's
		// not worth copying the dirty map for this key.)
		read, _ = m.read.Load().(readOnly)
		e, ok = read.m[key]
		if !ok && read.amended {
			e, ok = m.dirty[key]
			// Regardless of whether the entry was present, record a miss: this key
			// will take the slow path until the dirty map is promoted to the read
			// map.
			m.missLocked()
		}
		m.mu.Unlock()
	}
	if !ok {
		return nil, false
	}
	return e.load()
}

func (e *entry) load() (value interface{}, ok bool) {
	p := atomic.LoadPointer(&e.p)
	if p == nil || p == expunged {
		return nil, false
	}
	return *(*interface{})(p), true
}

// Store ...
func (m *Map) Store(key, value interface{}) {
	// m.read为atomic.Value类型, 从其中调用Load()方法取出存入的readOnly值.
	read, _ := m.read.Load().(readOnly)
	// 如果read存在这个键, 并且这个entry没有被标记删除, 尝试调用e.tryStore直接写入. 写入成功, 则结束第一次检测.
	// e为entry对象, 如果在这里tryStore()成功, 就不必往下执行了.
	if e, ok := read.m[key]; ok && e.tryStore(&value) {
		return
	}

	// e.tryStore()失败后, 执行如下尝试
	// 获取dirty map锁, 以进行对m.dirty的操作
	m.mu.Lock()
	defer m.mu.Unlock()

	read, _ = m.read.Load().(readOnly)
	if e, ok := read.m[key]; ok {
		// 如果entry之前被标记为删除, 表示在dirty map中没有被实际删除, 仍存在一个非nil的值.
		if e.unexpungeLocked() {
			m.dirty[key] = e
		}
		// 更新read map中的元素值
		e.storeLocked(&value)
	} else if e, ok := m.dirty[key]; ok {
		e.storeLocked(&value)
	} else {
		// read.amended == false表示曾将dirty map拷贝到read map中, 且此时dirty map为空.
		// 但是新增key操作, 要写dirty map. 需要将read map拷贝到dirty map, 再添加新元素.
		if !read.amended {
			// We're adding the first new key to the dirty map.
			// Make sure it is allocated and mark the read-only map as incomplete.
			m.dirtyLocked()
			m.read.Store(readOnly{m: read.m, amended: true})
		}
		m.dirty[key] = newEntry(value)
	}
}

// tryStore 将一个值存入entry对象, 需要这个entry没有被标记删除, 即e.p != expunged
// 如果entry被标记为已删除, 则返回false, 不做任何操作.
// caller: m.Store()
func (e *entry) tryStore(i *interface{}) bool {
	p := atomic.LoadPointer(&e.p)
	if p == expunged {
		return false
	}
	// 使用CAS方法尝试写入值, 但是由于CAS并非并发发安的, 其返回值也有可能是false,
	// 所以这里使用了for循环, 一直尝试写入(所说自旋锁也是这个原理).
	for {
		if atomic.CompareAndSwapPointer(&e.p, p, unsafe.Pointer(i)) {
			return true
		}
		// 如果p指针 == expunged, 表示此key已经被标记为删除状态, 返回false.
		p = atomic.LoadPointer(&e.p)
		if p == expunged {
			return false
		}
	}
}

// unexpungeLocked 确保元素没有被标记为删除.
// 如果这个元素之前被删除了, 它必须在解锁前(m.mu.Unlock())被添加到dirty map上.
func (e *entry) unexpungeLocked() (wasExpunged bool) {
	return atomic.CompareAndSwapPointer(&e.p, expunged, nil)
}

// storeLocked 无条件将值存储到entry中.
// entry必须是未被标记删除的状态.
// caller: m.Store()
func (e *entry) storeLocked(i *interface{}) {
	atomic.StorePointer(&e.p, unsafe.Pointer(i))
}

// Delete 删除key对应的值.
// 采用延迟删除策略, 当read map存在元素时, 将read map中的元素置为nil.
// 只有在提升dirty的时候才清理删除的数, 延迟删除可以避免后续获取删除的元素时候需要加锁.
// 当read map中不存在目标元素, 而dirty map中存在时, 直接删除dirty map中的元素.
func (m *Map) Delete(key interface{}) {
	read, _ := m.read.Load().(readOnly)
	e, ok := read.m[key]
	// 当read map中不存在目标元素, 而dirty map中存在时, 直接删除dirty map中的元素.
	if !ok && read.amended {
		m.mu.Lock()
		// 这里又加锁重新获取了一次read map...是为了保险吗?
		read, _ = m.read.Load().(readOnly)
		e, ok = read.m[key]
		if !ok && read.amended {
			delete(m.dirty, key)
		}
		m.mu.Unlock()
	}
	// 当read map存在元素时, 将元素置为nil
	if ok {
		// 好像不在意返回值...
		e.delete()
	}
}

// delete ...
// caller m.Delete()
func (e *entry) delete() (hadValue bool) {
	for {
		p := atomic.LoadPointer(&e.p)
		if p == nil || p == expunged {
			return false
		}
		if atomic.CompareAndSwapPointer(&e.p, p, nil) {
			return true
		}
	}
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's
// contents: no key will be visited more than once, but if the value for any key
// is stored or deleted concurrently, Range may reflect any mapping for that key
// from any point during the Range call.
//
// Range may be O(N) with the number of elements in the map even if f returns
// false after a constant number of calls.
func (m *Map) Range(f func(key, value interface{}) bool) {
	// We need to be able to iterate over all of the keys that were already
	// present at the start of the call to Range.
	// If read.amended is false, then read.m satisfies that property without
	// requiring us to hold m.mu for a long time.
	read, _ := m.read.Load().(readOnly)
	if read.amended {
		// m.dirty contains keys not in read.m. Fortunately, Range is already O(N)
		// (assuming the caller does not break out early), so a call to Range
		// amortizes an entire copy of the map: we can promote the dirty copy
		// immediately!
		m.mu.Lock()
		read, _ = m.read.Load().(readOnly)
		if read.amended {
			read = readOnly{m: m.dirty}
			m.read.Store(read)
			m.dirty = nil
			m.misses = 0
		}
		m.mu.Unlock()
	}

	for k, e := range read.m {
		v, ok := e.load()
		if !ok {
			continue
		}
		if !f(k, v) {
			break
		}
	}
}

// missLocked 将m.misses加1.
// 当m.misses的值大于等于dirty map的长度时, 将dirty map的内容拷贝到read map中,
// 然后将dirty map清空, miss值置0.
// caller: m.Load()
func (m *Map) missLocked() {
	m.misses++
	if m.misses < len(m.dirty) {
		return
	}
	m.read.Store(readOnly{m: m.dirty})
	m.dirty = nil
	m.misses = 0
}

// dirtyLocked 将成员从read map复制到dirty map.
// 注意: 此时dirty map必须为nil, 否则不操作. 因为此时
// @caller m.Store()
func (m *Map) dirtyLocked() {
	if m.dirty != nil {
		return
	}

	read, _ := m.read.Load().(readOnly)
	m.dirty = make(map[interface{}]*entry, len(read.m))
	for k, e := range read.m {
		if !e.tryExpungeLocked() {
			m.dirty[k] = e
		}
	}
}

func (e *entry) tryExpungeLocked() (isExpunged bool) {
	p := atomic.LoadPointer(&e.p)
	for p == nil {
		if atomic.CompareAndSwapPointer(&e.p, nil, expunged) {
			return true
		}
		p = atomic.LoadPointer(&e.p)
	}
	return p == expunged
}
