package logging

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/davecgh/go-spew/spew"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
	mu  sync.Mutex
)

type LogType int

const (
	LogTypeStandard LogType = iota
	LogTypeSimple
	LogTypeColored
	LogTypeJSON
)

func init() {
	fmt := &logrus.TextFormatter{
		ForceColors:      true,
		DisableSorting:   true,
		DisableTimestamp: true,
	}

	log.Out = os.Stdout
	log.Formatter = fmt
}

func Output() io.Writer {
	mu.Lock()
	defer mu.Unlock()
	return log.Out
}

func SetOutput(w io.Writer) {
	mu.Lock()
	defer mu.Unlock()
	log.Out = w
}

func SetDebug(enabled bool) {
	mu.Lock()
	defer mu.Unlock()

	if enabled {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}
}

func SetLogTypeFromString(typ string) {
	switch typ {
	case "standard":
		SetLogType(LogTypeStandard)
		break
	case "simple":
		SetLogType(LogTypeSimple)
		break
	case "colored":
		SetLogType(LogTypeColored)
		break
	case "json":
		SetLogType(LogTypeJSON)
		break
	default:
		Warn("unknown log type - set colored")
		SetLogType(LogTypeColored)
	}
}

func SetLogType(typ LogType) {
	mu.Lock()
	defer mu.Unlock()

	switch typ {
	case LogTypeStandard:
		log.Formatter = new(logrus.TextFormatter)
		break

	case LogTypeSimple:
		log.Formatter = &logrus.TextFormatter{
			DisableColors:    true,
			DisableSorting:   true,
			DisableTimestamp: true,
		}
		break

	case LogTypeColored:
		log.Formatter = &logrus.TextFormatter{
			ForceColors:      true,
			DisableSorting:   true,
			DisableTimestamp: true,
		}
		break

	case LogTypeJSON:
		log.Formatter = new(logrus.JSONFormatter)
		break
	}
}

func LogScriptError(name string, msg interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Errorf("script error [%s]: %v", name, msg)
}

func LogScriptWarn(name string, msg interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Warnf("script warn [%s]: %v", name, msg)
}

func Printf(format string, a ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Fprintf(log.Out, format, a...)
}

func Println(a ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Fprintln(log.Out, a...)
}

func Print(a ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Fprint(log.Out, a...)
}

func DDumpUnmarshaled(descr string, in []byte) {
	mu.Lock()
	defer mu.Unlock()

	if log.Level < logrus.DebugLevel {
		return
	}

	var res interface{}
	if err := ffjson.Unmarshal(in, &res); err != nil {
		panic("DumpUnmarshaled: unable to unmarshal input")
	}

	fmt.Fprintf(log.Out, "%s ------------------------- dump start ---------------------------------------\n", descr)
	spew.Fdump(log.Out, res)
	fmt.Fprintf(log.Out, "%s -------------------------  dump end  ---------------------------------------\n\n", descr)
}

func DDumpJSON(descr string, in interface{}) {
	mu.Lock()
	defer mu.Unlock()

	if log.Level < logrus.DebugLevel {
		return
	}

	out, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		panic("DumpJSON: unable to marshal input")
	}

	fmt.Fprintf(log.Out, "%s ------------------------- dump start ---------------------------------------\n", descr)
	fmt.Fprintln(log.Out, string(out))
	fmt.Fprintf(log.Out, "%s -------------------------  dump end  ---------------------------------------\n\n", descr)
}

func DumpJSON(descr string, in interface{}) {
	mu.Lock()
	defer mu.Unlock()

	out, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		panic("DumpJSON: unable to marshal input")
	}

	fmt.Fprintf(log.Out, "%s ------------------------- dump start ---------------------------------------\n", descr)
	fmt.Fprintln(log.Out, string(out))
	fmt.Fprintf(log.Out, "%s -------------------------  dump end  ---------------------------------------\n\n", descr)
}

func DDump(descr string, in interface{}) {
	mu.Lock()
	defer mu.Unlock()

	if log.Level < logrus.DebugLevel {
		return
	}

	fmt.Fprintf(log.Out, "%s ------------------------- dump start ---------------------------------------\n", descr)
	spew.Fdump(log.Out, in)
	fmt.Fprintf(log.Out, "%s -------------------------  dump end  ---------------------------------------\n\n", descr)
}

func Dump(descr string, in interface{}) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Fprintf(log.Out, "%s ------------------------- dump start ---------------------------------------\n", descr)
	spew.Fdump(log.Out, in)
	fmt.Fprintf(log.Out, "%s -------------------------  dump end  ---------------------------------------\n\n", descr)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Debug(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Warn(args...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Warning(args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Error(args...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Fatal(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Debugf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Warnf(format, args...)
}

// Warningf logs a message at level Warn on the standard logger.
func Warningf(format string, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Warningf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger.
func Fatalf(format string, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Fatalf(format, args...)
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Debugln(args...)
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Infoln(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Warnln(args...)
}

// Warningln logs a message at level Warn on the standard logger.
func Warningln(args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Warningln(args...)
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Errorln(args...)
}

// Panicln logs a message at level Panic on the standard logger.
func Panicln(args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Panicln(args...)
}

// Fatalln logs a message at level Fatal on the standard logger.
func Fatalln(args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	log.Fatalln(args...)
}
