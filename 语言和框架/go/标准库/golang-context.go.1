// Package context defines the Context type, which carries deadlines,
// cancelation signals, and other request-scoped values across API boundaries
// and between processes.
//
// Programs that use Contexts should follow these rules to keep interfaces
// consistent across packages and enable static analysis tools to check context
// propagation:
package context

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"
)

// Context ...
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

var Canceled = errors.New("context canceled")

var DeadlineExceeded error = deadlineExceededError{}

type deadlineExceededError struct{}

func (deadlineExceededError) Error() string   { return "context deadline exceeded" }
func (deadlineExceededError) Timeout() bool   { return true }
func (deadlineExceededError) Temporary() bool { return true }

type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*emptyCtx) Done() <-chan struct{} {
	return nil
}

func (*emptyCtx) Err() error {
	return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {
	return nil
}

func (e *emptyCtx) String() string {
	switch e {
	case background:
		return "context.Background"
	case todo:
		return "context.TODO"
	}
	return "unknown empty Context"
}

var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

// Background ...
func Background() Context {
	return background
}

// TODO ...
func TODO() Context {
	return todo
}

// CancelFunc ...
type CancelFunc func()

// WithCancel ...
func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
	c := newCancelCtx(parent)
	propagateCancel(parent, &c)
	// 这里返回的是&c, 即cancelCtx对象的指针,
	// 因为在定义cancelCtx的方法时, receiver写的是指针类型,
	// 所以是*cancelCtx实现了Context接口而不是cancelCtx.
	return &c, func() { c.cancel(true, Canceled) }
}

// newCancelCtx ...
func newCancelCtx(parent Context) cancelCtx {
	return cancelCtx{Context: parent}
}

// propagateCancel propagate(传播, 繁衍), 将子级ctx添加到父级ctx的children成员map中,
// 以便之后执行 cancel() 时可以实现传导效果.
// caller: WithCancel(), WithDeadline()
func propagateCancel(parent Context, child canceler) {
	if parent.Done() == nil {
		// 如果父级ctx无法被取消...比如Background(),
		// 相当于没有父级ctx, 就没有继续执行的必要了.
		// child 本身就可以作为 parent 存在了.
		return
	}
	if p, ok := parentCancelCtx(parent); ok {
		// 一般情况下会运行到这里
		p.mu.Lock()
		if p.err != nil {
			// 如果父级ctx已经被取消
			child.cancel(false, p.err)
		} else {
			// 将子级ctx添加到父级ctx的children成员map中
			if p.children == nil {
				p.children = make(map[canceler]struct{})
			}
			p.children[child] = struct{}{}
		}
		p.mu.Unlock()
	} else {
		// 如果p是emptyCtx, 即Background()返回的结果, 则会运行到这里.
		go func() {
			select {
			case <-parent.Done():
				child.cancel(false, parent.Err())
			case <-child.Done():
			}
		}()
	}
}

// parentCancelCtx 返回父级context的`*cancelCtx`类型成员.
// 可能是 cancelCtx, 也可能是 timerCtx(但timerCtx中也是有cancelCtx成员的),
// 但不可能是 valueCtx, 因为 valueCtx
// 最初调用时, parent应该是Background()得到的空ctx.
// caller: propagateCancel(), removeChild()
func parentCancelCtx(parent Context) (*cancelCtx, bool) {
	/*
		WithValue的使用场景一般是
		ctx, cancel := context.WithCancel(context.Background())
		// ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
		ctx = context.WithValue(ctx, "key", "val")
		defer cancel()

		即WithValue一般需要与其他WithXXX组合使用, 所以ta的Context一定有值,
		但没有办法直接确定传入的ctx参数是cancelCtx/timerCtx.
		for循环目的就是, 如果parent是valueCtx类型, 将parent重新赋值, 再做一次判断.
	*/
	for {
		switch c := parent.(type) {
		case *cancelCtx:
			return c, true
		case *timerCtx:
			return &c.cancelCtx, true
		case *valueCtx:
			parent = c.Context
		default:
			// 能运行到这里, 说明parent为emptyCtx, 即Background()的返回值.
			return nil, false
		}
	}
}

// removeChild removes a context from its parent.
func removeChild(parent Context, child canceler) {
	p, ok := parentCancelCtx(parent)
	if !ok {
		return
	}
	p.mu.Lock()
	if p.children != nil {
		delete(p.children, child)
	}
	p.mu.Unlock()
}

// *cancelCtx 和 *timerCtx 都实现了canceler接口, 实现该接口的类型都可以被直接cancel
type canceler interface {
	cancel(removeFromParent bool, err error)
	Done() <-chan struct{}
}

// closedchan is a reusable closed channel.
var closedchan = make(chan struct{})

func init() {
	close(closedchan)
}

type cancelCtx struct {
	// 直接把interface{}放到struct中?
	Context 

	// protects following fields
	mu   sync.Mutex
	done chan struct{}
	// created lazily, closed by first cancel call
	// set to nil by the first cancel call
	// canceler是接口类型, 这里应该是把map当set用了, 保证key唯一, 而不在乎val
	children map[canceler]struct{}
	// set to non-nil by the first cancel call
	err      error 
}

func (c *cancelCtx) Done() <-chan struct{} {
	c.mu.Lock()
	if c.done == nil {
		c.done = make(chan struct{})
	}
	d := c.done
	c.mu.Unlock()
	return d
}

func (c *cancelCtx) Err() error {
	c.mu.Lock()
	err := c.err
	c.mu.Unlock()
	return err
}

func (c *cancelCtx) String() string {
	return fmt.Sprintf("%v.WithCancel", c.Context)
}

// cancel 关闭 c.done 通道, 做了如下几件事
// 1. 设置c.err = err, c.children = nil
// 2. 依次遍历c.children，每个child分别cancel
// 3. 如果设置了`removeFromParent`，则将c从其parent的children中删除.
func (c *cancelCtx) cancel(removeFromParent bool, err error) {
	if err == nil {
		panic("context: internal error: missing cancel error")
	}
	c.mu.Lock()
	if c.err != nil {
		c.mu.Unlock()
		return // already canceled
	}
	c.err = err
	if c.done == nil {
		c.done = closedchan
	} else {
		close(c.done)
	}
	for child := range c.children {
		// NOTE: acquiring the child's lock while holding parent's lock.
		child.cancel(false, err)
	}
	c.children = nil
	c.mu.Unlock()

	if removeFromParent {
		removeChild(c.Context, c)
	}
}

// WithDeadline 如果父级ctx的deadline早于参数d,
// @param parent可能是Background()空context
// If the parent's deadline is already earlier than d,
// WithDeadline(parent, d) is semantically equivalent to parent.
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
	if cur, ok := parent.Deadline(); ok && cur.Before(d) {
		// 如果父级ctx的deadline早于参考d指定的时间, 可以直接返回*cancelCtx对象了.
		// 因为同样可以被取消, 而且父级ctx一定比子ctx早结束.
		return WithCancel(parent)
	}

	c := &timerCtx{
		cancelCtx: newCancelCtx(parent),
		deadline:  d,
	}
	propagateCancel(parent, c)

	// 下面与WithCancel()相比, 多出来的步骤: 设置定时器.
	dur := time.Until(d)
	if dur <= 0 {
		// deadline时间已经过了, 直接取消然后返回
		c.cancel(true, DeadlineExceeded)
		return c, func() { c.cancel(true, Canceled) }
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	if c.err == nil {
		// 启动定时器
		c.timer = time.AfterFunc(dur, func() {
			c.cancel(true, DeadlineExceeded)
		})
	}
	return c, func() { c.cancel(true, Canceled) }
}

type timerCtx struct {
	cancelCtx
	timer    *time.Timer // 也需要cancelCtx.mu加锁
	deadline time.Time
}

func (c *timerCtx) Deadline() (deadline time.Time, ok bool) {
	return c.deadline, true
}

func (c *timerCtx) String() string {
	return fmt.Sprintf(
		"%v.WithDeadline(%s [%s])", 
		c.cancelCtx.Context, c.deadline, time.Until(c.deadline),
	)
}

// cancel 先调用父级cancelCtx.cancel()方法, 多出来的操作是, 
// 销毁定时器成员, 释放资源.
func (c *timerCtx) cancel(removeFromParent bool, err error) {
	// 这里调用父级cancelCtx的cancel方法
	c.cancelCtx.cancel(false, err)
	if removeFromParent {
		// Remove this timerCtx from its parent cancelCtx's children.
		removeChild(c.cancelCtx.Context, c)
	}
	// 与cancelCtx相比, 多出来的操作就是销毁定时器成员, 释放资源.
	c.mu.Lock()
	defer c.mu.Unlock()
	// 停止并销毁计时器成员对象
	if c.timer != nil {
		c.timer.Stop()
		c.timer = nil
	}
}

// WithTimeout ...
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	return WithDeadline(parent, time.Now().Add(timeout))
}

// To avoid allocating when assigning to an interface{}, 
// context keys often have concrete type struct{}.
// Alternatively, exported context key variables' static type
// should be a pointer or interface.
// 注意: WithValue需要与其他WithXXX配合使用
func WithValue(parent Context, key, val interface{}) Context {
	if key == nil {
		panic("nil key")
	}
	if !reflect.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}
	return &valueCtx{parent, key, val}
}

// valueCtx 只是多了一对key/val, 和一个`Value()`方法.
type valueCtx struct {
	Context
	key, val interface{}
}

func (c *valueCtx) String() string {
	return fmt.Sprintf("%v.WithValue(%#v, %#v)", c.Context, c.key, c.val)
}

func (c *valueCtx) Value(key interface{}) interface{} {
	if c.key == key {
		return c.val
	}
	return c.Context.Value(key)
}
