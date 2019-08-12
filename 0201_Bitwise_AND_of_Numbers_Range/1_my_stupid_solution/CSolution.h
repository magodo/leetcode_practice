/*************************************************************************
 Author: Zhaoting Weng
 Created Time: Sat 25 Feb 2017 10:17:46 PM CST
 Description: 
 ************************************************************************/

class CSolution 
{
    public:
        int rangeBitwiseAnd(int m, int n);

    private:
        /* return the MSB bit index, starts from 0 */
        int msbBit(int n);

        /* maximum allowed MSB (starts from 0)*/
        static int m_max_msb;
};
