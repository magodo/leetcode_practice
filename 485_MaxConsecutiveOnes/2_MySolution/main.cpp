/*************************************************************************
 Author: Zhaoting Weng
 Created Time: Mon 27 Feb 2017 08:11:54 PM CST
 Description: 
 ************************************************************************/

#include <iostream>
#include <vector>
#include <cassert>

using namespace std;

class Solution
{
    public:
        int findMaxConsecutiveOnes(vector<int>& nums)
        {
            int max = 0, sum = 0;
            for (auto i: nums)
            {
                max = (max < (sum = (sum+1)*i))? sum:max;
            }
            return max;
        }
};

int main()
{
    Solution sol;

    vector<int> v{1,1,1,0,0,1,0,1,1,1,1,0,0};
    assert(sol.findMaxConsecutiveOnes(v) == 4);

    v = vector<int>{0};
    assert(sol.findMaxConsecutiveOnes(v) == 0);

    v = vector<int>{1};
    assert(sol.findMaxConsecutiveOnes(v) == 1);
}
