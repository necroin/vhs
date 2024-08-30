package logger

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"time"
)

func GetLogLevel(level string) int {
	result, ok := levelMap[level]
	if !ok {
		result = NoneLevel
	}
	return result
}

type Logger struct {
	out   io.Writer
	level int
	mutex *sync.Mutex
}

type LogEntry struct {
	logger       *Logger
	globalLabels []string
	localLabels  []string
	globalPrefix string
	localPrefix  string
}

func NewLogger(out io.Writer, level int) *Logger {
	return &Logger{
		out:   out,
		level: level,
		mutex: &sync.Mutex{},
	}
}

func (logger *Logger) WtihLabels(labels ...string) *LogEntry {
	prefixes := []string{}
	for _, label := range labels {
		prefix := fmt.Sprintf("[%s]", label)
		prefixes = append(prefixes, prefix)
	}

	return &LogEntry{
		logger:       logger,
		globalLabels: labels,
		localLabels:  labels,
		globalPrefix: strings.Join(prefixes, " "),
		localPrefix:  strings.Join(prefixes, " "),
	}
}

func (log *LogEntry) WtihLabels(labels ...string) *LogEntry {
	globalLabels := append(log.globalLabels[:], labels...)

	globalPrefixes := []string{}
	for _, label := range globalLabels {
		prefix := fmt.Sprintf("[%s]", label)
		globalPrefixes = append(globalPrefixes, prefix)
	}

	localPrefixes := []string{}
	for _, label := range labels {
		prefix := fmt.Sprintf("[%s]", label)
		localPrefixes = append(localPrefixes, prefix)
	}

	return &LogEntry{
		logger:       log.logger,
		globalLabels: globalLabels,
		localLabels:  labels,
		globalPrefix: strings.Join(globalPrefixes, " "),
		localPrefix:  strings.Join(localPrefixes, " "),
	}
}

func (log *LogEntry) Print(level string, message string, args ...any) {
	log.logger.mutex.Lock()
	defer log.logger.mutex.Unlock()
	fmt.Fprintf(log.logger.out, "%s %s: %s %s\n", time.Now().Format("2006-01-02 15:04:05"), level, log.globalPrefix, fmt.Sprintf(message, args...))
}

func (log *LogEntry) Error(message string, args ...any) {
	if log.logger.level >= ErrorLevel {
		log.Print("ERROR", message, args...)
	}
}

func (log *LogEntry) Info(message string, args ...any) {
	if log.logger.level >= InfoLevel {
		log.Print("INFO", message, args...)
	}
}

func (log *LogEntry) Verbose(message string, args ...any) {
	if log.logger.level >= VerboseLevel {
		log.Print("VERBOSE", message, args...)
	}
}

func (log *LogEntry) Debug(message string, args ...any) {
	if log.logger.level >= DebugLevel {
		log.Print("DEBUG", message, args...)
	}
}

func (log *LogEntry) Detailed(message string, args ...any) {
	if log.logger.level >= DetailedLevel {
		log.Print("DETAILED", message, args...)
	}
}

func (entry *LogEntry) NewError(message string, args ...any) error {
	message = fmt.Sprintf(message, args...)
	return fmt.Errorf("%s [Error] %s", entry.localPrefix, message)
}
