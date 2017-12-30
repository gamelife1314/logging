package logging

import (
	"bytes"
	"strings"
	"text/template"
	"time"
)

// logging.MessageFormatter
type MessageFormatter struct {

	// Message Format, you can use any fields of MessageRecord.
	// Example: {{.Color}}[{{.Time}}] {{.LevelString}}  {{.FuncName}} {{.ShortFileName}} {{.Line}} {{.ColorClear}} {{.Message}}\n
	Format string

	// Message Time Format
	// Default: time.RFC1123
	TimeFormat string
}

// logging.MessageFormatter.GetMessage, return formatted message string for output.
func (formatter *MessageFormatter) GetMessage(logger *Logger) string {
	if formatter.TimeFormat == "" {
		formatter.TimeFormat = time.RFC1123
	}
	logger.Record.Time = time.Now().Format(formatter.TimeFormat)
	stringBuffer := new(bytes.Buffer)
	tpl := template.Must(template.New("messageFormat").Parse(formatter.Format))
	tpl.Execute(stringBuffer, *logger.Record)
	message := stringBuffer.String()
	if strings.Index(message, "\n") != len(message)-1 {
		message += "\n"
	}
	return message
}
