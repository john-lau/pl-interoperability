# pl-interoperability
Comparing the interoperability of Golang with C/C++ and Python with C/C++

### Golang with C

cgo command
- enables for the creation of Go packages that call C code.

Points of interest to investigate
- Go is a garbage collected runtime while C is not. Passing data must be done carefully
- Call overhead: compare function call overhead differences between just Go vs CGo

### CPython3 with C
- CPython itself is written in C so there is a Python API that allows access to most aspects of Python's runtime system
- Writing C modules for Python allows for two major benefits:
  1. New built-in object types
  2. Ability to call C library functions and system calls
