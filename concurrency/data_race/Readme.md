# Data Races

- A data race is when two or more goroutines attempt to read and write to the resource at the same time.

- Race conditions can create bugs that appear totally random or can never surface as they corrupt data.

- Atomic functions and mutexes are a way to synchronize the access of shared resources between goroutines.
