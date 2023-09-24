# SpringBoot「全局统一处理异常封装」导致返回双重值问题

参考文章

1. [spring boot 2.4.5版本「全局统一处理异常封装」导致返回双重值问题解决](https://blog.csdn.net/u010739100/article/details/118071368)

spring boot: 2.7.13

## 问题描述

spring boot 工程中封装一个`ResponseData{status, message, data}`类, 作为接口响应体.

有一个接口返回`ResponseData`, 并将`V1Deployment`放到`data`字段中.

```java
	@RequestMapping(value = "/deployments/{name}", method = RequestMethod.GET)
	@ApiOperation(value = "获取集群详情", notes = "获取集群详情")
	@ResponseBody
	public ResponseData<V1Deployment> detail(@PathVariable String name) {
		ResponseData<V1Deployment> resp = new ResponseData<V1Deployment>();
		try {
			V1Deployment deploy = kubeclient.getAppsV1Api().readNamespacedDeployment(
				name, namespace, null
			);
			resp.setData(deploy);
		} catch (ApiException e) {
			resp.setStatus(1);
			resp.setMessage(e.getMessage());
		}
		return resp;
	}
```

但在获取某些 deployment 资源时, 会返回一些很奇怪的结果.

```json
{
    "status": 0,
    "message": "",
    "data": {V1Deployment 对象}
}{
    "status": 3,
    "message": "未知异常",
    "data": "Could not write JSON: Not an integer; nested exception is com.fasterxml.jackson.databind.JsonMappingException: Not an integer (strategy.rollingUpdate.maxSurge)"
}
```

没错, 返回了2个`ResponseData`对象, 看着像是 middleware/filter 出现异常后没有直接返回, 而是继续执行.

## 解决方案

参考文章1的排查思路和解决方案写的很详细, 跟我遇到的场景一致, 我也的确用到了全局统一处理异常的手段, 值得一看.

### 问题定位思路

1. 经过测试分析，在controller层正常返回值时，没有异常，只有在返回的值，进入response组装流程抛异常时，才会出现问题，初步断定，当数据进入response组装阶段之后，已经产生的数据自动会缓存进入响应体的body中了
2. 在组装途中，由于json为空值，jackson抛出了exception，被「全局异常统一处理代码」捕获，这时全局异常处理的代码再次对返回体做了报错的封装，将封装结果追加缓存到body中
3. 最后，系统将缓存到body中的数据全部读出来，一起返回给了前端，于是就包含了两次response的数据，这大概就是问题出现的原因所在

解决方法

1. 最简单的方案就是回退spring boot版本到上一个稳定版2.3.1.RELEASE
2. 但我就想用最新版怎么办？想办法，在异常处理，封装完成之后，先将body里面的缓存数据清空，再写入新的response数据即可


```java
    @ExceptionHandler(Exception.class)
    public ResponseData<String> handleException(HttpServletRequest req, Exception e, HttpServletResponse resp) {
        // 这里在捕获异常后, 把遗留的解析内容先清空再返回就可以了.
        try {
            resp.resetBuffer();
        }catch (Exception ex){
            e.printStackTrace();
        }
        return new ResponseData<>(3, "未知异常", e.getLocalizedMessage());
    }
```
