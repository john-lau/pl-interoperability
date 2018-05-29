#!/usr/bin/env python

import sys
from time import time

def poly(x):
    mu = 10.0
    pol =[0] * 100
    r = range(0,100)

    for j in r:
        pol[j] = mu = (mu + 2.0) / 2.0
    su = 0.0
    for j in r:
        su = x * su + pol[j]
    return su

if __name__ == "__main__":
    # start = time()
    n = int(sys.argv[1])
    x = float(sys.argv[2])
    pu = 0.0
    for _ in range(0,n):
        pu += poly(x)
    # print(pu)
    # end = time()
    # print (end - start)

