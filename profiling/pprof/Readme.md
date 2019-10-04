# http/pprof Profiling

Using the http/pprof support you can profile your web applications and services to see exactly where your performance or memory is being taken.

## Building and Running the Project

Build and run the example program.

```sh
    go build
    ./pprof
```

Test it is working.

```sh
    http://localhost:4000/sendjson
```

To add load to the service while running profiling we can run these command.

Send 1M request using 8 connections.

```sh
    hey -m POST -c 8 -n 1000000 "http://localhost:4000/sendjson"
```

## Raw http/pprof

We already added the following import so we can include the profiling route to our web service.

```sh
    import _ "net/http/pprof"
```

Look at the basic profiling stats from the new endpoint.

```sh
    http://localhost:4000/debug/pprof
```

Run a single search from the Browser and then refresh the profile information.

```sh
    http://localhost:4000/sendjson
```

Put some load of the web application. Review the raw profiling information once again.

```sh
    hey -m POST -c 8 -n 1000000 "http://localhost:4000/sendjson"
```

## Interactive Profiling

### Heap Profiling

Run the pprof tool.

```sh
    go tool pprof -<option> ./pprof http://localhost:4000/debug/pprof/heap
```

Options help to see pressure on heap over time:

```sh
    -alloc_space:       All allocations happened since program start (default).
    -alloc_objects:     Number of object allocated at the time of profile.
```

Options help to see the current status of the heap:

```sh
    -inuse_space:       Allocations live at the time of profile.
    -inuse_objects:     Number of bytes allocated at the time of profile.
```

If you want to reduce memory consumption, look at the `-inuse_space` profile collected during normal program operation.

If you want to improve execution speed, look at the `-alloc_objects` profile collected after significant running time or at program end.

### CPU Profiling

Run the Go pprof tool in another window or tab to review cpu information.

```sh
    go tool pprof http://localhost:4000/debug/pprof/profile
```

If you include the binary when using the browser tooling, you can get information down to the assembly.

```sh
    go tool pprof -http :3000 ./pprof http://localhost:4000/debug/pprof/profile
```

Note that goroutines in `syscall` state consume an OS thread, other goroutines do not(except for goroutines that called `runtime.LockOSThread`, which is, unfortunately, not visible in the profile).

Note that goroutines in `IO wait` state do NOT consume an OS thread. They are parked on the non-blocking network poller.

## Comparing Profiles

Take a snapshot of the current heap profile. Then do the same for the cpu profile.

```sh
    curl -s http://localhost:4000/debug/pprof/heap > base.heap
```

After some time, take another snapshot:

```sh
    curl -s http://localhost:4000/debug/pprof/heap > current.heap
```

Now compare both snapshots against the binary and get into the pprof tool:

```sh
    go tool pprof -inuse_space -base base.heap current.heap
```
