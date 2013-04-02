datafun: inv
=============

Compute the inverse.



Usage
-----

    Usage: inv [options]

    Options:
      -h, --help             display usage
      -v, --version          display version
      -i, --input <file>     input file, defaults to stdin
      -o, --output <file>    output file, defaults to stdout


Example
-------

### 1 Compute the inverse

    inv << EOF
    1,2,3,4
    5,6,7,8
    9,10,11,12
    13,14,15,16
    EOF

outputs...

    1,0.5,0.3333333333333333,0.25
    0.2,0.16666666666666666,0.14285714285714285,0.125
    0.1111111111111111,0.1,0.09090909090909091,0.08333333333333333
    0.07692307692307693,0.07142857142857142,0.06666666666666667,0.0625





