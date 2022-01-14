---Channels in Golang

# A channel is a medium through which a goroutine communicates with another goroutine. In simple term a channel is a techique which allows to let one goroutine to send data to another goroutine. By default channel is bidirectional. means the goroutines can send or receive data through the same chanel.

# Creating channel
Channel is created using chan keyword. It can only transfer data of same type. Different
types of data are not allowed to transport from the same channel.

var Channel_name chan Type

or 

channel_name := make(chan Type)

//Channels are statical type
	ch := make(chan int)

	ch <- 42  //Writing values to channel

	<-ch  //Reading the value of channel

	//All the channels operations are blocking in default

# When we close the channel, we can't send any data to that channel