package cqueue

type Task struct {
	f func()
}

func AddTask(queue string, f func()) {
	queueList.RLock()
	defer queueList.RUnlock()

	if q, ok := queueList.m[queue]; ok {
		q.task <- &Task{
			f: f,
		}
	}
}
