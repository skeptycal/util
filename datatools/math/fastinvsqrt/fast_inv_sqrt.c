float Q_rsqrt( float number )
{
    long i;                                                        // 32-bit number
        // standard binary integer
        // 00000000 00000000 00000000 00000000
    float x2, y;                                                // 32-bit decimal number
        /*
            //+ IEEE 754 standard floating point number (a.k.a. binary scientific notation)
            //+ normalised numbers

            M is 23 bits
            E is 8 bits
            S is always 1 (in this example)

            //+ bit representation is                2^23 * E + M (shift E by 23 bits and add M)
            //+ decimal number is                   (1 + M/2^23) * 2^(E-127)

                0   0000000 000000000000000000000000
                |           |           |
                |           |           |--> 23 bit mantissa - in binary the mantissa is unique, the only
                |           |                   non zero number before the decimal point is 1 ... so it is always 1
                |           |                   this means that the 1 is assumed and not represented
                |           |
                |           |                   range of numbers is 0 to 2^23-1
                |           |
                |           |--> 8 bit exponent (negative sign bit is used ... so the range is -127 to 128)
                |
                |--> sign bit (ignored in this algorithm ... Real square roots are always of positive numbers ... )
        */
    const float threehalfs = 1.5F;                  // 1.5 (also 32-bit)

    x2 = number * threehalfs;
    y = number;
    i = * ( long * ) &y;                                    // evil floating point bit hack
    i = 0x5f3759df - ( i >> 1 );                        // what the fuck?
    y = * ( float * ) &i;
    y = y * ( threehalfs - ( x2 * y * y ) );        // 1st iteration
    // y = y * ( threehalfs - ( x2 * y * y ) );     // 2nd iteration, can be removed
}
