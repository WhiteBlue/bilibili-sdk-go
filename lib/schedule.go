package lib

import "time"

type Schedule struct {
	t       *time.Ticker
	handler func() bool
}

//Init Schedule
func InitSchedule(delay time.Duration, handler func() bool) *Schedule {
	t := time.NewTicker(delay)

	return &Schedule{t:t, handler:handler}
}


//Start Schedule
func (this *Schedule) Start() {
	for {
		select {
		case <-this.t.C:
			this.handler()
		}
	}
}

//Stop inner timer
func (this *Schedule) Stop() {
	if this.t != nil {
		this.t.Stop()
	}
}