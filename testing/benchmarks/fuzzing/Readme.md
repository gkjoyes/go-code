# Go Fuzz

[Go-fuzz](https://github.com/dvyukov/go-fuzz) is a coverage-guided fuzzing solution for testing of Go packages. Fuzzing is mainly applicable to packages that parse complex inputs (both text and binary), and is especially useful for hardening of systems that parse inputs from potentially malicious users.

- Fuzzing allows you to find cases where your code panics.
- Once you identify data inputs that causes panics, code can be corrected and tests created.
- Table tests are an excellent choice for these input data panics.

## Working

- Run the `go-fuzz-build` tool against the package to generate the fuzz zip file. The zip file contains all the instrumented binaries go-fuzz is going to use while fuzzing. Any time the source code changes this needs to be re-run.

    ```sh
        go-fuzz-build github.com/george-kj/go-code/testing/benchmarks/fuzzing/example1
    ```

- Perform the actual fuzzing by running the `go-fuzz` tool and find data inputs that cause panics. Run this until you see an initial crash.

    ```sh
        go-fuzz -bin=./api-fuzz.zip -workdir=workdir/corpus
    ```

- Review the `crashers` folder under the `workdir/corpus` folders. This contains panic information. You will see an issue when the data passed into the web call is empty. Fix the `Process` function and add the table data to the test.

    ```sh
        {"/process", http.StatusBadRequest, []byte(""), `{"Error":"The Message"}`},
    ```

- Run the build and start fuzzing again.

    ```sh
        rm -rf workdir/crashers and workdir/supressions
        go-fuzz -bin=./api-fuzz.zip -dup -workdir=workdir/corpus
    ```
