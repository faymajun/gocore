package log

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestInit(t *testing.T) {
	Init()
	logger.Trace("this is a test trace log")
	logger.Debug("this is a test info log")
	logger.Info("this is a test info log")
	logger.Warn("this is a test warn log")
	logger.Error("this is a test error log")
	panic("test panic")
}

var logger = logrus.WithField("com", "log")

func TestLogFiled(t *testing.T) {
	Init()
	logger.Trace("this is a test trace log")
	logger.Debug("this is a test info log")
	logger.Info("this is a test info log")
	logger.Warn("this is a test warn log")
	logger.Error("this is a test error log")
}
