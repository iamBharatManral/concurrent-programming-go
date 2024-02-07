### Condition variables:
1. Rather than waiting indefinitely for lock to be released, we conditionally suspends the execution of thread.

### Semaphores:
1. Rather than single thread have exclusive access to a resource, semaphores allows multiple concurrent executions/accesses.
2. A semaphore with N = 1 concurrent executions gives us same functionality as mutexes. It is also known as binary semaphores.
3. 