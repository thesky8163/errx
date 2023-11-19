package errx

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func errinfo_time() (errmsg string) {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("Failed to get caller information")
	}
	funcName := runtime.FuncForPC(pc).Name()
	currentTime := time.Now().Format("2006-01-02 15:04:05.000")
	return fmt.Sprintf("%s %s:%d: %s", currentTime, file, line, funcName)
}
func errinfo() (errmsg string) {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("Failed to get caller information")
	}
	funcName := runtime.FuncForPC(pc).Name()
	return fmt.Sprintf("%s:%d: %s", file, line, funcName)
}

func Out(err error) bool {
	if err != nil {
		fmt.Printf("%s [err: %s]\n", errinfo_time(), err)
		return true
	}
	return false
}
func Outlog(err error) bool {
	if err != nil {
		log.Printf("%s [err: %s]\n", errinfo(), err)
		return true
	}
	return false
}

// func Printlog(err error, add func()) {
// 	if err != nil {
// 		fmt.Printf("%s", errinfo())
// 		add()
// 	}
// }
// func PrintlogReturn(err error, add func() (int, error)) (int, error) {
// 	if err != nil {
// 		fmt.Printf("%s", errinfo())
// 		return add()
// 	}
// 	return 0, err
// }
// func Print(err error, add func()) {
// 	if err != nil {
// 		fmt.Printf("%s", errinfo())
// 		add()
// 	}
// }
// func PrintRuturn(err error, add func()) bool {
// 	if err != nil {
// 		fmt.Printf("%s", errinfo_time())
// 		add()
// 		return true
// 	}
// 	return false
// }
