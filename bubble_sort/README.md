# Bubble Sort

Input N is the size of the left and right singular vectors.

# How To Use Scripts

### Measuring Time
The __time.rb__ ruby script simply runs the desired program 50 times. We are measuring time it takes to run the written code itself so we calculated the runtime within each respective program and then output it. As such, if we run the program 50 times we will output 50 different runtimes for each time the program is executed. Calculating runtime within the program itself is very simple in both Python and Golang; you can observe the commented out lines in both the Python/Golang versions of the program that measure time. 

### Measuring Memory
Measuring memory was a little harder because it varied depending on the implementation. We will detail those differences in the following sections.

##### Python and Python w/ C
This is simplest implementation to measure memory for as we simply used our __mem.rb__ script which runs the desired program 50 times while a concurrent thread constantly running the _ps_ command with the _rss_ flag to monitor the memory usage the program, tracking the max memory usage. 

##### Golang
We used the _runtime_ library, part of the Golang profiling library _pprof_, to capture the memory usage of each execution of the program. 

The lines that captured memory usage in the Go program are currently commented out. If you want to capture memory usage please uncomment.

##### CGo
Unfortunately the _runtime_ library doesn't capture any memory that is allocated in C. As such, we had to use a second Golang library, _cmemory_ in order to capture memory allocated in C. We combined the recored memory usage from both _runtime_ and _cmemory_ libraries in order to get the total memory usage of CGo programs.

The lines that captured memory usage in the Go program are currently commented out. If you want to capture memory usage please uncomment. Remember that both _runtime_ and _cmemory_ is used to capture the memory usage.



