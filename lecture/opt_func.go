package lecture

type Job struct {
	Name  string
	Count int
}

type JobBuilder struct {
	Name  string
	Count int
}

// Design pattern: Builder
type JobHandler func(*JobBuilder)

func WithName(name string) JobHandler {
	// do anything with name...
	return func(builder *JobBuilder) {
		builder.Name = name
	}
}

func WithCount(count int) JobHandler {
	// do anything with count...
	return func(builder *JobBuilder) {
		builder.Count = count
	}
}

func NewJob(handlers ...JobHandler) Job {
	// Array of handlers, thus loop to call each

	builder := new(JobBuilder)
	for _, hdl := range handlers {
		hdl(builder)
	}

	return Job{
		Name:  builder.Name,
		Count: builder.Count,
	}
}

func (j *Job) ProcessName() {

}

func Opt() {
	NewJob(
		WithName("Job 1"),
		WithCount(1),
	)
}
