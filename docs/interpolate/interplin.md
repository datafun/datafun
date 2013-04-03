datafun: interplin
==================

Perform linear interpolation on a one-dimensional data set.



Usage
-----

      Usage: interplin [options]

      Options:
        -h, --help              display usage
        -v, --version           display version
        -i, --input <file>      input file, defaults to stdin
        -o, --output <file>     output file, defaults to stdout
        -n, --number [count]    number of points to insert between each point


Examples
--------

Perform linear interpolation on the following dataset inserting two data points between each.

    interplin -n 2 << EOF
    0
    10
    20
    EOF

outputs...

    0
    3.3333333333333335
    6.666666666666667
    10
    13.333333333333334
    16.666666666666668
    20






