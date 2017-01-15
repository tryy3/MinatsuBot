package minatsubot

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const (
	// ErrorLevel level. Used for errors that requires an extra eye on.
	// Mostly used for errors like, not being able to write to a cricial file or similiar.
	// Can also be used before crashing.
	ErrorLevel Level = iota
	// WarnLevel level. Non-cricial messages, used to warn about things,
	// but not bad enough to crash the app or similiar.
	WarnLevel
	// InfoLevel level. Used for most general stuff, used to tell what the application
	// is doing, like starting up and such.
	InfoLevel
	// DebugLevel level. Used for debugging stuff, often very verbose messages.
	// Usually used for debugging.
	DebugLevel
)

type Level uint8

// GetLoggingLevel expects a string and returns a Level depending on the string.
// Always gonna return InfoLevel if the string is not a correct level.
func getLoggingLevel(level string) Level {
	switch strings.ToLower(level) {
	case "error":
		return ErrorLevel
	case "warn":
	case "warning":
		return WarnLevel
	case "debug":
		return DebugLevel
	}
	return InfoLevel
}

func newLogger(name string, level Level) *Logger {
	return &Logger{
		writers: []io.Writer{os.Stdout},
		prefix:  name,
		level:   level,
	}
}

type Logger struct {
	writers []io.Writer
	prefix  string
	level   Level
}

func (l *Logger) formatPrefix(level string, message string) string {
	date := time.Now()
	year, month, day := date.Date()
	hour, min, sec := date.Clock()
	return fmt.Sprintf("%d/%d/%d %d:%d:%d %s %s: %s", year, month, day, hour, min, sec, level, l.prefix, message)
}

func (l *Logger) log(message string) {
	for _, writer := range l.writers {
		fmt.Fprintln(writer, message)
	}
}

// Format

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.level >= DebugLevel {
		l.log(l.formatPrefix("DEBUG", fmt.Sprintf(format, args...)))
	}
}

func (l *Logger) Infof(format string, args ...interface{}) {
	if l.level >= InfoLevel {
		l.log(l.formatPrefix("INFO", fmt.Sprintf(format, args...)))
	}
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	if l.level >= WarnLevel {
		l.log(l.formatPrefix("WARNING", fmt.Sprintf(format, args...)))
	}
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	if l.level >= WarnLevel {
		l.log(l.formatPrefix("WARNING", fmt.Sprintf(format, args...)))
	}
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	if l.level >= ErrorLevel {
		l.log(l.formatPrefix("ERROR", fmt.Sprintf(format, args...)))
	}
}

// No format

func (l *Logger) Debug(args ...interface{}) {
	if l.level >= DebugLevel {
		l.log(l.formatPrefix("DEBUG", fmt.Sprint(args...)))
	}
}

func (l *Logger) Info(args ...interface{}) {
	if l.level >= InfoLevel {
		l.log(l.formatPrefix("INFO", fmt.Sprint(args...)))
	}
}

func (l *Logger) Warn(args ...interface{}) {
	if l.level >= WarnLevel {
		l.log(l.formatPrefix("WARNING", fmt.Sprint(args...)))
	}
}

func (l *Logger) Warning(args ...interface{}) {
	if l.level >= WarnLevel {
		l.log(l.formatPrefix("WARNING", fmt.Sprint(args...)))
	}
}

func (l *Logger) Error(args ...interface{}) {
	if l.level >= ErrorLevel {
		l.log(l.formatPrefix("ERROR", fmt.Sprint(args...)))
	}
}
