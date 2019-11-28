package main

import (
	log "github.com/sirupsen/logrus"
)

type msFormater struct{}

func (mf msFormater) Format(entry *log.Entry) ([]byte, error) {
	entry.Data["module"] = "safari"

	formatter := log.TextFormatter{}
	return formatter.Format(entry)
}

func main() {
	log.SetFormatter(msFormater{})
	log.SetLevel(log.DebugLevel)
	log.Info("hello")
}
