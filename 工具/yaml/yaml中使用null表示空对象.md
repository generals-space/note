# yaml中使用null表示空对象

在使用 kubectl patch 变更指定字段时, 如果目标字段是一个对象, 如

```yaml
        livenessProbe:
          httpGet:
            path: /healthz
            port: 8081
          initialDelaySeconds: 15
          periodSeconds: 20
```

想把`livenessProbe`置为空, 但是直接设置`livenessProbe: {}`是不行的, 无法覆盖.

需要使用

```yaml
        livenessProbe: null
```
