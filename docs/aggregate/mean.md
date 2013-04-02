datafun: mean
=============

Compute mean.



Usage
-----

    Usage: mean [options]

    Options:
      -h, --help              display usage
      -v, --version           display version
      -i, --input <file>      input file, defaults to stdin
      -o, --output <file>     output file, defaults to stdout
      -x, --horizontal        perform operation horizontally instead of vertically
      -c, --chunk <number>    chunk the input to treat as multiple separate input files with <number> of lines


Examples
--------

### 1. Compute mean of list of numbers:

    mean << EOF
    5
    10
    20
    EOF

outputs...

    11.66666666



### 2. Compute mean of a CSV input vertically:

The mean of the column.

i.e. (1*5*9*13)/4, (2,6,10,14)/4, etc

    mean << EOF
    1,2,3,4
    5,6,7,8
    9,10,11,12
    13,14,15,16
    EOF

outputs...

    7,8,9,10



### 3. Compute mean of a CSV input horizontally:

The mean of the row.

i.e. (1*2*3*4)/4, (5,6,7,8)/4, etc

    mean -x << EOF
    1,2,3,4
    5,6,7,8
    9,10,11,12
    13,14,15,16
    EOF

outputs...

    2.5
    4.5
    6.5
    8.5



### 4. Compute mean of a CSV input veritically (chunked):

You can think of this as treating the input as if it was separate input files split into chunks. You
could use the POSIX `split` utility or the `splitlines` util in the package to accomplish the same
thing. This just saves you from having to call an additional program.

In other words, this takes each column and splits it into new data sets that contain `n` items as input by `-c` or `--chunk`.

i.e. (1+5)/2, (2+6)/2 etc ... (9+13)/2...

    mean -c 2 << EOF
    1,2,3,4
    5,6,7,8
    9,10,11,12
    13,14,15,16
    EOF

outputs...

    3,4,5,6
    11,12,13,14


### 5. Compute mean of a CSV input horizontally (chunked):

This takes each row and splits it into new data sets that contain `n` items as input by `-c` or `--chunk`.

i.e. (1+2)/2, (3+4)/2 ... (5+6)/2...


    mean -c 2 -x << EOF
    1,2,3,4
    5,6,7,8
    9,10,11,12
    13,14,15,16
    EOF

outputs...

    1.5,3.5
    5.5,7.5
    9.5,11.5
    13.5,15.5



