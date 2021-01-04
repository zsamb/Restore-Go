package logger

import (
	"fmt"
	"os"
	"time"
)

type (
	Logger struct {
		accessPath, activityPath, systemPath string
		Debug                                bool
		queue                                chan LogItem
	}
	LogItem   interface{}
	AccessLog struct {
		method     string
		source     string
		location   string
		showAuth   bool
		authUser   string
		debug      bool
		accessPath string
	}
	SystemLog struct {
		error      bool
		message    string
		systemPath string
	}
	ActivityLog struct {
		activity     string
		message      string
		activityPath string
	}
)

func (logger *Logger) beginQueue(queue <-chan LogItem) {
	for job := range queue {
		switch job.(type) {
		case AccessLog:
			accessLog, ok := job.(AccessLog)
			if ok {
				accessLog.Send()
			}
		case SystemLog:
			systemLog, ok := job.(SystemLog)
			if ok {
				systemLog.Send()
			}
		case ActivityLog:
			activityLog, ok := job.(ActivityLog)
			if ok {
				activityLog.Send()
			}
		}
	}
}

func (logger *Logger) Access(method, source, location, authUser string, showAuth bool) {
	select {
	case logger.queue <- AccessLog{
		method:     method,
		source:     source,
		location:   location,
		showAuth:   showAuth,
		authUser:   authUser,
		debug:      logger.Debug,
		accessPath: logger.accessPath,
	}:
	default:
	}
}

func (logger *Logger) System(message string, error bool) {
	select {
	case logger.queue <- SystemLog{
		error:      error,
		message:    message,
		systemPath: logger.systemPath,
	}:
	default:
	}
}

func (logger *Logger) Activity(activity, message string) {
	select {
	case logger.queue <- ActivityLog{
		activity:     activity,
		message:      message,
		activityPath: logger.activityPath,
	}:
	default:
	}
}

//Check logging locations are available
func (logger *Logger) Validate() error {
	paths := []string{logger.activityPath, logger.accessPath, logger.systemPath}
	present := time.Now()
	for _, path := range paths {

		file, err := os.OpenFile(path+present.Format("02-01-2006.log"), os.O_CREATE|os.O_RDONLY, 0644)
		if err != nil {
			return fmt.Errorf("could not open log files: %s", err)
		}
		_ = file.Close()

	}
	return nil
}

//Create Initiates new logger
func Create(accessPath, activityPath, systemPath string, debug bool) *Logger {
	queueChannel := make(chan LogItem, 100)
	instance := Logger{
		accessPath:   accessPath,
		activityPath: activityPath,
		systemPath:   systemPath,
		Debug:        debug,
		queue:        queueChannel,
	}
	go instance.beginQueue(queueChannel)
	return &instance
}
