// Package for multi-level logging.
package logdump

import (
	"fmt"
	"log"
	"os"
	"io"
)

// LogsOutput represents an active logging objects
// that generates lines of output to an io.Writer.
type LogsOutput struct {
	File []FileOutput // contains log type and output writer
}

// FileOutput contains log level type and io.Writer for output.
// Each logging operation makes a single call to 
// the Writer's Write method.
type FileOutput struct {
	Logtype string // contains log level type
	Logoutput io.Writer // contains io.Writer for output
}

var logsOutput LogsOutput

// Inits log structure and assigns output.
// FileOutput.Logoutput points to io.Writer
// depending on the specified level.
// After init log output sets to default os.Stderr.
func Init(logstruct map[string]interface{}) {
	for k, v := range logstruct {
		f, err := os.OpenFile(fmt.Sprintf("%v", v), 
			os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		if err != nil {
		    log.Fatalf("Error opening log file: %v", err)
		}
		fileOutput := FileOutput{Logtype: k, Logoutput: f}
		logsOutput.File = append(logsOutput.File, fileOutput)
		log.SetOutput(fileOutput.Logoutput)
		log.Println("Start", k, "logging")
	}
	log.SetOutput(os.Stderr)
}

// Write dumps message to the specified level.
// On error prints message to default os.Stderr.
func Write(logtype string, message string) {
	success := false
	for _, v := range logsOutput.File {
		if v.Logtype == logtype {
			log.SetOutput(v.Logoutput)
			log.Println(message)
			success = true
		}
	}
	log.SetOutput(os.Stderr)
	if !success {
		log.Println("Type", logtype, "was not assigned for log output")
	}
}
