参考文章

1. [Set environment variables Delve](https://stackoverflow.com/questions/50872412/set-environment-variables-delve)

```
AWS_ENV=development AWS_REGION=eu-west-1 dlv debug main.go
```

在执行 dlv 之前, 使用 export 设置环境变量也可以.
