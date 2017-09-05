/*************************************************************************
 Author: Zhaoting Weng
 Created Time: Tue 05 Sep 2017 09:56:54 AM CST
 Description: 
 ************************************************************************/

#include <vector>
#include <string>
#include <algorithm>

using std::vector;
using std::string;
using std::max;

class Solution
{
    public:
        int findMaxForm(const vector<string>& strs, int n_zero, int n_one)
        {
            /* Construct 2-D matrix (n_zero+1 * n_one+1). Each one holds the maximum amount
             * of strings could form. The initial value of them are 0. */

            vector<vector<int>> score(n_zero+1, vector<int>(n_one+1, 0));

            for (const auto str: strs)
            {
                int zeros{0};
                int ones{0};

                for (const auto c: str)
                    switch (c)
                    {
                        case '0':
                            ++zeros;
                            break;
                        case '1':
                            ++ones;
                            break;
                    }

                /* Now "zeros" and "ones" mean the amount of 0 and 1 in one string.
                 * We update the matrix of score matrix in each iteration, and we 
                 * update them from (n_zero, n_one) to (0,0). */

                for (int i = n_zero; i >= zeros; --i)
                    for (int j = n_one; j >= ones; --j)
                        score[i][j] = max(score[i][j], score[i-zeros][j-ones]+1);
            }

            return score[n_zero][n_one];
        }
};

#include <iostream>

int main()
{
    Solution sol;
    std::cout << sol.findMaxForm({"1001","1", "0"}, 3,3) << std::endl;
}
