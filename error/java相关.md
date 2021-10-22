# Java相关错误收集



## Unable to make protected final java.lang.Class java.lang.ClassLoader.defineClass(java.lang.String,byte[],int,int,java.security.ProtectionDomain) throws java.lang.ClassFormatError accessible: module java.base does not "opens java.lang" to unnamed module @7ab2bfe1

添加命令: java --add-opens java.base/java.lang=ALL-UNNAMED

异常是由Java 9及以上版本中引入的Java Platform Module System引起的，特别是强封装的实现。
它仅在特定条件下允许access，最突出的条件是:

- 类型必须是公共的
- 必须导出拥有的软件包

对于反射，导致异常的代码尝试使用相同的限制。
更确切地说，异常是由对 `setAccessible` 的调用引起的。

[参考文档-> 👉](https://www.cnblogs.com/stcweb/articles/15114266.html)

