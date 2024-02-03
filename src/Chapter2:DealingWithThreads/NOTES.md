### OS level primitives for concurrency:
1. Process: provides separate execution context
2. Thread: also knows as light-weight process. executes in shared memory space

### Threading models:
1. M:1: multiple user-level threads runs within the context of single kernel-level thread
2. M:N: multiple user-levels threats maps to multiple kernel-level threads. a model used by Golang


### Concurrency Vs Parallelism
1. Concurrency: attribute of the program code. It's about planning how to do many tasks at the same time
2. Parallelism: property of the executing program. It's about performing many tasks at the same time.