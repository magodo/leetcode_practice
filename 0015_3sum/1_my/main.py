#!/usr/bin/env python3
# -*- coding: utf-8 -*-

#########################################################################
# Author: Zhaoting Weng
# Created Time: Sat 10 Aug 2019 03:41:48 PM CST
# Description:
#########################################################################
from typing import List

class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        if len(nums) < 3:
            return []
        if len(nums) == 3:
            if sum(nums) == 0:
                return [nums]
            else:
                return []
        nums.sort()
        output = []
        idx1 = 0
        while idx1 < len(nums) - 2:
            subsum = 0 - nums[idx1]
            idx2, idx3 = idx1+1, len(nums)-1
            while idx2 < idx3:
                if nums[idx2] + nums[idx3] == subsum:
                    output.append([nums[idx1], nums[idx2], nums[idx3]])
                    idx2 += 1
                    idx3 -= 1
                    while idx2 < idx3 and nums[idx2-1] == nums[idx2]:
                        idx2 += 1
                    while idx2 < idx3 and nums[idx3+1] == nums[idx3]:
                        idx3 -= 1
                elif nums[idx2] + nums[idx3] > subsum:
                    idx3 -= 1
                else:
                    idx2 += 1
            idx1 += 1
            while idx1 < len(nums) - 2 and nums[idx1] == nums[idx1-1]:
                idx1 += 1
        return output

if __name__ == '__main__':
    import collections

    s = Solution()
    print(s.threeSum([-1, 0, -1, 1, 2, -4]))
    print(s.threeSum([0,0,0,0]))
