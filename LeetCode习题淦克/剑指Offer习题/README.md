# 剑指🗡Offer

## [剑指 Offer 04. 二维数组中的查找](https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/)

![image-20211206101452475](https://ccurj.oss-cn-beijing.aliyuncs.com/image-20211206101452475.png)

> **解题方案：**
>
> 思路：
>
> - 从矩阵的左下角看，上方的数字都比其小，右侧的数字都比他大，所以可以一句此规律去判断数字是否存在![二维数组查找元素](https://ccurj.oss-cn-beijing.aliyuncs.com/%E4%BA%8C%E7%BB%B4%E6%95%B0%E7%BB%84%E6%9F%A5%E6%89%BE%E5%85%83%E7%B4%A0.gif)
>
> ```java
> class Solution {
>     public boolean findNumberIn2DArray(int[][] matrix, int target) {
>         if(matrix.length == 0)
>             return false;
> 
>         int x = 0;
>         int y = matrix.length - 1;
> 
>         while(x < matrix[0].length && y >= 0){
>             if(matrix[y][x] > target) {
>                 y--;
>             } else if(matrix[y][x] < target) {
>                 x++;
>             } else {
>                 return true;
>             }
>         }
> 
>         return false;
>     }
> }
> ```
>
> 

<hr style="background:#ffd04c;margin: 0 60px">
## [剑指 Offer 05. 替换空格](https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/)

![image-20211206102151052](https://ccurj.oss-cn-beijing.aliyuncs.com/image-20211206102151052.png)

> **解题方案**
>
> 思路：
>
> - 增加一个新字符串，遍历原来的字符串，遍历过程中，如果非空格则将原来的字符直接拼接到新字符串中，如果遇到空格则将`%20`拼接到新字符串中
>
> ```java
> class Solution {
>     public String replaceSpace(String s) {
>         StringBuilder sb = new StringBuilder();
>         for(int i = 0 ; i < s.length(); i++){
>             char ch = s.charAt(i);
>             if(ch == ' ') {
>                 sb.append("%20");
>             }
>             else {
>                 sb.append(ch);
>             }
>         }
>         return sb.toString();
>     }
> }
> ```
>
> **其他知识 - String、StringBuilder 和 StringBuffer**
>
> - 可变与不可变
>
>   String类是一个不可变类，如果一个字符串需要经常被改变应该使用其他两个类
>
> - 初始化方式
>
>   ```java
>   String str = new String("java");
>   String str2 = "java";
>   StringBuffer sb = new StringBuffer("hello");
>   ```
>
> - 字符串修改方式
>
>   String字符串修改方法是首先创建一个StringBuffer，其次调用StringBuffer的append方法，最后调用StringBuffer的toString()方法把结果返回,StringBuffer和StringBuilder在修改字符串方面比String的性能要高。
>
> - 是否实现了equals和hashCode方法
>
>   String实现了equals()方法和hashCode()方法，new String("java").equals(new String("java"))的结果为true；而StringBuffer没有实现equals()方法和hashCode()方法，因此，new StringBuffer("java").equals(new StringBuffer("java"))的结果为false，将StringBuffer对象存储进Java集合类中会出现问题。
>
>
> - 是否线程安全
>
>   StringBuilder是线程不安全的，StringBuffer是线程安全的，如果只是在单线程中使用字符串缓冲区，则StringBuilder的效率会高些，但是当多线程访问时，最好使用StringBuffer。
>
> 综上，在执行效率方面，StringBuilder最高，StringBuffer次之，String最低，对于这种情况，一般而言，如果要操作的数量比较小，应优先使用String类；如果是在单线程下操作大量数据，应优先使用StringBuilder类；如果是在多线程下操作大量数据，应优先使用StringBuilder类。

<hr style="background:#ffd04c;margin: 0 60px">

