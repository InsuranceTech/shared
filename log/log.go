package log

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var (
	lock    sync.Locker
	counter int64
)

type TaggedLogger struct {
	Tag string
}

func write(mtype string, text string, tag string) {
	t := time.Now()
	if tag == "" {
		fmt.Println("[" + t.Format("2006-01-02 15:04:05") + "] [" + mtype + "] : " + text)
	} else {
		fmt.Println("[" + t.Format("2006-01-02 15:04:05") + "] [" + mtype + "] [" + tag + "]: " + text)
	}
}

func CreateTag(tagName string) *TaggedLogger {
	atomic.AddInt64(&counter, 1)
	return &TaggedLogger{
		Tag: tagName + "#" + strconv.Itoa(int(counter)),
	}
}

func CreateTagID(tagName string) *TaggedLogger {
	return &TaggedLogger{
		Tag: tagName,
	}
}

// region Taglı kullanım
func (l *TaggedLogger) InfoF(format string, a ...any) {
	text := fmt.Sprintf(format, a)
	write("INFO", text, l.Tag)
}

func (l *TaggedLogger) Info(a ...any) {
	text := fmt.Sprint(a)
	write("INFO", text, l.Tag)
}

func (l *TaggedLogger) WarningF(format string, a ...any) {
	text := fmt.Sprintf(format, a)
	write("WARNING", text, l.Tag)
}

func (l *TaggedLogger) Warning(a ...any) {
	text := fmt.Sprint(a)
	write("WARNING", text, l.Tag)
}

func (l *TaggedLogger) ErrorF(format string, a ...any) {
	text := fmt.Sprintf(format, a)
	write("ERROR", text, l.Tag)
}

func (l *TaggedLogger) Error(a ...any) {
	text := fmt.Sprint(a)
	write("ERROR", text, l.Tag)
}

func (l *TaggedLogger) LogF(format string, a ...any) {
	text := fmt.Sprintf(format, a)
	write("LOG", text, l.Tag)
}

func (l *TaggedLogger) Log(a ...any) {
	text := fmt.Sprint(a)
	write("LOG", text, l.Tag)
}

func (l *TaggedLogger) FatalF(format string, a ...any) {
	text := fmt.Sprintf(format, a)
	write("FATAL", text, l.Tag)
	os.Exit(1)
}

func (l *TaggedLogger) Fatal(a ...any) {
	text := fmt.Sprint(a)
	write("FATAL", text, l.Tag)
	os.Exit(1)
}

//endregion

// region Genel kullanım
func InfoF(format string, a ...any) {
	text := fmt.Sprintf(format, a)
	write("INFO", text, "")
}

func Info(a ...any) {
	text := fmt.Sprint(a)
	write("INFO", text, "")
}

func WarningF(format string, a ...any) {
	text := fmt.Sprintf(format, a)
	write("WARNING", text, "")
}

func Warning(a ...any) {
	text := fmt.Sprint(a)
	write("WARNING", text, "")
}

func ErrorF(format string, a ...any) {
	text := fmt.Sprintf(format, a)
	write("ERROR", text, "")
}

func Error(a ...any) {
	text := fmt.Sprint(a)
	write("ERROR", text, "")
}

func LogF(format string, a ...any) {
	text := fmt.Sprintf(format, a)
	write("LOG", text, "")
}

func Log(a ...any) {
	text := fmt.Sprint(a)
	write("LOG", text, "")
}

func FatalF(format string, a ...any) {
	text := fmt.Sprintf(format, a)
	write("FATAL", text, "")
	os.Exit(1)
}

func Fatal(a ...any) {
	text := fmt.Sprint(a)
	write("FATAL", text, "")
	os.Exit(1)
}

//endregion
