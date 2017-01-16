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

	fmt.Println(newLog)
	// Output:
	// map[error:error.log warning:warning.log notice:notice.log]
}

func ExampleWrite() {
	logdump.Write("not_created", "hi")
    // Output:
	// 
}
