# cqueue

A queue design with golang chan.

## Install

```go get github.com/hanson/cqueue```

## Usage

```go
import github.com/hanson/cqueue

func main() {
    cqueue.Run()
    
    cqueue.AddQueue("task_a", 10, true)
    cqueue.AddQueue("task_b", 10, false)
    
    var conn *conn
    cqueue.AddQueue(cqueue.GenQueueKey(conn), 10, false)
    
	AddTask("task_a", func() {
	    // do something
	})
}
```