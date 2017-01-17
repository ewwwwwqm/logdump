# Logdump

[![Godoc Reference][godoc-img]][godoc]
[![Go Report Card][goreportcard-img]][goreportcard]

Package for writing logs to multiple files

## Usage
``` $ go get github.com/ewwwwwqm/logdump ```

## Example
``` main.go ```
```go
package main

import (
	"log"

	"github.com/ewwwwwqm/logdump"
)

func main() {
	// create message zones and assign files for output
	newLog := map[string]interface{}{
		"error": "error.log",
		"warning": "warning.log",
		"notice": "notice.log",
	}

	// init logger
	logdump.Init(newLog)

	// for ex. dump error
	logdump.Write("error", "something had an error")

	// default output
	log.Println("Standard output to os.Stderr")

	// few more errors
	logdump.Write("notice", "some notice")
	logdump.Write("error", "another error")
	logdump.Write("warning", "fancy warning")

	// wrong message zone, output to os.Stderr
	logdump.Write("wrong_zone", "fail")
}
```
Before running:

![Before run](https://lh3.googleusercontent.com/qu3mHs6Ug0fxwtFFDKb5BiJdsQRG1gpSyygB4mi4869ukJpdTtjoVaNzpwGzyBeSk1eZOpyvejSib56FayhNa1PP9FPIHvZ9jEED_bPO7rrLlcsNJBQzmuydzC2Kww5d-s1DlChnFOpiv3PV_bzYHYhVLODXJeq6IC8DfuzRQUq5jTBJIifzwdGXmaFtBW3islxQffHn4OJTbAqldgXXyPL4dX-5QO3GFagcjom-FN2Jm0c069H3ZhqvR6c_N1NbHuDsDR7pOggKq8R8cRbmdd5n4ZW2zrQbopzvH_H87pJdfVovAB8W1A9XxB6iebZk4PcRp1eHmFygImkVlCEd1xYWUfGFuBG_29YJ0VxVyh1e_uEtJ2uUnFh7lqerYHcSdyZVjaeTpZvLnkB4u-ppwmsU3KzyBKIU-x1_9nYQWjVkV5VUU8bFvFY2lFF4misTR7U-jaG6HKMVLhIaBVLY-RO7zQpOpDSesV3AFIo5sHhoiPNzMvwpEt1pOOjHRenwfkGQT0vvpIm7zrv00NJAGAXkmJviKC3qVK8n7xkO36-um4icFbvArQxIyQqtj_0OW8n7z2pb5eDKicR4qrNA066jHGoTcuaS7PSQZAK-37f2OyR32N-h1L3lLef_1WAepheI9A8G7eDQkvPN_reG3p_ylzscQz5jsXlmTHYJeg=w609-h84-no)

Run the app:

``` $ go run main.go ```

![After run](https://lh3.googleusercontent.com/ADOHxXnHMnZO1Dt5oqHhWBkuTq-5oM13mhQUcQtoMZ2LsPQqecsACPepOTbXZO9bemgJb0rHIKQAD7Sa45s6CXccUS7ugNbYehJZ69tXcj8y2PYBXLnT0SgzyXTmPVlY30r-BKuyBVFoMTsoZIE5Fx6w9LP2hDfd_Nb_vpPzy2sQRVSwxRXGilJKwSuc6M0b8YriUk4EB6e_d4FWGF6LTV9F6NeXg-WRKp5JS5XoHi5WAkW9ntlpKOd_EbMuNnw4IjJIbjI0d-geZEpXoT7RsKTRVjARUBuPR3K8LZ4MgPQMLM7e4MAC5-uH9JqB6Sp5Dbu9viEbrtlBG8HrhQmPlztJ7KBfBGxVmauzoRIqY4nOBffhXu7Wd89_jw5PsBtxr4tyf8w7bApdET2q98V2_UtDFz47L5mZ44OYPip6wEoSBEixSOhgajUdNY-Sz_HAWZDdszqyYA0R_H1SvGpMD-AVeTyTH849UAyl7yHJ4s2uON14VPztpJP-WvkRhpBnv5sDw-LM30uAOsuYfAEh8WLSoKxGdqxPlJsHnwDeRb_rvxAyKKp8PAz_bA0k8kywdIM8V6MgA8_5NIE5xCh-qfOgqAxBMog9qVSnSue5Tmd5bfgpmIQlQ5HuK9qUK2cjz6Utce_Eq70FdGprruyiT-n1ymg7fnPv1SJDHtYh6w=w609-h84-no)

Files were created.

#### Output
``` os.Stderr: ```
```
2017/01/13 17:45:48 Standard output to os.Stderr
2017/01/13 17:45:48 Type wrong_zone was not assigned for log output
```
``` error.log ```
```
2017/01/13 17:45:48 Start error logging
2017/01/13 17:45:48 something had an error
2017/01/13 17:45:48 another error
```
``` warning.log ```
```
2017/01/13 17:45:48 Start warning logging
2017/01/13 17:45:48 fancy warning
```
``` notice.log ```
```
2017/01/13 17:45:48 Start notice logging
2017/01/13 17:45:48 some notice
```

[godoc]: http://godoc.org/github.com/ewwwwwqm/logdump
[godoc-img]: https://godoc.org/github.com/ewwwwwqm/logdump?status.svg
[ci-img]: https://travis-ci.org/ewwwwwqm/logdump.svg?branch=master
[cov-img]: https://coveralls.io/repos/github/ewwwwwqm/logdumpbadge.svg?branch=master
[ci]: https://travis-ci.org/ewwwwwqm/logdump
[cov]: https://coveralls.io/github/ewwwwwqm/logdump?branch=master
[goreportcard-img]: https://goreportcard.com/badge/github.com/ewwwwwqm/logdump
[goreportcard]: https://goreportcard.com/report/github.com/ewwwwwqm/logdump
