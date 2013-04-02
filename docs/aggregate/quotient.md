datafun: quotient
=================

Compute quotient from left-to-right or top-to-bottom.



Usage
-----

    Usage: quotient [options]

    Options:
      -h, --help              display usage
      -v, --version           display version
      -i, --input <file>      input file, defaults to stdin
      -o, --output <file>     output file, defaults to stdout
      -x, --horizontal        perform operation horizontally instead of vertically
      -c, --chunk <number>    chunk the input to treat as multiple separate input files with <number> of lines


Examples
--------

### 1. Compute quotient of list of numbers:

    quotient << EOF
    100
    10
    5
    EOF

outputs...

    2



### 2. Compute quotient of a CSV input vertically:

The quotient of the column.

i.e. (200/50/2), (30/5/2), etc

    product << EOF
    200,30
    50,5
    2,2
    EOF

outputs...

    2,3    







 



