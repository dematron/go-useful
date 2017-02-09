package common

import (
	"log"
	"runtime"
)

func CheckError(err error) error {
	if err != nil {
		log.Printf("Error %s", err)
		// trace
		pc := make([]uintptr, 10) // at least 1 entry needed
		runtime.Callers(2, pc)
		f := runtime.FuncForPC(pc[0])
		file, _ := f.FileLine(pc[0])
		log.Printf("%s %s\n", file, f.Name())
		//
	}
	return err
}
