package timer

import (
	"fmt"
	"go-debug/output"
	"time"
)

type Timer struct {
	Name      string
	StartTime *time.Time
	EndTime   *time.Time
	Since     time.Duration
}

func NewTimer(s ...string) *Timer {

	if len(s) > 0 {
		return &Timer{Name: s[0]}
	}
	return &Timer{}
}

func (t *Timer) Start() {
	start := time.Now()
	t.StartTime = &start
}
func (t *Timer) Stop() {
	end := time.Now()
	t.EndTime = &end
	t.Since = t.EndTime.Sub(*t.StartTime)

	var msg string
	if t.Name != "" {
		// If a name is set, include it in the message
		msg = fmt.Sprintf("%s time elapsed: %v", t.Name, t.Since)
	} else {
		// If no name is set, output the time elapsed only
		msg = fmt.Sprintf("Time elapsed: %v", t.Since)
	}
	output.Info(msg)
}
func (t *Timer) TimeFunction() func() {
	return func() {
		t.Start()
		defer t.Stop()
		output.Infof("Time elapsed: %v\n", t.Since)

	}
}

func MesureRestOfFunc() {
	start := time.Now()
	defer func() {
		output.Infof("Time elapsed: %v\n", time.Since(start))
	}()
}

func Time(s ...string) func() {

	if len(s) == 1 {
		return func() {
			t := NewTimer(s[0])
			t.Start()
			defer t.Stop()
		}
	} else {
		return func() {
			t := NewTimer()
			t.Start()
			defer t.Stop()
		}
	}

}
