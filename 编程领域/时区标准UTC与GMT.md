# 时区标准UTC与GMT

参考文章

1. [UTC和GMT什么关系？](https://www.zhihu.com/question/27052407)

UTC是根据原子钟来计算时间(世界上最精确的原子钟50亿年才会误差1秒), 是时间的计算依据.

而GMT是`Greenwich Mean Time`, 是划分时区的基准. GMT（格林威治时间）、CST（北京时间）、PST（太平洋时间）等等是具体的时区.

GMT: UTC +0    =    GMT: GMT +0
CST: UTC +8    =    CST: GMT +8
PST: UTC -8    =    PST: GMT -8

即, 其他时区都是以GMT为基准`+n`或`-n`来表示的.