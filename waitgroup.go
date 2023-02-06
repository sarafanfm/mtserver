package mtserver

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Job struct {
	Run      func() error
	Shutdown func()
}

func run(onJobError CallbackError, jobs ...*Job) {
	chTerm := make(chan os.Signal)
	chErr := make(chan error)

	wg := new(sync.WaitGroup)

	for i := 0; i < len(jobs); i++ {
		j := jobs[i]
		if j.Run != nil {
			wg.Add(1)
			go func() {
				defer wg.Done()

				if err := j.Run(); err != nil {
					select {
					case chErr <- err:
					default:
					}
				}
			}()
		}
	}

	wg.Add(1)

	go func() {
		defer wg.Done()

		signal.Notify(chTerm, syscall.SIGTERM, syscall.SIGINT)

		select {
		case <-chTerm:
			shutdown(jobs...)
		case err := <-chErr:
			if onJobError != nil {
				onJobError(err)
			}
			shutdown(jobs...)
		}
	}()

	wg.Wait()
}

func shutdown(jobs ...*Job) {
	for i := len(jobs) - 1; i >= 0; i-- {
		if jobs[i].Shutdown != nil {
			jobs[i].Shutdown()
		}
	}
}
