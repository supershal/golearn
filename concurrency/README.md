# Go Concurrency

### Go routines and GoMAXPROCS:
https://www.ardanlabs.com/blog/2014/01/concurrency-goroutines-and-gomaxprocs.html
**Process:** The job of the process is to act like a container for all the resources the application uses and maintains as it runs. These resources include things like a memory address space, handles to files, devices and threads.

**Thread:** A thread is a path of execution that is scheduled by the operating system to execute the code we write in our functions against a processor.
A process starts out with one thread, the main thread, and when that thread terminates the process terminates. This is because the main thread is the origin for the application. The main thread can then in turn launch more threads and those threads can launch even more threads.

The operating system schedules a thread to run on an available processor regardless of which process the thread belongs to.

**Goroutines and Parallelism**:
Any function or method in Go can be created as a goroutine. We can consider that the main function is executing as a goroutine, however the Go runtime does not start that goroutine. Goroutines are considered to be lightweight because they use little memory and resources plus their initial stack size is small. 

_The operating system schedules threads to run against available processors and the Go runtime schedules goroutines to run within a logical processor that is bound to a single operating system thread._
By default, the Go runtime allocates a single logical processor to execute all the goroutines that are created for our program

but if you want to run goroutines in parallel, Go provides the ability to add more via the GOMAXPROCS environment variable or runtime function.
runtime.GOMAXPROCS(3)

**Concurrency Vs Parallelism**:
Parallelism is when two or more threads are executing code simultaneously against different processors.
If you configure the runtime to use more than one logical processor, the scheduler will distribute goroutines between these logical processors which will result in goroutines running on different operating system threads.

The problem with building concurrency into our applications is eventually our goroutines are going to attempt to access the same resources, possibly at the same time. Read and write operations against a shared resource must always be atomic

_Channels are the way in Go we write safe and elegant concurrent programs that eliminate race conditions_

### Dectecting race
https://www.ardanlabs.com/blog/2013/09/detecting-race-conditions-with-go.html

```
go build -race
```

### Basic Go sync code
```go
import "sync"

var m sync.Mutex

m.Lock()
m.Unlock()

var wg sync.WaitGroup
wg.Add(2)
go func1(){
  wg.Done()
}
go func2(){
  wg.Done()
}
wg.Wait()
```

### Go channels
https://www.ardanlabs.com/blog/2014/02/the-nature-of-channels-in-go.html
_Channels are type safe message queues that have the intelligence to control the behavior of any goroutine attempting to receive or send on it_
A channel acts as a conduit between two goroutines and will synchronize the exchange of any resource that is passed through it.

 When a channel is created with no capacity, it is called an **unbuffered channel**. In turn, a channel created with capacity is called a **buffered channel.**
 
 **Unbuffered Channel**
 Unbuffered channels have no capacity and therefore require both goroutines to be ready to make any exchange. When a goroutine attempts to send a resource to an unbuffered channel and there is no goroutine waiting to receive the resource, the channel will lock the sending goroutine and make it wait. When a goroutine attempts to receive from an unbuffered channel, and there is no goroutine waiting to send a resource, the channel will lock the receiving goroutine and make it wait.
Synchronization is inherent in the interaction between the send and the receive. One can not happen without the other. The nature of an unbuffered channel is guaranteed synchronization.

**Buffered Channel**
When a goroutine attempts to send a resource to a buffered channel and the channel is full, the channel will lock the goroutine and make it wait until a buffer becomes available. If there is room in the channel, the send can take place immediately and the goroutine can move on. When a goroutine attempts to receive from a buffered channel and the buffered channel is empty, the channel will lock the goroutine and make it wait until a resource has been sent.

 if the buffer is full or if there is nothing to receive, a buffered channel will behave very much like an unbuffered channel.

*create channel*
append `chan` before type. `ch := make(chan int)`. This is bidirectional channel.
to send value to the channel `ch <- val` 
to recive value from channel `val := <-ch`

**bhehvior of channel**
https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html
when it comes to channels, I think about one thing: signaling. A channel allows one goroutine to signal another goroutine about a particular event. Signaling is at the core of everything you should be doing with channels. Thinking of channels as a signaling mechanism will allow you to write better code with well defined and more precise behavior.

To understand how signaling works, we must understand its three attributes:
- Guarantee Of Delivery
- State
- With or Without Data

*Gaurantee of Delivery*:
unbuffered channel gaaurantees deliver. buffered does not

*State*
The state of the channel can be nil, open or closed.

var ch chan string // nil state
ch := make(chan string) // its in open state

close(ch) // closed state

**Signals are sent and received through a channel. Don’t say read/write because channels don’t perform I/O.**

`send` on closed channel causes `panic` .  data can be still received on closed channel.

*with or without data*
When you signal with data, it’s usually because:

  - A goroutine is being asked to start a new task.
  - A goroutine reports back a result.

You signal without data by closing a channel.
When you signal without data, it’s usually because:

  - A goroutine is being told to stop what they are doing.
  - A goroutine reports back they are done with no result.
  - A goroutine reports that it has completed processing and shut down.

*Signaling Without Data*
Signaling without data is mainly reserved for cancellation. It allows one goroutine to signal another goroutine to cancel what they are doing and move on
In most cases you want to use the standard library context package to implement signaling without data. The context package uses an Unbuffered channel underneath for the signaling and the built-in function close to signal without data.

If you choose to use your own channel for cancellation, rather than the context package, your channel should be of type chan struct{}. It is the zero-space, idiomatic way to indicate a channel used only for signalling.

** Iterate over channel **
```go
ch := make(chan int, 5)
go func(){
  for _, p := range ch {
    fmt.Println("receinved", p)
  }
}

const work = 20
for w:=0; w < work; w++{
  select{
    case ch <- "paper":
      fmt.Println("manager send work")
    default:
    //If you can’t perform the send, then you know your box is full and the employee is at capacity. At this point the new work needs to be discarded so things can keep moving.
     fmt.Println("manager: drop")
  }
}
close(ch

```




 





 

