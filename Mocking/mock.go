package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go"
	countDownStart = 3
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}
type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
