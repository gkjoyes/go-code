# Blocking Profiling

Testing and Tracing allows us to see blocking profiles.

## Running a Test Based Blocking Profile

We can get blocking profiles by running a test.

Generate a block profile from running the test.

```sh
    go test -blockprofile block.out
```

Run the pprof tool to view the blocking profile.

```sh
    go tool pprof block.out
```

Review the TestLatency function.

```sh
    list TestLatency
```

## Running a Trace

Once you have a test established you can use the *-trace* *trace.out* option with the go test tool.

Generate a trace from running the test.

```sh
    go test -trace trace.out
```

Run the trace tool to review the trace.

```sh
    go tool trace trace.out
```
