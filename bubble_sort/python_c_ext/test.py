#!/usr/bin/env python

from ctypes import *
import random
import testmodule
import time
import sys

def py_sort(iarray,length):
        for i in range(length):
                for j in range(i,length):
                        if iarray[i]>iarray[j]:
                                tmp=iarray[i]
                                iarray[i]=iarray[j]
                                iarray[j]=tmp

def main(argv):
        start_time = time.time()
        n = 100
        if len(sys.argv) > 1:
                n = int(sys.argv[1])

        sample_size=int(sys.argv[1])

        lr= random.sample(range(10000000),sample_size)

        print(lr)

        # Call C function here
        testmodule.c_sort(sample_size, lr)


        py_sort(lr, len(lr))
        print("finished sorting array")
        print(lr)

        
        time_taken = time.time() - start_time


if __name__ == "__main__":
    main(sys.argv)
