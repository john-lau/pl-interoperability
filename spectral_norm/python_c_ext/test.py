#!/usr/bin/env python

from array import array
from math import sqrt
from sys import argv
from time import time
import testmodule
import sys
xrange = range

def main():
    # start = time()
    n = int(argv[1])
    u = [1.0] * n
    v = [1.0] * n

    for dummy in xrange (10):
        testmodule.eval_AtA_times_u(n, u, v)
        testmodule.eval_AtA_times_u(n, v, u)

    vBv = vv = 0

    for ue, ve in zip(u, v):
        vBv += ue * ve
        vv  += ve * ve

    # end = time()
    # print (end - start)
    # print("%0.9f" % (sqrt(vBv/vv)))

main()
