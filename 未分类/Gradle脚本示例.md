# Gradle脚本示例

```groovy
/*
*	@name:downloadJars
*	@function:将dependencies中的依赖包下载到当前目录的libs目录下
*			  见task downloadJars
* */

apply plugin: 'java'

repositories{
        mavenLocal()
        //maven字段必须指定url, 不能以mavenLocal()之类替代
        maven{
                name "oschina"
            url "http://maven.oschina.net/content/groups/public/"
        }
        mavenCentral()
}

dependencies{
		//将下载spring的编译时依赖
        compile "org.springframework:spring-context:4.2.2.RELEASE"
}

task downloadJars(type: Copy){
        from configurations.runtime
        //将下载到当前目录的libs目录下
        into 'libs'		
}
```