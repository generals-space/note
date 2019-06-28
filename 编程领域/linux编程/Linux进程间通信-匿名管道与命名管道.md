# Linux进程间通信-匿名管道与命名管道

参考文章

[Linux进程间通信之管道(pipe)、命名管道(FIFO)与信号(Signal)](http://www.cnblogs.com/biyeymyhjob/archive/2012/11/03/2751593.html)

[【Linux/OS/Network】匿名管道（pipe）和命名管道（FIFO）](http://blog.csdn.net/SuLiJuan66/article/details/50588885)

示例

`mkpipe.c`

```c
#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>
#define MAXLINE 1024
int main()
{
    int n;
    int fd[2];
    pid_t pid;
    char line[MAXLINE];
    
    //fd是一个拥有两个成员的数组, 使用pipe()函数后会生成读写两个fd
    //它们成对存在, fd[0]是读, fd[1]是写
    
    if(pipe(fd) < 0)
    {
        exit(0);
    }

    if((pid = fork()) < 0)
    {
        perror("hehe, i am exiting...");
        exit(0);
    }
    else if (pid > 0)
    {
        printf("I am the parent process\n");
        //父进程关闭读fd, 而向写fd里面写数据
        close(fd[0]);
        write(fd[1], "\nhello world\n", 14);
    }
    else
    {
        printf("I am the child process\n");
        close(fd[1]);
        n = read(fd[0], line, MAXLINE);
        // STDOUT_FILENO表示的文件描述符应该为1
        write(STDOUT_FILENO, line, n);
    }
    exit(0);
}

```

```
$ gcc -c mkpipe.c 
$ gcc -o mkpipe mkpipe.o
$ ./mkpipe
I am the parent process
I am the child process

hello world
```