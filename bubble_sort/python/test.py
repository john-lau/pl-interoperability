from ctypes import *
import random
import sys

def convert_list_to_array(lll):
        intarraytype=c_int*len(lll)
        intarray=intarraytype()
        index=0
        for l in lll:
                intarray[int(index)]=int(l)
                index=index+1
        return intarray

def py_sort(iarray,length):
        for i in range(length):
                for j in range(i,length):
                        if iarray[i]>iarray[j]:
                                tmp=iarray[i]
                                iarray[i]=iarray[j]
                                iarray[j]=tmp

def main(argv):
        n = 100
        if len(sys.argv) > 1:
                n = int(sys.argv[1])

        sample_size=int(sys.argv[1])

        lr= random.sample(range(10000000),sample_size)
        iarray = convert_list_to_array(lr)

        py_sort(iarray, len(lr))
        print("finished sorting array")

if __name__ == "__main__":
    main(sys.argv)