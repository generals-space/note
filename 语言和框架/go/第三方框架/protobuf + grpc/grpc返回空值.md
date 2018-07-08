# grpc返回空值

参考文章

1. [以小见大——那些基于 protobuf 的五花八门的 RPC（5 完）](https://blog.csdn.net/gzlaiyonghao/article/details/6323900)

2. [Can I define a grpc call with a null request or response?](https://stackoverflow.com/questions/31768665/can-i-define-a-grpc-call-with-a-null-request-or-response/31772973)

> ... we as developers are really bad at guessing what we might want in the future. So I recommend being safe by always defining custom params and results types for every method, even if they are empty.

