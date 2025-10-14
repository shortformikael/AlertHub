package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/shortformikael/AlertHub/src/config"
	"github.com/shortformikael/AlertHub/src/utils/colors"
)

type LogType string

const (
	LogTypeInfo    LogType = "INFO"
	LogTypeWarning LogType = "WARNING"
	LogTypeError   LogType = "ERROR"
	LogTypeDebug   LogType = "DEBUG"
	LogTypeFatal   LogType = "FATAL"
)

type LogEntry struct {
	Component    string
	SubComponent string
	LogType      LogType
	Description  string
	Timestamp    string
	PID          string
	TID          string // Thread ID
	GoroutineID  string // Goroutine ID
}

// getThreadID returns the current thread ID (TID)
func getThreadID() string {
	// Lock the current goroutine to an OS thread
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// Get the thread ID using syscall
	tid := syscall.Gettid()
	return strconv.Itoa(tid)
}

// getGoroutineID returns the current goroutine ID
func getGoroutineID() string {
	buf := make([]byte, 64)
	buf = buf[:runtime.Stack(buf, false)]
	// Extract goroutine ID from stack trace
	// Format: "goroutine 123 [running]:"
	for i := 0; i < len(buf); i++ {
		if buf[i] == ' ' {
			// Find the next space or bracket
			for j := i + 1; j < len(buf); j++ {
				if buf[j] == ' ' || buf[j] == '[' {
					return string(buf[i+1 : j])
				}
			}
			break
		}
	}
	return "unknown"
}

func (le *LogEntry) Init(component string, logType LogType, description string, subComponent ...string) {
	le.Component = component
	le.LogType = logType
	le.Description = description
	le.Timestamp = time.Now().Format("15:04:05")
	le.PID = strconv.Itoa(os.Getpid())
	le.TID = getThreadID()
	le.GoroutineID = getGoroutineID()
	// Handle optional subComponent with default value
	if len(subComponent) > 0 && subComponent[0] != "" {
		le.SubComponent = subComponent[0]
	} else {
		le.SubComponent = "" // Default value
	}
}

func (le *LogEntry) PrintString() string {
	name := colors.BrightWhite("[" + config.Name + "]")
	time := colors.BrightBlue("[" + le.Timestamp + "]")
	component := colors.Yellow("[" + le.Component + "]")
	subComponent := ""
	if le.SubComponent != "" {
		subComponent = colors.BrightGreen("[" + le.SubComponent + "]")
	}
	logType := le.getLogTypeString()
	description := le.Description

	if le.SubComponent != "" {
		return fmt.Sprintf("%s %s %s %s %s: %s",
			name,
			time,
			component,
			subComponent,
			logType,
			description)
	}
	return fmt.Sprintf("%s %s %s %s: %s",
		name,
		time,
		component,
		logType,
		description)
}

func (le *LogEntry) getLogTypeString() string {
	switch le.LogType {
	case LogTypeInfo:
		return colors.Green("[" + string(le.LogType) + "]")
	case LogTypeWarning:
		return colors.Yellow("[" + string(le.LogType) + "]")
	case LogTypeError:
		return colors.Red("[" + string(le.LogType) + "]")
	case LogTypeDebug:
		return colors.Cyan("[" + string(le.LogType) + "]")
	case LogTypeFatal:
		return colors.Orange("[" + string(le.LogType) + "]")
	}
	return colors.White("[" + string(le.LogType) + "]")
}

func saveEntry(logEntry LogEntry) {
	filename := "development_" + time.Now().Format("2006-01-02") + ".log"
	filepath := config.Dir.Logs + "/" + filename
	file, openError := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if openError != nil {
		fmt.Println("Error opening file: ", openError)
		return
	}
	defer file.Close()
	jsonData, jsonError := json.Marshal(logEntry)
	if jsonError != nil {
		fmt.Println("Error marshalling log entry: ", jsonError)
		return
	}
	_, writeError := file.WriteString(string(jsonData) + ",\n")
	if writeError != nil {
		fmt.Println("Error writing to file: ", writeError)
		return
	}
}

func Log(component string, logType LogType, description string, subComponent ...string) {
	var le LogEntry
	le.Init(component, logType, description, subComponent...)
	fmt.Println(le.PrintString())
	saveEntry(le)
}
