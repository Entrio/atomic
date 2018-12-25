package LogManager

import (
	"fmt"
	"time"
)

type (
	LogManager struct {
		startTime *time.Time
	}
)

// Initiate a new instance of log manager and return it
func NewLogManager(t *time.Time) *LogManager {
	lm := &LogManager{
		startTime: t,
	}
	lm.LogfTime("Logger initialized")
	return lm
}

// Output log information with the timestamp of occurrence
func (*LogManager) Logf(format string, params ... interface{}) {
	f := fmt.Sprintf("[INFO] %s > %s", time.Now().Local().Format("15:04:05"), fmt.Sprintf(format, params...))
	fmt.Println(f)
}

// Log current event with the timestamp AND elapsed time since program start
func (l *LogManager) LogfTime(format string, params ... interface{}) {
	f := fmt.Sprintf("[INFO] %s(elapsed %fs) > %s", time.Now().Local().Format("15:04:05"), time.Now().Local().Sub(*l.startTime).Seconds(), fmt.Sprintf(format, params...))
	fmt.Println(f)
}
