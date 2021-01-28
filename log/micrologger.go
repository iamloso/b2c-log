package log

import (
	"fmt"
	microlog "github.com/micro/go-micro/v2/debug/log"
)

// MicroLogger struct
type MicroLogger struct{}

// Print - Log Formatter
// Read reads log entries from the logger
func (o *MicroLogger) Read(...microlog.ReadOption) ([]microlog.Record, error) {
	return nil, nil
}

// Write writes records to log
func (o *MicroLogger) Write(r microlog.Record) error {
	str := fmt.Sprintf("%s %s %v\n", r.Timestamp, r.Metadata, r.Message)
	_, err := Hook().Write([]byte(str))
	return err
}

// Stream log records
func (o *MicroLogger) Stream() (microlog.Stream, error) {
	return nil, nil
}
