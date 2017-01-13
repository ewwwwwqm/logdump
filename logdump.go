package logdump

import (
	"fmt"
	"log"
	"os"
	"io"
)

type LogsOutput struct {
	File []FileOutput
}
type FileOutput struct {
	Logtype string
	Logoutput io.Writer
}

var logsOutput LogsOutput

func Init(logstruct map[string]interface{}) {
	for k, v := range logstruct {
		f, err := os.OpenFile(fmt.Sprintf("%v", v), os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
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
