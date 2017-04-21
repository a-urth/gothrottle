#### Idea
Implement general purpose throttling library with simple interface in different ways

##### SimpleThrottle
Naive implementation with approach to split time in frames given size and just calculate number of requests for each frame.
`Pros` - easy to implement, works like a charm
`Cons` - not that accurate

##### ChannelThrottler
Sort of (leaky bucket)[https://en.wikipedia.org/wiki/Leaky_bucket] implementation using go' channel and ticker
`Pros` - accurate and still simple and easy to implement
`Cons` - not good that we'll just live as it is ticker and goroutine which works with it

##### TODO
Throttler using linked list or array of records with times
`Pros` - no concurrency means more simplicity
`Cons` - no concurrency means overhead on manual calculation which records should be removed
