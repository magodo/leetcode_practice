/*************************************************************************
 Author: Zhaoting Weng
 Created Time: Thu 02 Mar 2017 10:30:30 AM CST
 File Name: solution.cpp
 Description: 

    The solution is to divide problem to sub-problems:

    T(n,m) = T(n,m-1) + C, where C means how many numbers whose index i in range [n,m-1], nums[i] > 2*nums[m]

    So for the problem, we can deduce into following procedure:

    T(0,n) = T(0,n-1) + Cn = T(0,n-2) + Cn-1 + Cn =... T(0,1) + C2 +  ... Cn = C1 +... Cn

    Each Ci could be calculated by following procedure:

    Let's say we want to get Cn, then we need to check how many numbers whose index i in range [0,n-1],
    nums[i] > 2*nums[n]. We can arrange each number in nums(0,n-1) into a BST, and each node in this tree
    contains a counter recording "sum of count of node in its right sub-tree + amount of this node". 
    So when searching for nums[n] in an existed tree, if an existed node's key:

    * <  nums[n]->key: search the right sub-tree 
    * == nums[n]->key: add count of the node and stop
    * > nums[n]->key: search the left sub-tree and add count of the node

    When the tree is constructed, when insert a node, when it goes through, if an existed node's key:

    * < the key to insert, continue check on the left sub-tree
    * == the key to insert, increase count of the node and stop
    * > the key to insert, increase count of the node and continue check on right sub-tree
    
    We can start from the 1st element in "nums", caculate C1; Then insert 2nd element, calculate C2; And so forth.

 ************************************************************************/

#include <vector>
#include <iostream>
#include <cassert>
#include <chrono>

using namespace std::chrono;
using namespace std;

/***********************
 * Tuner Wrapper
 * *********************/
template<typename F>
class TimerWrapper
{
    public:
        TimerWrapper(F func, clock_t& elapsedTime):
            call(func),
            startTime_(::clock()),
            elapsedTime_(elapsedTime)
        {}

        ~TimerWrapper()
        {
            const clock_t endTime_ = ::clock();
            elapsedTime_ += endTime_ - startTime_;
        }

        F call;
    private:
        const clock_t startTime_;
        clock_t& elapsedTime_;
};

template<typename F>
TimerWrapper<F> getWrapper(F func, clock_t& elapsedTime)
{ return TimerWrapper<F>(func, elapsedTime); }

/***********************
 * Implementation
 * *********************/

class Solution
{
    public:

        struct Node
        {
            Node(int cnt, int key):
                left(nullptr),
                right(nullptr),
                cnt(cnt),
                height(1),
                key(key)
            {}

            Node* left;     // ptr to left node
            Node* right;    // ptr to right node
            int cnt;        // amount of "node on its sub-tree" + "node itself"
            int height;     // height of node
            int key;        

            /**
             * Adjust height of a node, the precondition is the children's heights are adjusted
             */
            void adjust_height()
            {
                height = max((left?left->height:0), (right?right->height:0)) + 1;
            }

            /**
             * balance factor of the node
             */
            int bfactor()
            {
                int lh = (left==nullptr)? 0 :left->height;
                int rh = (right==nullptr)? 0 :right->height;
                return lh-rh;
            }
        };

        int reversePairs(vector<int> &nums)
        {
            int sum = 0;
            Node *root = nullptr;

            /* return 0 if input list is empty */
            if (nums.size() == 0)
                return 0;

            for (auto it = nums.begin(); it < nums.end() -1; ++it)
            {
            //    high_resolution_clock::time_point start = high_resolution_clock::now();
                root = insert(root, *it);
            //   cout << *it << ": " << duration_cast<microseconds>(high_resolution_clock::now() - start).count() << endl;
            //  start = high_resolution_clock::now();
                sum += search(root, static_cast<long>(*(it+1))*2);
            // cout << *it << ": " << duration_cast<microseconds>(high_resolution_clock::now() - start).count() << endl;
            }
            return sum;
        }


        Node *insert(Node *node, long key)
        {
            if (node == nullptr)
                return new Node(1, key);

            if (key < node->key)
            {
                node->left = insert(node->left, key);
                
                // After adding an left child, LL/LR might occur
                if (node->bfactor() == 2)
                {
                    // LL
                    if (key < node->left->key)
                        node = rightRotate(node);
                    // LR
                    else
                        node = leftRightRotate(node);
                }
                else
                {
                    // no need to rotate, just adjust height
                    node->adjust_height();
                }
            }
            else if (key > node->key)
            {
                ++node->cnt;
                node->right = insert(node->right, key);
                
                // After adding an right child, RR/RL might occur
                if (node->bfactor() == -2)
                {
                    // RR
                    if (key > node->right->key)
                        node = leftRotate(node);
                    // RL
                    else
                        node = rightLeftRotate(node);
                }
                else
                {
                    // no need to rotate, just adjust height
                    node->adjust_height();
                }
            }
            else // (key == node->key)
            {
                ++node->cnt;
            }

            return node;
        }

        int search(Node *node, long key)
        {
            if (node == nullptr)
                return 0;
            
            if (key == node->key)
                return search(node->right, key);

            else if (key < node->key)
            {
                return (node->cnt + search(node->left, key));
            }

            else if (key > node->key)
            {
                return search(node->right, key);
            }
        }


        // Right rotate for LL case
        // (this will also update count and height for the nodes affected)
        Node *rightRotate(Node *node)
        {
            Node *pivot = node->left;
            if (pivot != nullptr)
            {
                node->left = pivot->right;
                pivot->right = node;
                
                // adjust count of right sub-tree + itself
                // (only for the node whose right sub-tree changes)
                pivot->cnt += node->cnt;
                
                // adjust height
                node->adjust_height();
                pivot->adjust_height();
            }
            return pivot;
        }

        // Left rotate for RR case
        // (this will also update count and height for the nodes affected)
        Node *leftRotate(Node *node)
        {
            Node *pivot = node->right;
            if (pivot != nullptr)
            {
                node->right = pivot->left;
                pivot->left = node;
                // adjust count of right sub-tree + itself
                // (only for the node whose right sub-tree changes)
                node->cnt -= pivot->cnt;
                
                // adjust hegiht
                node->adjust_height();
                pivot->adjust_height();
            }
            return pivot;
        }

        // for RL case
        Node *rightLeftRotate(Node *node)
        {
            node->right = rightRotate(node->right);
            return leftRotate(node);
        }

        // for LR case
        Node *leftRightRotate(Node *node)
        {
            node->left = leftRotate(node->left);
            return rightRotate(node);
        }

        void traverse(Node *node)
        {
            if (node == nullptr)
                return;
            traverse(node->left);
            cout << "[" << node->key << ", (" << "cnt: "<< node->cnt <<", height: " << node->height<< ")] ";
            traverse(node->right);
        }
};

int main()
{
    Solution sol;
    vector<int> v;

    v = vector<int>{};
    assert(sol.reversePairs(v) == 0);
    v = vector<int>{2147483647,2147483647,2147483647,2147483647,2147483647,2147483647};
    assert(sol.reversePairs(v) == 0);
    v = vector<int>{1,3,2,3,1};
    assert(sol.reversePairs(v) == 2);
    v = vector<int>{2,4,3,5,1};
    assert(sol.reversePairs(v) == 3);

    v = vector<int>{};
    for (int i = 0; i < 50000; ++i)
        v.push_back(i);

    assert(sol.reversePairs(v) == 0);
}
