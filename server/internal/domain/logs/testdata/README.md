# Log testdata

Some of the test data is not pushed to Git because of it's size.

For the list of ignored files, see the `.gitignore`

## Generating the data

To generate the data, go into this package, and run:

```sh
go run . > chunk_long.log
```

## Benchmarks

To run the benchmarks, open `process_chunk_bench_test.go`, uncomment the contents, then go into the `logs` domain directory and run

```sh
go test -bench=.
```
