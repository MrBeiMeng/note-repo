## 1 两数之和

<span style="font-weight:bold;font-size:14px">难度：</span> <span style="background:#3de1ad;border-radius:5px;padding:1px 5px;font-weight:bold;color:#ffffff">简单</span>  <span style="font-weight:bold;font-size:14px">地址：</span> [两数之和](https://leetcode.cn/problems/two-sum/)

给定一个整数数组 `nums` 和一个整数目标值 `target`，请你在该数组中找出 **和为目标值** `target` 的那 **两个** 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。


<span style="font-size:14px;font-weight:bold">示例1:</span>

<p style="background:#f2f2f4;padding:10px 15px;border-radius:5px">
  <span style="font-size:14px;font-weight:bold">输入：</span>
  <span style="font-size:14px;">nums = [2,7,11,15], target = 9</span>
  <br>
  <span style="font-size:14px;font-weight:bold">输出：</span>
  <span style="font-size:14px;">[0,1]</span>
  <br>
  <span style="font-size:14px;font-weight:bold">解释：</span>
  <span style="font-size:14px;">因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。</span>
</p>

<span style="font-size:14px;font-weight:bold">示例2:</span>

<p style="background:#f2f2f4;padding:10px 15px;border-radius:5px">
  <span style="font-size:14px;font-weight:bold">输入：</span>
  <span style="font-size:14px;">nums = [3,2,4], target = 6</span>
  <br>
  <span style="font-size:14px;font-weight:bold">输出：</span>
  <span style="font-size:14px;">[1,2]</span>
</p>

<span style="font-size:14px;font-weight:bold">示例3:</span>

<p style="background:#f2f2f4;padding:10px 15px;border-radius:5px">
  <span style="font-size:14px;font-weight:bold">输入：</span>
  <span style="font-size:14px;">nums = [3,3], target = 6</span>
  <br>
  <span style="font-size:14px;font-weight:bold">输出：</span>
  <span style="font-size:14px;">[0,1]</span>
</p>


<span style="font-size:14px;font-weight:bold">提示:</span>

<ul>
  <li>2 <= nums.length <= 104</li>
  <li>-109 <= nums[i] <= 109</li>
  <li>-109 <= target <= 109</li>
  <li>只会存在一个有效答案</li>
</ul>
<span style="font-size:14px;font-weight:bold">进阶: </span> <span style="font-size:14px;">你可以想出一个时间复杂度小于 O(n2) 的算法吗？</span>




<div style="text-align:center">
  <span>🔑 解题思路</span>
  <p style="display: none">用一个值去寻找另一个值</p>

[我的题解-java](../code_place/JavaCode/src/main/java/TwoNum.java)
[我的题解-python](../code_place/PythonCode/src/two_num.py)
[我的题解-golang](../code_place/GoCode/two_num.go)
</div>



## 文档工具

<span style="background:#3de1ad;border-radius:5px;padding:1px 5px;font-weight:bold;color:#ffffff">简单</span> 
<span style="background:#ffa400;border-radius:5px;padding:1px 5px;font-weight:bold;color:#ffffff">中等</span> 
<span style="background:#f35336;border-radius:5px;padding:1px 5px;font-weight:bold;color:#ffffff">困难</span>

你好世界