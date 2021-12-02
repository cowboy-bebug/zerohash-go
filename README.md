> 📢 Please check out my Python implementation as well: [cowboy-bebug/zerohash-py]([https://github.com/cowboy-bebug/zerohash-py])

---

## To get started:

Since we don't have to build a binary, simply run:

```console
$ git clone git@github.com:cowboy-bebug/zerohash-go.git
$ cd zerohash-go
$ go mod download
$ go run main.go
```

To test:

```console
$ go test -v ./...
```

To format:

```console
$ go fmt ./...
```

To lint:

```console
$ go vet ./...
```

## Design Considerations

- A simple `producer -> (channel) -> consumer` model is used for each trading pair
  - a dedicated channel is used for each pair of consumer + producer goroutines to achieve isolation
- Some notes on functions:
  - `Consume()`
    - assumed it's a standard practice in Go to pass a structure by reference and mutate its values
    - so that's the pattern used (`Price` is passed by reference)
  - `Produce()`
    - `panic` is used crash the program if we can't establish websocket subscription (as it's likely that we made a mistake)
  - `computeVwap()`
    - logic to compute volume-weighted average price is extracted into this function for unit testing purposes

## Assumptions

Assumptions remain the same as mentioned in [zerohash-py#assumptions](https://github.com/cowboy-bebug/zerohash-py#assumptions)
