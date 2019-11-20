# lockedsource

In the Go standard library, the `rand.Source` returned by `rand.NewSource(seed int64)`  is not concurrency safe. This project provides a concurrent-safe `rand.Source`.

### Usage

Relace`rand.Source(seed int64)` with `lockedsource.New(seed int64)`.
