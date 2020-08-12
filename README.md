# Jarl (Pronounced Yarl)

Yet Another Retry Library -- with emphasis on "simple"

## Why?
Same as everyone else: I didn't like the implementations I found out there.

I wanted a clean, simple api centered around linear time delays, with support for infinite retries. I may add logarithmic progression at some point, but not today. When I do, infinite retries will throw an error.

## Wait, infinite retries?

Yep, in the event that, say, you're reading from a hot-pluggable hardware device, I don't want the service to bail. Instead, I want it to keep trying to connect, and then carry on its merry way when it does.

However, it's important to be safe, and mitigate memory leaks in what is by design a runaway loop. Because of that, error catching and iteration tracking are both disabled for an infinite retry. Additionally, unrelated to infinite retries, for very large max retries, I limit the number of collected errors to 512, which is probably still an obscene amount to capture.

## TODO:

- Maybe improve the tiny `onRetry` implementation
- Tests
