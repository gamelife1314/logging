package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// logging.MessageRecord
type MessageRecord struct {
	Level         MessageLevel
	LevelString   string
	Message       string
	Pid           int
	Program       string
	Time          string
	FuncName      string
	LongFileName  string
	ShortFileName string
	Line          int
	Color         string
	ColorClear    string
}

// logging.getMessageRecord, make a record and return it's reference.
func GetMessageRecord(level MessageLevel, format string, a ...interface{}) *MessageRecord {
	message := fmt.Sprintf(format, a...)
	pc, file, line, _ := runtime.Caller(3)
	record := &MessageRecord{
		Level:         level,
		Message:       message,
		Pid:           os.Getpid(),
		Program:       filepath.Base(os.Args[0]),
		Time:          "",
		FuncName:      runtime.FuncForPC(pc).Name(),
		LongFileName:  file,
		ShortFileName: filepath.Base(file),
		Line:          line,
		Color:         LevelColorFlag[level],
		ColorClear:    LevelColorSeqClear,
		LevelString:   LevelString[level],
	}
	return record
}
