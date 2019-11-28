package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {

	logrus.SetOutput(ioutil.Discard)
	log.SetOutput(ioutil.Discard)

	logger = logrus.New()
	logger.SetOutput(ioutil.Discard)
	logger.SetNoLock()
	logger.SetFormatter(&logrus.TextFormatter{
		DisableSorting:   true,
		DisableColors:    true,
		DisableTimestamp: true,
	})
}

var testString = `Quisque venenatis ipsum vel ornare porta. 
Aliquam at tristique purus, non maximus urna. 
Morbi et massa vel diam laoreet vestibulum in nec nulla.`

var logrusAdvantageMap = func() map[string]interface{} {
	out := make(map[string]interface{})
	for i := 0; i < 1000; i++ {
		out[fmt.Sprint(i)] = i
	}
	return out
}()

var testMapSmallMembers = map[string]interface{}{
	"hello":  "world",
	"module": "safari",
	"foobar": 45,
	"f":      "g",
	"h":      "i",
}

var testMap = map[string]interface{}{
	"hello": "world",
	"module": map[string]string{
		"safari": "logrus",
	},
	"foobar": 45,
}

func BenchmarkLogrusBestCase(b *testing.B) {

	for i := 0; i < b.N; i++ {
		logger.Print(logrusAdvantageMap)
	}
}

func BenchmarkLogBestCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		log.Print(logrusAdvantageMap)
	}
}

func BenchmarkLogrus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logger.Print(testString)
	}
}

func BenchmarkLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		log.Print(testString)
	}
}

func BenchmarkLogrusMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logger.Print(testMap)
	}
}

func BenchmarkLogMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		log.Print(testMap)
	}
}

func BenchmarkLogrusMapSmallMembers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logger.Print(testMapSmallMembers)
	}
}

func BenchmarkLogMapSmallMembers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		log.Print(testMapSmallMembers)
	}
}

func BenchmarkLogrusCallerReport(b *testing.B) {
	logrus.SetReportCaller(true)
	for i := 0; i < b.N; i++ {
		logger.Print(testString)
	}
	logrus.SetReportCaller(false)
}

func BenchmarkLogCallerReport(b *testing.B) {
	log.SetFlags(log.Llongfile | log.Ltime)
	for i := 0; i < b.N; i++ {
		log.Print(testString)
	}
	log.SetFlags(log.Ltime)
}
