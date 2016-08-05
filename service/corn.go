package service

import (
	"time"
	"github.com/anacrolix/sync"
	"github.com/go-playground/log"
)

type CornTaskImpl interface {
	Run() error
	Success()
	Failure(error)
	GetName() string
	GetDuration() time.Duration
	GetLastRun() time.Time
	SyncLastRunTime()
}

type CornTask struct {
	Name     string
	Duration time.Duration
	LastRun  time.Time
}

func (t *CornTask) Run() error {
	return nil
}

func (t *CornTask) Success() {

}

func (t *CornTask) Failure(err error) {

}

func (t *CornTask) GetName() string {
	return t.Name
}

func (t *CornTask) GetDuration() time.Duration {
	return t.Duration
}

func (t *CornTask) GetLastRun() time.Time {
	return t.LastRun
}

func (t*CornTask) SyncLastRunTime() {
	t.LastRun = time.Now()
}


//execute task
func exec(f CornTaskImpl) {
	log.Info("invoke task, taskname: ", f.GetName())
	defer func() {
		if r := recover(); r != nil {
			log.Error(r)
		}
	}()

	if err := f.Run(); err != nil {
		f.Failure(err)
	} else {
		f.Success()
	}
	log.Info("run task end, taskname: ", f.GetName())
}

type CornService struct {
	ticker *time.Ticker
	tasks  []CornTaskImpl
	lock   sync.Mutex
	done   chan struct{}
}

func (c *CornService) RegisterTask(task CornTaskImpl) {
	task.SyncLastRunTime()
	go exec(task)
	c.tasks = append(c.tasks, task)
}

func (c *CornService) syncTaskList(nowTime time.Time) {
	for _, task := range c.tasks {
		//Unix timestamp => duration
		between := time.Duration(nowTime.Unix() - task.GetLastRun().Unix()) * time.Second
		if between >= task.GetDuration() {
			task.SyncLastRunTime()
			exec(task)
		}
	}
}

func (c *CornService) loop() {
	for {
		select {
		case <-c.done:
			log.Info("corn loop stopped....")
			return
		case nowTime := <-c.ticker.C:
			go c.syncTaskList(nowTime)
		}
	}
}

func (c *CornService) Start() {
	go c.loop()
}

func (c *CornService) Stop() {
	c.ticker.Stop()
	close(c.done)
}

func NewCornService() *CornService {
	return &CornService{
		ticker: time.NewTicker(time.Minute),
		tasks:[]CornTaskImpl{},
		lock:sync.Mutex{},
		done:make(chan struct{}, 1),
	}
}