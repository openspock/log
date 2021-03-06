package log

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Disabled enables toggling logging on and off.
var Disabled = false

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

// Type is an enumeration that identifies a log type.
type Type int

const (
	// AppLog identifies an application function log
	AppLog Type = iota
	// SysLog identifies a system log, for e.g. HeartBeat.
	SysLog
	// SysMsg identifies a system message, including but not limited to heartbeats.
	SysMsg
	// AppMsg identifies an application message.
	AppMsg
)

func (t Type) String() string {
	return [...]string{"app_log", "sys_log", "sys_msg", "app_msg"}[t]
}

// Info logs information logs to configured logger
func Info(msg string, t Type, params map[string]interface{}) {
	if !Disabled {
		log.WithFields(params).WithFields(log.Fields{"type": t.String()}).Info(msg)
	}
}

// Warn logs warning logs to configured logger
func Warn(msg string, t Type, params map[string]interface{}) {
	if !Disabled {
		log.WithFields(params).WithFields(log.Fields{"type": t.String()}).Warn()
	}
}

// Error logs error logs to configured logger
func Error(msg string, t Type, params map[string]interface{}) {
	if !Disabled {
		log.WithFields(params).WithFields(log.Fields{"type": t.String()}).Fatal()
	}
}

// Crit logs critical logs to configured logger
func Crit(msg string, t Type, params map[string]interface{}) {
	if !Disabled {
		log.WithFields(params).WithFields(log.Fields{"type": t.String()}).Panic()
	}
}
