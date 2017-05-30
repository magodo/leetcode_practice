/*************************************************************************
 Author: Zhaoting Weng
 Created Time: Sat 25 Feb 2017 09:40:16 PM CST
 Description: 
 ************************************************************************/

#include "CSolution.h"
#include "math.h"

int CSolution::m_max_msb = 30;

int CSolution::msbBit(int n)
{
    int upper = m_max_msb+1, lower = 0;
    int msb= (upper+lower)/2;
    int tmp;

    while ((upper - lower) > 1)
    {
        tmp = pow(2, msb);
        if (n < tmp)
            upper = msb;
        else if (n > tmp)
            lower = msb;
        else // n == pow(2,msb)
            break;
        msb = (lower+upper)/2;
    }
    return msb;

}

int CSolution::rangeBitwiseAnd(int m, int n)
{
    /*
     * If the lower limit and upper limit have different numbers of
     * valid bits, then the bitwise AND of numbers in the range is 0;
     * If they have the same length, then the result equals to:
     * same MSBs of m and n, attch 0 as the other LSBs (from the bit
     * when they are different).
     *
     * E.g. 0b101011001
     *      to
     *      0b101010000
     *
     *      Results into:
     *
     *      ob101010000
     */

    /* if the bit size is not the same, then n AND m == 0 */
    if (msbBit(n) != msbBit(m))
    {
        return 0;
    }
    
    /* From msb to lsb, find the 1st different bit. Before this point(
     * exclusive), the bit remain as it was; After this point, the bit
     * are reset to 0
     */

    int msb = msbBit(n);
    int result = 0;
    int mask = 1<<msb;

    while (((n & mask) == (m & mask)) && (mask > 0))
    {
        result += n&mask;
        mask >>= 1;
    }
    return result;
}
