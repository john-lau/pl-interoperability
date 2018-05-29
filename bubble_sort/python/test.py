#!/usr/bin/env python

from ctypes import *
import random
import sys
import time

# def convert_list_to_array(lll):
#         intarraytype=c_int*len(lll)
#         intarray=intarraytype()
#         index=0
#         for l in lll:
#                 intarray[int(index)]=int(l)
#                 index=index+1
#         return intarray

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

        # iarray = convert_list_to_array(lr)

        # print(lr)
        py_sort(lr, len(lr))
        # print("finished sorting array")
        time_taken = time.time() - start_time

        with open("bubble_sort_python_25000.txt", "a") as text_file:
            print("%s" % time_taken)
            text_file.write(str(time_taken) + "\n")

        

        # print(lr)

if __name__ == "__main__":
    main(sys.argv)
