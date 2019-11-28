package main

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
	log.SetOutput(ioutil.Discard)
}

var testString = `Quisque venenatis ipsum vel ornare porta. 
Aliquam at tristique purus, non maximus urna. 
Morbi et massa vel diam laoreet vestibulum in nec nulla.`

var testMapSmallMembers = map[string]interface{}{
	"hello": "world",
	"module": "safari"
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

func BenchmarkLogrus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logrus.Print(testString)
	}
}

func BenchmarkLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		log.Print(testString)
	}
}

func BenchmarkLogrusMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logrus.Print(testMap)
	}
}

func BenchmarkLogMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		log.Print(testMap)
	}
}

func BenchmarkLogrusMapSmallMembers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logrus.Print(testMapSmallMembers)
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
		logrus.Print(testString)
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
