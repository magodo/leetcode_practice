/*************************************************************************
 Author: Zhaoting Weng
 Created Time: Sat 25 Feb 2017 10:14:36 PM CST
 Description: 
 ************************************************************************/

#include <iostream>
#include <cassert>
#include "math.h"
#include "CSolution.h"

using namespace std;

int main()
{
    CSolution sol;

    assert(4 == sol.rangeBitwiseAnd(5,7));
    assert(0 == sol.rangeBitwiseAnd(1,3));
    assert(0b101010000 == sol.rangeBitwiseAnd( 0b101010000, 0b101011001));
    assert((pow(2,31)-2)== sol.rangeBitwiseAnd(pow(2,31)-2, pow(2,31)-1));
}
