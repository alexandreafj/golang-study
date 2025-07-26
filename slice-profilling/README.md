# Profiling Go Slice Performance

This guide demonstrates how to benchmark Go code and visualize performance data to understand the overhead of slice re-allocations.

## Step 1: Run the Benchmarks with Memory Stats

First, we'll execute the tests to observe the raw performance and memory statistics. Navigate to your `slice-profiling` directory in the terminal and run the following command:

```bash
go test -bench=. -benchmem
````

  - `go test`: The command used to run tests.
  - `-bench=.`: Instructs Go to run all benchmarks in the current directory.
  - `-benchmem`: This crucial flag tells Go to include memory allocation statistics in the output.

You will see output similar to this, which includes the `B/op` (bytes per operation) and `allocs/op` (allocations per operation) columns:

```
goos: darwin
goarch: amd64
pkg: github.com/alexandreafj/golang-study/slice-profilling
cpu: VirtualApple @ 2.50GHz
BenchmarkCreateSliceNonPerformant-10                1785            600045 ns/op         4101398 B/op         28 allocs/op
BenchmarkCreateSlicePerformant-10                   9686            107022 ns/op          802823 B/op          1 allocs/op
PASS
ok      github.com/alexandreafj/golang-study/slice-profilling   2.709s
```

**Immediate Observation:**

  * **Non-Performant:** Took 600,045 nanoseconds and resulted in 28 allocations.
  * **Performant:** Took only 107,022 nanoseconds and made just 1 allocation (the initial `make` call). It is over five times as fast and significantly reduces memory allocation overhead\!

## Step 2: Generate a Memory Profile

The primary difference highlighted here is memory allocation. Let's create a profile specifically focused on this aspect.

Run this command:

```bash
go test -bench=NonPerformant -memprofile memprofile.out
```

  - `-bench=NonPerformant`: Runs only the specific benchmark we intend to profile.
  - `-memprofile memprofile.out`: Directs Go to save the memory profiling data to a file named `memprofile.out`.

### Generating a CPU Profile

This profile shows which functions are consuming the most CPU time.

Run the following command in your terminal:

```bash
go test -bench=. -cpuprofile cpu.out
```

This will run the benchmarks and create a new file in your directory named cpu.out. This is the file you will upload.

## Step 3: Analyze the Profile with `pprof`

Now, we will use `go tool pprof` to analyze the generated file. This tool is a powerful component of Go's profiling suite.

Execute this command to open an interactive `pprof` session:

```bash
go tool pprof memprofile.out
```

You will see a prompt like `(pprof)`. You are now inside the `pprof` tool.



## Step 4: Visualize the Data (The "Aha\!" Moment)

The most effective way to understand the cost is through a visual graph. Type `web` into the `pprof` prompt and press Enter.

```
(pprof) web
```

This command will:

  * Generate a visual graph (as an `.svg` file) illustrating where your program spent its time allocating memory.
  * Open this `.svg` file in your default web browser.

You will see a graph similar to the following (visual representation not included here, but imagine a call graph).

**How to Read the Graph:**

  * **Boxes:** Each box represents a function call.
  * **Box Size:** The larger the box, the more memory that function allocated.
  * **Arrows:** Indicate which function called another. A thicker arrow signifies a more frequently taken call path.

In the generated graph, you will observe a very large box for `runtime.growslice`. This is the internal Go function invoked to re-allocate and copy a slice when it exceeds its current capacity. You can clearly see that it is being called by our `createSliceNonPerformant` function and is responsible for nearly all of the memory allocation within the program.

If you were to run the same profiling steps on the `Performant` version, the `runtime.growslice` box would be minuscule or entirely absent. This provides clear visual evidence of the performance cost associated with unnecessary re-allocations\!