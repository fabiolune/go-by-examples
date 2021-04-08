# Buffered channels and semaphore pattern

In this example there are 100 http calls to some external apis. They could potentially run simultaneosly, but threads (even gree nthreads) can be expensive.
Then we introduce a semaphore to allow only a certain amount (10) of concurrent goroutines.
The semaphore is implemented using a buffered channel (channel with only a limited amount of possible slots): the acquisition of the right of running the goroutine is based on the fact that the blocking operation of pushing items to the channel is only possible when some slots are free.
Slots are released as soon as a goroutine executes by receiving one of the values available on the channel.

The reason why the program prints out more than the specified amount of concurrent operations (which is 10) depends on the fact that other routines are actually running that are not directly launched by the program (the `main` routine, the garbage collector routine, ...)

# References

based on [Go (Golang) Semaphore Pattern Tutorial](https://www.youtube.com/watch?v=yayBIP1v0z8)