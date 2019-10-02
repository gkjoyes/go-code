# Benchmark Profiling

Using benchmarks you can profile your programs and see exactly where your performance or memory is being taken.

## Memory Profiling

Run the benchmark.

```sh
    go test -run none -bench . -benchtime 3s -benchmem -memprofile p.out
```

Run the pprof tool.

```sh
    go tool pprof -<option> p.out
```

Run these pprof commands.

```sh
    (pprof) list algOne
    (pprof) web list algOne
```

Run the pprof tool using the browser based tooling.

```sh
    go tool pprof -<option> -http :3000 p.out
```

__Note:__ You need `dot` and `gv` installed.

```sh
    apt-get install graphviz gv
```

Navigate the drop down menu in the UI.

Run the pprof tool using including the `memcpu.test` binary

```sh
    go tool pprof -<option> -http :3000 memcpu.test p.out
```

When you do this, you can get profiling information down to the assembly level.

If you want to reduce memory consumption, look at the `-inuse_space` profile collected during normal program operation.

```sh
    -inuse_space        Allocations live at the time of profile.
    -inuse_objects      Number of bytes allocated at the time of profile.
```

If you want to improve execution speed, look at the `-alloc_objects` profile collected after significant running time or at program end.

```sh
    -alloc_space        All allocations happened since program start ** default.
    -alloc_objects      Number of objects allocated at the time of profile.
```

## CPU Profiling

Run the benchmark.

```sh
    go test -run none -bench . -benchtime 3s -benchmem -cpuprofile p.out
```

Run the pprof tool using the command line tooling.

```sh
    go tool pprof p.out
```

Run these pprof commands.

```sh
    (pprof) list algOne
    (pprof) web list algOne
```

Run the pprof tool using the browser based tooling.

```sh
    go tool pprof -http :3000 p.out
```

Navigate the drop down menu in the UI.

Run the pprof tool using including the `memcpu.test` binary.

```sh
    go tool pprof -http :3000 memcpu.test p.out
```

When you do this, you can get profiling information down to the assembly level.
