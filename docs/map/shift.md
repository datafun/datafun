datafun: shift
=============

Shift the data set by a value.



Usage
-----

    Usage: shift [options]

    Options:
      -h, --help             display usage
      -v, --version          display version
      -i, --input <file>     input file, defaults to stdin
      -o, --output <file>    output file, defaults to stdout
      -n, --number <num>     the number to shift by


Example
-------

### 1 Shift the dataset

Shift the data to the left 2

    scale -n -2 << EOF
    1,2,3,4
    5,6,7,8
    9,10,11,12
    13,14,15,16
    EOF

outputs...

   -1,0,1,2
    3,4,5,6
    7,8,9,10
    11,12,13,14





