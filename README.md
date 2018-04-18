# pl-interoperability
Comparing the interoperability of Golang with C/C++ and Python with C/C++

# golang with C

cgo command
- enables for the creation of Go packages that call C code.

Points of interest to investigate
- Go is a garbage collected runtime while C is not. Passing data must be done carefully
- Call overhead: compare function call overhead differences between just Go vs CGo
