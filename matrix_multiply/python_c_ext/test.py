#!/usr/bin/env python

import array
import sys
import testmodule
import time
xrange = range

def matmul(a, b):
    c = [array.array("d", [0.0]) * len(b) for _ in xrange(len(b[0]))]
    for i in xrange(len(b[0])):
        for j in xrange(len(b)):
            c[i][j] = b[j][i]

    d = [array.array("d", [0.0]) * len(b[0]) for _ in xrange(len(a))]

    for i in xrange(len(a)):
        for j in xrange(len(b[0])):
            s, ai, cj = 0.0, a[i], c[j]
            for k in xrange(len(b)):
                s += ai[k] * cj[k]
            d[i][j] = s
    return d

def build_matrix(n):
    t = 1.0 / n / n
    m = [[0.0] * n for _ in xrange(n)]
    for i in xrange(n):
        for j in xrange(n):
            m[i][j] = t * (i - j) * (i + j)
    return m

def main(argv):
    start_time = time.time()
    n = 100
    if len(sys.argv) > 1:
        n = int(sys.argv[1])

    a = build_matrix(n)
    b = build_matrix(n)
    # print(a)
    # print(b)

    result = testmodule.matmul(n, a, b)

    # print(result)

    # print("--- %s seconds ---" % (time.time() - start_time))

    time_taken = time.time() - start_time

    with open("mat_mul_python_c_1000.txt", "a") as text_file:
        print("%s" % time_taken)
        text_file.write(str(time_taken) + "\n")

if __name__ == "__main__":
    main(sys.argv)
