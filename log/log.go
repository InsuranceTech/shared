package log

import (
	"fmt"
	"time"
)

func write(mtype string, text string) {
	t := time.Now()
	fmt.Println("[" + t.Format("2006-01-02 15:04:05") + "] [" + mtype + "] : " + text)
}

func InfoF(format string, a ...any) {
	text := fmt.Sprintf(format, a)
	write("INFO", text)
}

func Info(a ...any) {
	text := fmt.Sprint(a)
	write("INFO", text)
}

func WarningF(format string, a ...any) {
	text := fmt.Sprintf(format, a)
	write("WARNING", text)
}

func Warning(a ...any) {
	text := fmt.Sprint(a)
	write("WARNING", text)
}

func ErrorF(format string, a ...any) {
	text := fmt.Sprintf(format, a)
	write("ERROR", text)
}

func Error(a ...any) {
	text := fmt.Sprint(a)
	write("ERROR", text)
}

func LogF(format string, a ...any) {
	text := fmt.Sprintf(format, a)
	write("LOG", text)
}

func Log(a ...any) {
	text := fmt.Sprint(a)
	write("LOG", text)
}
