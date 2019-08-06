package msg

import (
	"strconv"
	"strings"
)

var msgs map[int]string

func init() {
	msgs = make(map[int]string)
	msgs[1000] = "{0} success"
	msgs[9000] = "no authorization"
	msgs[9001] = "authorization expired"
	msgs[9002] = "{0} wrong"
	msgs[9003] = "{0} not found"
	msgs[9004] = "{0} error"
	msgs[9005] = "you can only show yourself"
	msgs[9006] = "{0} fail"
	msgs[9007] = "{0} can not be empty"
	msgs[9008] = "{0} must be a number"
	msgs[9009] = "{0} already existed"
	msgs[9010] = "{0} can not be zero"
	msgs[9011] = "{0} must be bigger than {1}"
	msgs[9012] = "{0} must be between {1} and {2}"
}

// GetMsg returns the specified msg
func GetMsg(id int, args ...string) string {
	msg := msgs[id]
	for i, v := range args {
		msg = strings.ReplaceAll(msg, "{"+strconv.Itoa(i)+"}", v)
	}
	return msg
}
