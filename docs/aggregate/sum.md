datafun: sum
=============

Compute sum.



Usage
-----

    Usage: sum [options]

    Options:
      -h, --help              display usage
      -v, --version           display version
      -i, --input <file>      input file, defaults to stdin
      -o, --output <file>     output file, defaults to stdout
      -x, --horizontal        perform operation horizontally instead of vertically
      -c, --chunk <number>    chunk the input to treat as multiple separate input files with <number> of lines


Examples
--------

### 1. Compute sum of list of numbers:

    sum << EOF
    5
    10
    20
    EOF

outputs...

    35



### 2. Compute sum of a CSV input vertically:

The sum of the column.

i.e. (1+5+9+13), (2+6+10+14), etc

    sum << EOF
    1,2,3,4
    5,6,7,8
    9,10,11,12
    13,14,15,16
    EOF

outputs...

    28,32,36,40



### 3. Compute sum of a CSV input horizontally:

The sum of the row.

i.e. (1+2+3+4), (5+6+7+8) etc

    sum -x << EOF
    1,2,3,4
    5,6,7,8
    9,10,11,12
    13,14,15,16
    EOF

outputs...

    10
    26
    42
    58



### 4. Compute sum of a CSV input veritically (chunked):

You can think of this as treating the input as if it was separate input files split into chunks. You
could use the POSIX `split` utility or the `splitlines` util in the package to accomplish the same
thing. This just saves you from having to call an additional program.

In other words, this takes each column and splits it into new data sets that contain `n` items as input by `-c` or `--chunk`.

i.e. (1+5), (2+6) etc ... (9+13)...

    sum -c 2 << EOF
    1,2,3,4
    5,6,7,8
    9,10,11,12
    13,14,15,16
    EOF

outputs...

    6,8,10,12
    22,24,26,28


### 5. Compute sum of a CSV input horizontally (chunked):

This takes each row and splits it into new data sets that contain `n` items as input by `-c` or `--chunk`.

i.e. (1+2), (3+4) ... (5+6)...


    sum -c 2 -x << EOF
    1,2,3,4
    5,6,7,8
    9,10,11,12
    13,14,15,16
    EOF

outputs...

    3,7
    11,15
    19,23
    27,31



