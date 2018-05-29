#!/usr/bin/env python

from sys import argv
from time import time
import testmodule
import sys


if __name__ == "__main__":
    # start = time()
    n = int(argv[1])
    x = float(argv[2])
    pu = 0.0
    for _ in range(0,n):
        pu += testmodule.poly(x)
    # print(pu)
    # end = time()
    # print (end - start)
