# AT&T汇编与intel汇编语法区别

## 1. 大小写

intel格式的指令使用大写字母，而AT&T格式的使用小写字母。

- intel:	MOV EAX, EBX

- AT&T:	movl %eax, %ebx

## 2. 操作数赋值方向

在intel语法中，第一个表示目的操作数，第二个是源操作数，赋值方向从右到左;

AT&T语法中，第一个表示源操作数，第二个是目的操作数，方向从左到右，合乎自然。

- intel	`MOV EAX, EBX` ; 将EBX寄存器中的值放入EAX中

- AT&T	`movl %EAX, %EBX` ; 将EAX寄存器中的值放入EBX

## 3. 关于前缀

在intel语法中寄存器和立即数不需要前缀;

AT&T中，寄存器需要加前缀'%'，立即数要加前缀'$'。


```
intel	MOV EAX, 1; 将EAX寄存器中的值置为1
AT&T	MOV $1, %EAX; 同上
```

在AT&T汇编中，

符号常数直接引用，不需要加前缀，如：`movl value %ebx;` value为一常数;

在符号常数前加前缀'$'表示引用符号常数的地址，如：`movl $value, %ebx;`是将value的地址放到`ebx`中。

## 4. 关于后缀

AT&T语法中大部分指令操作码的最后一个字母表示操作数大小。

'b'表示byte(一个字节)，'w'表示word(两个字节)，'l'表示long(四个字节)。

intel中处理内存操作数时也有类似语法如：`BYTE PTR`，`WORD PTR`，`DWORD PTR`。

```
intel	MOV AL, BL	MOV EAX, DWORD PTR [EBX]
AT&T	movb %bl,%al	movl (%ebx), %eax
```

## 5. 间接寻址

intel中基地址使用中括号'[]'，AT&T中使用小括号'()'。

另外处理复杂操作数的语法也有所不同：

intel中为`section:[base + index * scale + disp]`，

在AT&T中等价语法为`section:disp(base, index, sale)`。

其中`section`，`index`，`scale`和`disp`(偏移地址)都是可选的，在指定`index`而没有显式指定scale的情况下使用默认值1，scale和disp不需要加前缀'&'。

```
movl -4(%ebp),%eax;把(ebp - 4)中的内容装入eax
```