package tool

import (
	"fmt"
	"time"
)

func ArgsToS(args ...interface{}) []string {
	strArgs := []string{}
	for _, arg := range args {
		var strV string
		switch f := arg.(type) {
		case time.Time:
			strV = f.Format("2006-01-02 15:04:05")
		default:
			strV = fmt.Sprintf("%v", arg)
		}
		strArgs = append(strArgs, strV)
	}
	return strArgs
}
