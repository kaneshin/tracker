package tracker

import (
	"log"
	"os"
	"time"
)

var (
	Logger *log.Logger
	Format func(time.Duration) string
	Hook   func(time.Duration)
)

func init() {
	// Initialize default logger to print stdout.
	Logger = log.New(os.Stdout, "", 0)

	// A representing of the duration form is "72h3m0.5s".
	Format = func(d time.Duration) string {
		return d.String()
	}
}

// Time calculates and prints elapsed time by given time.
func Time(s time.Time, args ...interface{}) {
	elapsed := time.Since(s)
	if Format != nil {
		args = append(args, Format(elapsed))
	} else {
		args = append(args, elapsed.String())
	}
	if Logger != nil {
		Logger.Println(args...)
	}
	if Hook != nil {
		Hook(elapsed)
	}
}
