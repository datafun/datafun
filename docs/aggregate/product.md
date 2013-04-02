datafun: product
=============

Compute product.



Usage
-----

    Usage: product [options]

    Options:
      -h, --help              display usage
      -v, --version           display version
      -i, --input <file>      input file, defaults to stdin
      -o, --output <file>     output file, defaults to stdout
      -x, --horizontal        perform operation horizontally instead of vertically
      -c, --chunk <number>    chunk the input to treat as multiple separate input files with <number> of lines


Examples
--------

### 1. Compute product of list of numbers:

    product << EOF
    5
    10
    20
    EOF

outputs...

    1000



### 2. Compute product of a CSV input vertically:

The product of the column.

i.e. (1*5*9*13), (2*6*10*14), etc

    product << EOF
    1,2,3,4
    5,6,7,8
    9,10,11,12
    13,14,15,16
    EOF

outputs...

    585,1680,3465,6144



### 3. Compute product of a CSV input horizontally:

The product of the row.

i.e. (1*2*3*4), (5*6*7*8) etc

    product -x << EOF
    1,2,3,4
    5,6,7,8
    9,10,11,12
    13,14,15,16
    EOF

outputs...

    24
    1680
    11880
    43680



### 4. Compute product of a CSV input veritically (chunked):

You can think of this as treating the input as if it was separate input files split into chunks. You
could use the POSIX `split` utility or the `splitlines` util in the package to accomplish the same
thing. This just saves you from having to call an additional program.

In other words, this takes each column and splits it into new data sets that contain `n` items as input by `-c` or `--chunk`.

i.e. (1*5), (2*6) etc ... (9*13)...

    product -c 2 << EOF
    1,2,3,4
    5,6,7,8
    9,10,11,12
    13,14,15,16
    EOF

outputs...

    5,12,21,32
    117,140,165,192


### 5. Compute product of a CSV input horizontally (chunked):

This takes each row and splits it into new data sets that contain `n` items as input by `-c` or `--chunk`.

i.e. (1*2), (3*4) ... (5*6)...


    product -c 2 -x << EOF
    1,2,3,4
    5,6,7,8
    9,10,11,12
    13,14,15,16
    EOF

outputs...

    2,12
    30,56
    90,132
    182,240    



