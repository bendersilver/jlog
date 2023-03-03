package jlog

import (
	"bytes"
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path"
	"runtime"
	"runtime/debug"
)

// export JLOGSD=1
var jLogSD = false

var (
	emergLog   *log.Logger
	alertLog   *log.Logger
	critLog    *log.Logger
	errLog     *log.Logger
	warningLog *log.Logger
	noticeLog  *log.Logger
	infoLog    *log.Logger
	debugLog   *log.Logger
)

func write(lg *log.Logger, args ...any) {
	var buf bytes.Buffer
	defer buf.Reset()

	_, fl, line, _ := runtime.Caller(3)
	fmt.Fprintf(&buf, "%s:%s:%d: ", path.Base(path.Dir(fl)), path.Base(fl), line)
	fmt.Fprintln(&buf, args...)
	lg.Printf("%s", buf.Bytes())

}

// Fatal -
func Fatal(a ...any) {
	write(emergLog, a...)
	os.Exit(1)
}

// Fatalf -
func Fatalf(format string, a ...any) {
	write(emergLog, fmt.Sprintf(format, a...))
	os.Exit(1)
}

// Emerg -
func Emerg(a ...any) {
	write(emergLog, a...)
}

// Emergf -
func Emergf(format string, a ...any) {
	write(emergLog, fmt.Sprintf(format, a...))
}

// Alert -
func Alert(a ...any) {
	write(alertLog, a...)
}

// Alertf -
func Alertf(format string, a ...any) {
	write(alertLog, fmt.Sprintf(format, a...))
}

// Crit -
func Crit(a ...any) {
	write(critLog, a...)
}

// Critf -
func Critf(format string, a ...any) {
	write(critLog, fmt.Sprintf(format, a...))
}

// Error -
func Error(a ...any) {
	write(errLog, a...)
}

// Errorf -
func Errorf(format string, a ...any) {
	write(errLog, fmt.Sprintf(format, a...))
}

// Warning -
func Warning(a ...any) {
	write(warningLog, a...)
}

// Warningf -
func Warningf(format string, a ...any) {
	write(warningLog, fmt.Sprintf(format, a...))
}

// Notice -
func Notice(a ...any) {
	write(noticeLog, a...)
}

// Noticef -
func Noticef(format string, a ...any) {
	write(noticeLog, fmt.Sprintf(format, a...))
}

// Info -
func Info(a ...any) {
	write(infoLog, a...)
}

// Infof -
func Infof(format string, a ...any) {
	write(infoLog, fmt.Sprintf(format, a...))
}

// Debug -
func Debug(a ...any) {
	write(debugLog, a...)
}

// Debugf -
func Debugf(format string, a ...any) {
	write(debugLog, fmt.Sprintf(format, a...))
}

// Recover -
func Recover(f func()) {
	defer func() {
		if r := recover(); r != nil {
			Critf("%v\n%s", r, debug.Stack())
		}
	}()
	f()
}

func EmergLogger() *log.Logger {
	return emergLog
}

func AlertLogger() *log.Logger {
	return alertLog
}

func CritLogger() *log.Logger {
	return critLog
}

func ErrLogger() *log.Logger {
	return errLog
}

func WarningLogger() *log.Logger {
	return warningLog
}

func NoticeLogger() *log.Logger {
	return noticeLog
}

func InfoLogger() *log.Logger {
	return infoLog
}

func DebugLogger() *log.Logger {
	return debugLog
}

func init() {
	jLogSD = os.Getenv("JLOGSD") != ""
	var err error
	if jLogSD {
		emergLog, err = syslog.NewLogger(syslog.LOG_EMERG, 0)
		if err != nil {
			fmt.Println(err)
		}

		alertLog, err = syslog.NewLogger(syslog.LOG_ALERT, 0)
		if err != nil {
			fmt.Println(err)
		}

		critLog, err = syslog.NewLogger(syslog.LOG_CRIT, 0)
		if err != nil {
			fmt.Println(err)
		}

		errLog, err = syslog.NewLogger(syslog.LOG_ERR, 0)
		if err != nil {
			fmt.Println(err)
		}

		warningLog, err = syslog.NewLogger(syslog.LOG_WARNING, 0)
		if err != nil {
			fmt.Println(err)
		}

		noticeLog, err = syslog.NewLogger(syslog.LOG_NOTICE, 0)
		if err != nil {
			fmt.Println(err)
		}

		infoLog, err = syslog.NewLogger(syslog.LOG_INFO, 0)
		if err != nil {
			fmt.Println(err)
		}

		debugLog, err = syslog.NewLogger(syslog.LOG_DEBUG, 0)
		if err != nil {
			fmt.Println(err)
		}
	}

	if emergLog == nil {
		emergLog = log.New(os.Stderr, "EMR: ", 0)
	}

	if alertLog == nil {
		alertLog = log.New(os.Stdout, "ALR: ", 0)
	}

	if critLog == nil {
		critLog = log.New(os.Stderr, "CRT: ", 0)
	}

	if errLog == nil {
		errLog = log.New(os.Stderr, "ERR: ", 0)
	}

	if warningLog == nil {
		warningLog = log.New(os.Stderr, "WRN: ", 0)
	}

	if noticeLog == nil {
		noticeLog = log.New(os.Stdout, "NTC: ", 0)
	}

	if infoLog == nil {
		infoLog = log.New(os.Stdout, "INF: ", 0)
	}

	if debugLog == nil {
		debugLog = log.New(os.Stdout, "DBG: ", 0)
	}
}
