datafun: scale
=============

Scale by a scaling factor.



Usage
-----

    Usage: scale [options]

    Options:
      -h, --help             display usage
      -v, --version          display version
      -i, --input <file>     input file, defaults to stdin
      -o, --output <file>    output file, defaults to stdout
      -n, --number <num>     the number to scale by


Example
-------

### 1 Scale the dataset

    scale -n 0.5 << EOF
    1,2,3,4
    5,6,7,8
    9,10,11,12
    13,14,15,16
    EOF

outputs...

    0.5,1,1.5,2
    2.5,3,3.5,4
    4.5,5,5.5,6
    6.5,7,7.5,8





