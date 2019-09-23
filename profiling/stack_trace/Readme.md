# Stack Traces and Core Dumps

Having some basic skills in debugging Go programs can save any programmer a good amount of time trying to identify problems. I believe in logging as much information as you can, but sometimes a panic occurs and what you logged is not enough. Understanding the information in a stack trace can sometimes mean the difference between finding the bug now or needing to add more logging and waiting for it to happen again. We can also stop a running program and get Core Dump which also generates a stack trace.

- Stack traces are an important tool in debugging an application.
- The runtime should never panic so the trace is everything.
- You can see every goroutine and the call stack for each routine.
- You can see every value passed into each function on the stack.
- You can generate core dumps and use these same techniques.
