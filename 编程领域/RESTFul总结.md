# RESTFul总结

参考文章

1. [RESTful Alternatives to DELETE Request Body](https://stackoverflow.com/questions/14323716/restful-alternatives-to-delete-request-body)
    - 使用RESTful 规范思维去设计批量删除接口是不合适的, DELETE请求无法附加请求体.

2. [我所认为的RESTful API最佳实践](https://www.scienjus.com/my-restful-api-best-practices/)

3. [RESTful GET，如果存在大量参数，是否有必要变通一下？](https://www.zhihu.com/question/36706936)
    - 对于复杂search做成post，不见得就违反了Restful。可以理解为，每次search的结果是一个新创建的临时集合。


> 无法很好地处理这种问题, 无法很顺理成章地得出一个合理的解决方案, 其根本原因是**大家的应用都不是在玩REST设计, 只是在实现层面上“看着像REST”而已**. 你不是使用资源进行系统建模, 不是以资源的角度来进行设计, 自然遇到问题不会从资源的角度去考虑, 最后和REST需要的资源第一位的观点冲突, 把自己绕死这种伪REST其实很要不得, 要么你就把REST丢掉, 只留下“URL好看点不错”这样的目标, 要么你就玩纯粹基于资源的设计和实现
