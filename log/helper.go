package log

import (
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

func ErrorStack(msg string) {
	logrus.Errorf("%s\n%s", msg, debug.Stack())
}
