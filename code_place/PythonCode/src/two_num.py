from typing import List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        tmp_dir = {}

        for i in range(len(nums)):
            tmp_target = target - nums[i]

            if tmp_target in tmp_dir.keys():
                return [tmp_dir.get(tmp_target), i]

            tmp_dir[nums[i]] = i

        return []


if __name__ == '__main__':
    s = Solution()
    print(s.twoSum([2, 7, 11, 15], 9))
    print(s.twoSum([3, 2, 4], 6))
    print(s.twoSum([3, 3], 6))
