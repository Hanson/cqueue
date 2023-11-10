package cqueue

import (
	"fmt"
	"sync"
)

type Queue struct {
	task    chan *Task
	stop    chan struct{}
	routine bool
}

type QueueList struct {
	m map[string]*Queue
	sync.RWMutex
}

var queueList *QueueList

func AddQueue(key string, concurrent int, routine bool) {
	queueList.Lock()
	defer queueList.Unlock()

	q := &Queue{}
	q.routine = routine
	q.task = make(chan *Task, concurrent)
	q.stop = make(chan struct{})
	queueList.m[key] = q
	q.RunQueue()
}

func DelQueue(key string) {
	queueList.Lock()
	defer queueList.Unlock()

	if q, ok := queueList.m[key]; ok {
		q.stop <- struct{}{}
	}
}

func (q Queue) RunQueue() {
	go func() {
		for {
			select {
			case t := <-q.task:
				if q.routine {
					go t.f()
				} else {
					t.f()
				}
			case <-q.stop:
				return
			}
		}
	}()
}

func Run() {
	queueList = &QueueList{}
	queueList.m = make(map[string]*Queue)
}

func GenQueueKey(obj interface{}) string {
	return fmt.Sprintf("%p", obj)
}
