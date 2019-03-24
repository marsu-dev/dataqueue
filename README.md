# dataqueue
N-M producer/consumer in go

This package implement a cancelable N-M producer/consumer model.

## usage
see [Exemples](exemples/)

```go
    func main() {
        ctx := context.Background()
        ctx, ctxCancel := context.WithTimeout(ctx, 1500*time.Millisecond)
        defer ctxCancel()

        startTime := time.Now()

        dataqueue.StartSimpleConsumer(ctx,
            &ChalengeFinder{
                difficulty: 1 << 30,
                chalenge:   1000,
            },
            runtime.NumCPU()*4,
        )

        log.Printf("Execution time: %s", time.Now().Sub(startTime).Round(time.Millisecond))
    }
```
