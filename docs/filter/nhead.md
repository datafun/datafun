datafun: nhead
==============

Not head or inverse head. If `head -n 5` returns the first 5 lines of a file. `nhead -n 5` skips the first 5 lines



Usage
-----

    Usage: nhead [options]

    Options:
      -h, --help display usage
      -V, --version display version
      -i, --input input file, defaults to stdin
      -o, --output output file, defaults to stdout
      -n, --number number of lines to skip, defaults to 0


Example
-------

skip first two lines:

    nhead -n 2 << EOF
    5
    10
    15
    20
    EOF

outputs...

    15
    20





