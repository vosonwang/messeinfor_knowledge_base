package logrus

import (
	 "github.com/sirupsen/logrus"
	"testing"
	"os"
)

func TestLogurs(t *testing.T)  {
	logrus.Info("hello, world.")
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}

func TestFileLog(t *testing.T)  {
	var log = logrus.New()
	file, err := os.OpenFile("logrus.log", os.O_RDWR|os.O_CREATE, 0)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	log.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}