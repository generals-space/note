参考文章

1. [Go错误处理：错误链使用指南](https://tonybai.com/2023/05/14/a-guide-of-using-go-error-chain/)
    - "Tony Bai"出品, 必属精品

错误链感觉有点像 python 中的分层异常.

```py
    try:
        pass
    except OSError as e:
        print(e)
    except Exception as e:
        print(e)
```
