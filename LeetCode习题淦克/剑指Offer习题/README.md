# 🗡剑指Offer

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
> 。
> ```
>
> 

