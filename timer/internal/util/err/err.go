package err

import "timer/internal/util/log"

func GetErr(msg string, err error) {
	if err != nil {
		log.Error.Fatalf("error:%v err->%v\n", msg, err)
	}
}
