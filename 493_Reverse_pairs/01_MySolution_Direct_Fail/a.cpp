/*************************************************************************
 Author: Zhaoting Weng
 Created Time: Tue 28 Feb 2017 08:32:31 PM CST
 Description: 
 ************************************************************************/

#include <vector>
#include <cassert>
#include <climits>
#include <iostream>

using namespace std;

class Solution {
public:
    int reversePairs(vector<int>& nums) {
        int cnt = 0;
        unsigned int inner_index, outer_index, nums_sz;
        int half_int_max = INT_MAX/2;
        int half_int_min = INT_MIN/2;

        nums_sz = nums.size();

        if (nums_sz > 0)
        {
#if 0
            for (outer_index= 0; outer_index< nums_sz-1; ++outer_index)
            {
                for (inner_index = outer_index+1; inner_index < nums_sz; ++inner_index)
                {
                    cnt += (nums[outer_index] > 2*static_cast<long>(nums[inner_index]));
                }
            }
#endif
#if 1
            for (outer_index= 0; outer_index< nums_sz-1; ++outer_index)
            {
                for (inner_index = outer_index+1; inner_index < nums_sz; ++inner_index)
                {
                    if ((nums[inner_index] < 0) || (nums[outer_index] > nums[inner_index]))  // this might help...
                    {
                        /**
                         *  Here we divide nums[i] by 2 rather than multiply nums[j] by 2 so that we will not
                         *  endure the overflow issue.
                         *  However, we need to ensure when "nums[i] > nums[j]*2", then "nums[i]/2 > nums[j]"
                         *  and vice versa.
                         *
                         *  For negative numbers:
                         *
                         *      nums[i]            nums[j]         nums[i]/2       nums[j]*2
                         *
                         *      -3                  -1              -1              -2            (1)
                         *      -3                  -2              -1              -4            (2)
                         *
                         *      (1): nums[i]    <   nums[j]*2
                         *           nums[i]/2  ==  nums[j]
                         *
                         *      (2): nums[i]    >   nums[j]*2
                         *           nums[i]/2  >   nums[j]
                         *
                         *      (note: when a value is divided, it might be bigger than it should be for negative)
                         *
                         *      Therefore, no need for any trick on negative values.
                         *
                         *  For positive numbers:
                         *
                         *      nums[i]            nums[j]         nums[i]/2       nums[j]*2
                         *      3                   1               1               2             (1)
                         *      3                   2               1               4             (2)
                         *
                         *      (1): nums[i]    >   nums[j]*2
                         *           nums[i]/2  ==  nums[j]
                         *
                         *      (2): nums[i]    <   nums[j]*2
                         *           nums[i]/2  <   nums[j]
                         *
                         *      (note: when a value is divided, it might be less than it should be for negative)
                         *
                         *      Therefore, when (nums[i]/2 == nums[j]) and (nums[i]%2 != 0), then it means nums[i]
                         *      suffers attenuation by dividing. In this case we should regard nums[i]/2 == nums[j]
                         *      as a hit (i.e. the judge is nums[i]/2 >= nums[j])
                         *
                         */

                        int half = nums[outer_index] / 2;

                        if (nums[outer_index] >= 0)
                        {
                            int remainder = nums[outer_index] % 2;
                            cnt += (remainder)? (half >= nums[inner_index]) : (half > nums[inner_index]);
                        }
                        else
                        {
                            cnt += (half > nums[inner_index]);
                        }
                    }
                }
            }

#endif
        }
        return cnt;
    }
};

int main()
{
    Solution sol;
    
    vector<int> v;

    v = vector<int>{};
    assert(sol.reversePairs(v) == 0);

    v = vector<int>{1,3,2,3,1};
    assert(sol.reversePairs(v) == 2);

    v = vector<int>{2,4,3,5,1};
    assert(sol.reversePairs(v) == 3);

    v = vector<int>{2147483647,2147483647,2147483647,2147483647,2147483647,2147483647};
    assert(sol.reversePairs(v) == 0);

    v = vector<int>{2147483647,2147483647,-2147483647,-2147483647,-2147483647,2147483647};
    assert(sol.reversePairs(v) == 9);
}
