datafun: abs
=============

Compute absolute value.



Usage
-----

    Usage: abs [options]

    Options:
      -h, --help             display usage
      -v, --version          display version
      -i, --input <file>     input file, defaults to stdin
      -o, --output <file>    output file, defaults to stdout


Example
-------

### 1 Compute absolute value of a list

    abs << EOF
    1
    -3
    -3.4
    4.4
    EOF

outputs...

    1
    3
    3.4
    4.4


### 2 Compute absolute vlaue of CSV input

    abs << EOF
    1,-2,3
    -4,5,6
    7,8,-9
    EOF

outputs...

    1,2,3
    4,5,6
    7,8,9


