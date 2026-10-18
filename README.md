# Go Concurrency with Goroutines and Channels

Hands-on Go examples for learning goroutines, channels, wait groups, worker
pools, and the producer-consumer pattern.

Repository: [manojtharindu11/go-concurrency-with-routine-and-channels](https://github.com/manojtharindu11/go-concurrency-with-routine-and-channels)

## Requirements

- Go 1.27.1 or newer

Check your installed Go version with:

```bash
go version
```

## GitHub topics

Suggested repository topics:

`go` `golang` `concurrency` `goroutines` `channels` `waitgroup`
`producer-consumer` `worker-pool` `go-examples` `learning-go`

## Project structure

```text
go-concurrency/
├── main.go
├── go.mod
└── examples/
    ├── channels.go
    ├── combined.go
    ├── goroutines.go
    ├── producerconsumer.go
    └── waitgroups.go
```

The root `main.go` is the program entry point. The learning examples are
organized in the `examples` package and expose one exported function each.

## Run the project

From the repository root:

```bash
go run .
```

The default entry point runs the producer-consumer example:

```text
Producing jobs: 1
Consuming job: 1
...
Consumer finished
Program finished
```

The order of producer and consumer messages can vary because both operations
run concurrently.

## Available examples

| Function | File | Concepts |
| --- | --- | --- |
| `RoutineExample` | `examples/goroutines.go` | Starting a goroutine |
| `ChannelReceiving` | `examples/channels.go` | Sending and receiving values through a channel |
| `WaitGroupTask` | `examples/waitgroups.go` | Waiting for multiple goroutines |
| `WorkerWorks` | `examples/combined.go` | Worker pool, jobs, results, and channel closing |
| `JobHandler` | `examples/producerconsumer.go` | Producer-consumer coordination |

To run a different example, change the function called from `main.go`. For
example:

```go
func main() {
    examples.WorkerWorks()
}
```

Then run:

```bash
go run .
```

## Key ideas demonstrated

- A goroutine runs a function concurrently with the calling code.
- An unbuffered channel synchronizes a sender and receiver.
- A buffered channel can hold values until a receiver is ready.
- Directional channel types document whether a function sends or receives.
- Closing a channel tells receivers that no more values will be sent.
- `sync.WaitGroup` waits for a group of goroutines to finish.
- Results should be collected only after workers have finished sending them.

## Test and format

Run all package tests:

```bash
go test ./...
```

Format all Go source files:

```bash
gofmt -w main.go examples/*.go
```
