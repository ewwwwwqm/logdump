package logdump_test

import (
	"fmt"

    "github.com/ewwwwwqm/logdump"
)

func ExampleInit() {
	newLog := map[string]interface{}{
		"error": "error.log",
		"warning": "warning.log",
		"notice": "notice.log",
	}

	logdump.Init(newLog)
	
	fmt.Println(newLog)
	// Output:
	// map[error:error.log warning:warning.log notice:notice.log]
}

func ExampleWrite() {
	newLog := map[string]interface{}{
		"error": "error.log",
		"warning": "warning.log",
		"notice": "notice.log",
	}

	logdump.Init(newLog)	
	logdump.Write("error", "there was an error")
    // Output:
	// 
}
