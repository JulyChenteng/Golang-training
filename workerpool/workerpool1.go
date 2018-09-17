package workerpool

type Job struct{
	Data interface{}
	Proc func(interface{})
}

var JobQueue chan Job

type Worker struct {
	WorkerPool chan (chan Job)
	JobChannel chan Job
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),    // 分配器中，会将任务交给jobChannel，下面会从这里读取到job
}

func (w Worker) Start() {
	go func() {
		for {
            // 将限制的JobChannel（chan） 丢入WorkerPool
			w.WorkerPool <- w.JobChannel

			select {
			case job,ok := <-w.JobChannel:  // 当闲置的jobChannel中有job时，
                job.Payload // job开始工作
                if !ok {
                	return
                }
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}


type Dispatcher struct {
    // 执行池，
	WorkerPool chan chan Job
	maxWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, maxWorkers: maxWorkers}
}

// 初始化worker池，并启动woker池，并开始接受新的job
func (d *Dispatcher) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

// 开始调度，接收新的job
func (d *Dispatcher) dispatch() {
	for {
		select {
		case job, ok := <-JobQueue:
			go func(job Job) {
				fmt.Println("[UPYUN] Dispatcher get JOB")
				
                // 从pool中获取空闲的job channel
				jobChannel := <-d.WorkerPool

                // 将job塞入 job channel中
				jobChannel <- job
			}(job)

			if !ok {
				for v := range d.WorkerPool{
					close(v)
				}
			}
		}
	}
}


