# LocalDateTime

```java
import java.time.LocalDateTime;
import java.time.ZoneOffset;
```

```java
		LocalDateTime now = LocalDateTime.now();
		System.out.println(now); // 2020-07-23T19:57:18.599
		System.out.println(now.toLocalDate()); // 2020-07-23
		System.out.println(now.toLocalTime()); // 19:57:18.599

		System.out.println(now.getHour()); // 20
		System.out.println(now.getSecond()); // 18

		// 得到时间戳, 单位为秒(js和golang都是精确到毫秒的, java这也太low了)
		// ...原来时间戳也是分时区的???
		System.out.println(now.toEpochSecond(ZoneOffset.of("+0"))); // 1595535307
		// 这个才更符合我们常用的时间戳, 本地时区.
		System.out.println(now.toEpochSecond(ZoneOffset.of("+8"))); // 1595506507
```
