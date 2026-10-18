# Go Concurrency Examples

Small, runnable examples for learning goroutines, channels, wait groups, and
the producer-consumer pattern in Go.

## Project layout

```text
.
├── main.go
└── examples/
    ├── channels.go
    ├── combined.go
    ├── goroutines.go
    ├── producerconsumer.go
    └── waitgroups.go
```

The root `main.go` selects the example to run. The example implementations are
kept in the `examples` package so each topic can be studied independently.

## Run the current example

From the project root:

```bash
go run .
```

The current entry point runs `examples.JobHandler`, which demonstrates:

- A producer sending jobs through a channel.
- A consumer receiving jobs with a receive-only channel.
- Closing a channel after production is complete.
- Waiting for the consumer with `sync.WaitGroup`.

## Try another example

Change the function called in `main.go` to one of the exported examples:

```go
examples.RoutineExample()
examples.ChannelReceiving()
examples.WaitGroupTask()
examples.WorkerWorks()
examples.JobHandler()
```

Then run the project again:

```bash
go run .
```

## Test all packages

```bash
go test ./...
```
