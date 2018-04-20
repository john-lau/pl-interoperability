# pl-interoperability
Comparing the interoperability of Golang with C/C++ and Python with C/C++

### Golang
Benefits of using Golang?
- Very fast compile time
- Concurrency support at the language level
- Garbage Collection is implemented
- Extremely scalable: Go is built out of asynchronous tasks with message passing 
- Overall sort of like a low-level language with high-level syntax 



### Golang with C

cgo command
- Enables for the creation of Go packages that call C code
- Generate code to access C library functions and global variables
- Go's solution for integration with exiting software, infrastructure, and protocols


Points of interest to investigate
- Go is a garbage collected runtime while C is not. Passing data must be done carefully
- Call overhead: compare function call overhead differences between just Go vs CGo



### CPython3 with C
- CPython itself is written in C so there is a Python API that allows access to most aspects of Python's runtime system
- Writing C modules for Python allows for two major benefits:
  1. New built-in object types
  2. Ability to call C library functions and system calls
