package main

import "os"

type Logger struct {
	Path string
	file *os.File
}

func (l *Logger) Write(content string) {
	if l.file == nil {
		var err error
		l.file, err = os.OpenFile("tmp.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return
		}
	}
	l.file.WriteString(content)
}