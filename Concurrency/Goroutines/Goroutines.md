---Goroutines

1. Goroutines are cheaper than threads

2. Goroutines are stored in the stack and the size of the stack can grow and shrink     according to the requirement of the program. But In the threads, the size of the stack is fixed.

3. Goroutines can communicate using the channel.

4. Due to the presence of channel, one go routine can communicate with other goroutine with low latency. 