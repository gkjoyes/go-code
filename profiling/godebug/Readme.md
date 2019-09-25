# GODEBUG

We can get specific trace information about the garbage collection and  scheduler using the GODEBUG environmental variable. The variable will cause the runtime to emit tracing information.

## Schedule tracing

```sh
    export GODEBUG=schedtrace=1000
```

- __scheddetail__: setting schedtrace=X and scheddetail=1 causes the scheduler to emit detailed multiline info every X milliseconds, describing state of the scheduler, processor, threads and goroutines.
- __schedtrace__: setting schedtrace=X causes the scheduler to emit a single line to standard error every X milliseconds, summarizing the scheduler state.

```txt
SCHED 7030ms: gomaxprocs=2 idleprocs=0 threads=6 spinningthreads=0 idlethreads=2 runqueue=129 [128 68]

gomaxprocs=2        Number of logical processors.
idleprocs=0         Number of idle logical processors.
threads=6           Threads in use.
spinningthreads=0   Threads in hold state.
idlethreads=2       Threads not in use.
runqueue=129        Goroutines in the global queue.
[128 68]            Goroutines in each of the logical processors.
```

## Generating a Schedular Trace

### Case 1

- Build and run the example program using a single logical processor.

    ```sh
        go build
        GOMAXPROCS=1 GODEBUG=schedtrace=1000 ./godebug
    ```

- Put some load of the web application.

    ```sh
        hey -m POST -c 8 -n 1000000 "http://localhost:4000/sendjson"
    ```

- Look at the load on the logical processor. We can only see runnable goroutines. After 5 seconds we don't see any more goroutines in the trace.

### Case 2

- Run the example program but leak goroutines.

    ```sh
    GOMAXPROCS=1 GODEBUG=schedtrace=1000 ./godebug leak
    ```

- Put some load of the web application.

    ```sh
        hey -m POST -c 8 -n 1000000 "http://localhost:4000/sendjson"
    ```

- Look at the load on the logical processor. We can only see runnable goroutines. After 5 seconds we still see goroutines in the trace.

## GC Tracing

There is no way to identify specifically in the code where a leak is occurring. We can validate if a memory leak is present and  which functions or methods are producing the most allocations.

Setting __gctrace=1__ causes the garbage collector to emit a single line to standard error at each collection, summarizing the amount of memory collected and the length of the pause. Setting gctrace=2 emits the same summary but also repeats each collection.

## Generating a GC Trace

- Build and run the example program.

    ```sh
        go build
        GODEBUG=gctrace=1   ./godebug
    ```

- Put some load of the web application.

    ```sh
        hey -m POST -c 8 -n 1000000 "http://localhost:4000/sendjson"
    ```

- Review the gc trace.

    ```txt
        gc 1 @4.248s 0%: 0.018+15+0.006 ms clock, 0.075+31/15/0+0.024 ms cpu, 4->4->0 MB, 5 MB goal, 4 P
    ```

    ```txt
        gc 1                        the GC number, incremented at each GC.
        @4.248s                     time in seconds since program start.
        0%                          percentage of time spent in GC since program start.
        0.018+15+0.006 ms clock     wall-clock time for the phases of the GC.
        0.075+31/15/0+0.024 ms cpu  CPU time for the phases of the GC.
        4->4->0 MB                  heap size at GC start, at GC end, and live heap.
        5 MB goal                   goal heap size.
        4 P                         number of processors used.
    ```

- __wall-clock__ time is a measure of the real time that elapses from start to end, including time that passes due to programmed delays or waiting for resources to become available.
- __CPU time__ is the amount of time for which a central processing unit(CPU) was used for processing instructions of a computer program or operating system.
- You can get more details by adding the __gcpacertrace=1__ flag. This causes the garbage collector to print information about the internal state of the concurrent pacer.

    ```sh
        export GODEBUG=gctrace=1, gcpacertrace=1
    ```

- Setting gctrace to any value > 0 also causes the garbage collector to emit a summary when memory is released back to the system. This process of returning memory to the system is called scavenging.

    ```txt
        scvg: 1 MB released
    ```

    ```txt
        1 MB printed only if non-zero
    ```

    ```txt
        scvg: inuse: 2, idle: 60, sys: 63, released: 58, consumed: 4 (MB)
    ```

    ```txt
        inuse: 2        MB used or partially used spans.
        idle: 60        MB spans pending scavenging.
        sys: 63         MB mapped from the system.
        released: 58    MB released to the system.
        consumed: 4     MB allocated from the system.
    ```

