# 665. Non-decreasing Array

## Problem

### Description

Given an array with n integers, your task is to check if it could become non-decreasing by modifying at most 1 element.

We define an array is non-decreasing if array[i] <= array[i + 1] holds for every i (1 <= i < n).
### Example
```cpp
Input: [4,2,3]
Output: True
Explanation: You could modify the first 4 to 1 to get a non-decreasing array.

Input: [4,2,1]
Output: False
Explanation: You can't get a non-decreasing array by modify at most one element.
```
> Note: 如果array [i] <= array [i + 1]对于每个i（1 <= i <n）成立，我们定义一个数组是非递减的。

>优先把 nums[i]降为nums[i+1]，这样可以减少 nums[i+1] > nums[i+2] 的风险。

## Solution
### Approach1

### Approach2
