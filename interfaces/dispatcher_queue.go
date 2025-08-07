package jobqueue


type JobQueue interface {
	Enqueue() error
	Dequeue() error 
	Ack() error
}

