package worker

// A buffered channel that we can send work requests on.
var JobQueue chan Job

type ExecutorInterface interface {
	Handle() error
}

// Job represents the job to be run
type Job struct {
	Executor ExecutorInterface
}

func (j *Job) SetExecutor(exec ExecutorInterface) {
	j.Executor = exec
}
