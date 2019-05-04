`gc.set_threshold()`

`gc.set_debug(flags)`
　　- `gc.DEBUG_COLLETABLE`:  打印可以被垃圾回收器回收的对象
　　- `gc.DEBUG_UNCOLLETABLE`:  打印无法被垃圾回收器回收的对象，即定义了__del__的对象
　　- `gc.DEBUG_SAVEALL`: 当设置了这个选项，可以被拉起回收的对象不会被真正销毁（free），而是放到gc.garbage这个列表里面，利于在线上查找问题
