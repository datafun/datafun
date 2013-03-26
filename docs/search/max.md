

max
---

Computes the maximum number from the input. Can also output the index of the maximum.


### usage

    Usage: max [options]

    Options:

      -V, --version                 output program version
      -h, --help                    output help information
      -i, --input <path>            specify input file [stdin]
      -o, --output <path>           specify output file [stdout]
      -n, --number [number]         specify the number [0]
      -s, --start [number]          the start of the data, x0 [0]
      -x, --show-index              show the index along with the maximum 


### examples

**input: (data.txt)**

    45
    23.0
    -10000.34
    10442
    0.42424
    2
    45.0
    -5.678
    6


#### 1. find maximum

**cmd:**

    max << data.txt


**output:**

    10442


#### 2. also get the index, or the line that it resides at

**cmd:**

    max -x << data.txt


**output:**

    3,10442.000000


#### 3. get the index, but convert it to something significant that is applicable to your data set

**cmd:**

    max -s 100 -n 0.5 << data.txt


**output:**

    101.500000,10442.000000



