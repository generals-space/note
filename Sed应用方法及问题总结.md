# Sed应用方法及问题总结

## FAQ

### 1. 使用sed时出现问题

```
general@ubuntu:/tmp$ sed -n '0,3p' ./result 
sed: -e expression #1, char 4: invalid usage of line address 0
...
general@ubuntu:/tmp$ cat ./result | sed -n '0,3p'
sed: -e expression #1, char 4: invalid usage of line address 0
```

出现原因:

...sed的行号是从1开始的, 所以它不明白第0行在哪里.

解决办法: 

把行号从0改成1就行了.