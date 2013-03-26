datafun: splitlines
===================

Like POSIX `split`. POSIX `split` has a line limitation. `splitlines` doesn't.



Usage
-----

    Usage: splitlines [options] filetemp

    Options:
      -h, --help display usage
      -V, --version display version
      -i, --input input file, defaults to stdin
      -o, --output output file, defaults to stdout


Example
-------

split the file `data.txt`:

**data.txt**:

    1
    2
    3
    4
    5

running splitlines:

    splitlines -i data.txt -n 2 /tmp/splits.txt

outputs...

    /tmp/0-splits.txt
    /tmp/1-splits.txt
    /tmp/2-splits.txt

with each file having at most 2 lines