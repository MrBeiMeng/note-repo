# 剑指🗡Offer

## 数组和字符串📶

### [剑指 Offer 04. 二维数组中的查找](https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/)

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

### [剑指 Offer 05. 替换空格](https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/)


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
### [剑指 Offer 11. 旋转数组的最小数字](https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/)

![image-20211206103909848](file://C:\Users\11923\AppData\Roaming\Typora\typora-user-images\image-20211206103909848.png?lastModify=1638758748)

> **解题方案**
>
> 思路：二分查找
>
> - 数组是有大小顺序的，从左到右，第一个非递增的数就是答案
> - 二分查找时，如果**nums[mid]**值比**nums[right]**值大，那么answer(**第一个非递增的数**)在mid的右边，left = mid+1;
> - 二分查找时，如果nums[mid]值比nums[right]小，说明mid到right处于第二个递增中，answer 在 mid左边，right=mid;
> - 二分查找时，如果mid 和 right相等right--;
> - 保证left < right
>
> <video src="https://ccurj.oss-cn-beijing.aliyuncs.com/二分查找最小数字.mp4"></video>
>
> ```java
> class Solution {
>     public int minArray(int[] numbers) {
>         int left = 0;
>         int right = numbers.length - 1;
>         while (left < right) {
>             int mid = (right + left) / 2;
>             if (numbers[mid] < numbers[right]) {
>                 right = mid;
>             } else if (numbers[mid] > numbers[right]) {
>                 left = mid + 1;
>             } else {
>                 right --;
>             }
>         }
>         return numbers[left];
>     }
> }
> ```

<hr style="background:#ffd04c;margin: 0 60px">

### [剑指 Offer 17. 打印从1到最大的n位数](https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/)

![image-20211206112930787](https://ccurj.oss-cn-beijing.aliyuncs.com/image-20211206112930787.png)

> 此题的测试集有一些问题，本题的愿意应该是处理一些大数。比如当n=10的情况
>
> 有两种可能溢出的情况
>
> - int最大是2^31-1 = 2,147,483,647 是九位，当n>=10时,数组中就会有些元素大于int上限。
> - java 数组长度受限于int,一旦那个数大于int上限，那么由他组成的数组长度也会超出int上限。
>
> 但此题规定的返回值类型时int,有点自相矛盾。
>
> ```java
> class Solution {
>     public int[] printNumbers(int n) {
>         int num = (int)Math.pow(10,n)-1;
>         int[] answer = new int[num];
>         for(int i=0;i<num;i++)
>             answer[i] = i+1;
>         return answer;
>     }
> }
> ```

<hr style="background:#ffd04c;margin: 0 60px">

### [剑指 Offer 21. 调整数组顺序使奇数位于偶数前面](https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/)

![image-20211206131703503](https://ccurj.oss-cn-beijing.aliyuncs.com/image-20211206131703503.png)

> **解题方案**
>
> 思路 双指针
>
> - 首先指定前指针 `start` 和后指针 `end`，然后前指针定位偶数，后指针定位奇数，定位到之后将两个值互换，直到数组遍历完成
> - 时间复杂度：O(n)，空间复杂度：O(1)
>
> **算法流程**
>
> 1. 初始化前指针 start = 0，后指针 end = nums.length - 1
> 2. 当 start < end 时表示该数组还未遍历完成，则继续进行奇数和偶数的交换
> 3. 当 nums[start] 为奇数时，则 start++，直到找到不为奇数的下标为止
> 4. 当 nums[end] 为偶数时，则 end--，直到找到不为偶数的下标为止
> 5. 交换 nums[start] 和 nums[end]，继续下一轮交换
> 6. 返回 nums，即为交换后的结果
>
> ```java
> class Solution {
>     public int[] exchange(int[] nums) {
>         int start = 0;
>         int end = nums.length - 1;
>         while(start < end) {
>             while(start < end && (nums[start] % 2) == 1) {
>                 start++;
>             }
>             while(start < end && (nums[end] % 2) == 0) {
>                 end--;
>             }
>             int tmp = nums[start];
>             nums[start] = nums[end];
>             nums[end] = tmp;
>         }
>         return nums;
>     }
> }
> ```

<hr style="background:#ffd04c;margin: 0 60px">

