# gradle引入guava[com.google.common.collect]

参考文章

1. [google相关包引入报错com.google.common](https://www.jianshu.com/p/e147cf16cd79)
2. [Google guava工具类的介绍和使用](https://juejin.im/post/6844903667498221581)
3. [Visual Studio Code - Java - Import Errors and More](https://stackoverflow.com/questions/45743779/visual-studio-code-java-import-errors-and-more)

使用 vscode 创建的 gradle 工程的`build.gradle`配置文件默认如下.

```groovy
plugins {
	id 'org.springframework.boot' version '2.3.3.RELEASE'
	id 'io.spring.dependency-management' version '1.0.10.RELEASE'
	id 'java'
}

group = 'com.example'
version = '0.0.1-SNAPSHOT'
sourceCompatibility = '8'

repositories {
	maven { url 'http://maven.aliyun.com/nexus/content/groups/public'}
	// mavenCentral()
}

dependencies {
	implementation 'org.springframework.boot:spring-boot-starter-web'
	testImplementation('org.springframework.boot:spring-boot-starter-test') {
		exclude group: 'org.junit.vintage', module: 'junit-vintage-engine'
	}
}

test {
	useJUnitPlatform()
}

```

在`dependencies`中添加如下依赖.

```groovy
dependencies {
	implementation 'org.springframework.boot:spring-boot-starter-web'
	testImplementation('org.springframework.boot:spring-boot-starter-test') {
		exclude group: 'org.junit.vintage', module: 'junit-vintage-engine'
	}
	compile 'com.google.guava:guava:26.0-jre'
}
```

> group: 'com.google.guava', name: 'guava', version: '28.0-jre'

然后执行`gradle dependencies`就可重新下载依赖.

------

添加了新的依赖, 但是编译器里`import com.google.common`部分仍然有红色下划波浪线, 提示`The import com.google cannot be resolved`.

我试了试, `gradle build`可以成功编译, vscode 也可以`Run/Debug`, 但是有这个下划线就没法进行代码提示了, 所以很不爽.

在找解决方法时, 发现了参考文章3.

采纳答案说在 vscode 界面`ctrl+shift+p`, 执行`Java Clean`清理 workspace 空间, 这会导致当前窗口重启, 但是有效.

现在不报错了, 也有了代码补全, good.
