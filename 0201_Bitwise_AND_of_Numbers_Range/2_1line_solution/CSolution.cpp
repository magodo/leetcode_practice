/*************************************************************************
 Author: Zhaoting Weng
 Created Time: Sat 25 Feb 2017 11:42:08 PM CST
 Description: 
 ************************************************************************/

#include "CSolution.h"

int CSolution::rangeBitwiseAnd(int m, int n)
{
    return (m < n)? (rangeBitwiseAnd(m>>1, n>>1) << 1):m;
}
