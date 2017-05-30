/*************************************************************************
 Author: Zhaoting Weng
 Created Time: Sun 05 Mar 2017 08:22:18 PM CST
 Description: 
 ************************************************************************/

#include <algorithm>
#include <vector>
#include <iostream>

using namespace std;

class Solution {
public:
    int maxArea(vector<int>& height) {
        int i = 0, j = height.size()-1;
        int h, capacity = 0;
        while (i < j)
        {
            h = min(height[i], height[j]);
            capacity = max(capacity, h*(j-i));
            while ((height[i] <= h) && (i < j)) ++i;
            while ((height[j] <= h) && (i < j)) --j;
        }
        return capacity;
    }
};

int main()
{
    Solution sol;
    vector<int> v;

    v = vector<int>{1,2};
    cout << sol.maxArea(v) << endl;
}
