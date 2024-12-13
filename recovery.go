package recovery

import (
	"fmt"
	"os"
	"runtime/debug"

	"pkg.moe/pkg/logger"
)

func RecoverPanic(errFormat string, mailSend ...bool) {
	if err := recover(); err != nil {
		var errStr string
		if hostname, _ := os.Hostname(); hostname != "" { // nolint: errcheck
			errStr = fmt.Sprintf("[%s] %s", hostname, fmt.Sprintf(errFormat, err))
		} else {
			errStr = fmt.Sprintf(errFormat, err)
		}
		logger.Get().Error(errStr + string(debug.Stack()))

		if len(mailSend) == 0 || mailSend[0] != false {
			go logger.SendMail(errStr + "\r\n\r\n" + string(debug.Stack())) // nolint: errcheck
		}
	}
}
