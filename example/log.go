// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package main

import (
	"time"

	"github.com/ergoapi/log"
	"github.com/ergoapi/log/factory"
	"github.com/ergoapi/log/survey"
	"github.com/sirupsen/logrus"
)

func dolog() {
	logfile := log.GetFileLogger("/tmp/dofile1.log")
	logfile.SetLevel(logrus.DebugLevel)
	logfile.Debug("debug level")
}

func dolog2() {
	logfile := log.GetFileLogger("/tmp/dofile2.log")
	logfile.SetLevel(logrus.DebugLevel)
	logfile.Debug("debug level")
}

func Survey(f factory.Factory) {
	l := f.GetLog()
	answer, err := l.Question(&survey.QuestionOptions{
		Question:     "Multiple sync configurations found. Which one do you want to use?",
		DefaultValue: "1",
		Options:      []string{"1", "2", "3"},
	})
	if err != nil {
		l.Panic(err)
	}
	l.WriteString("\n")
	l.Infof("an: %s", answer)
}

func main() {
	f := factory.DefaultFactory()
	Survey(f)
	dolog()
	dolog2()
	flog := log.GetInstance()
	flog.SetLevel(logrus.DebugLevel)
	log.StartFileLogging()
	flog.StartWait("debug level")
	flog.Debug("debug66666")
	time.Sleep(time.Second * 2)
	flog.Info("dididiidd")
	flog.StopWait()
	// flog.SetLevel(logrus.InfoLevel)
	flog.StartWait("info level")
	flog.Debug("debug777777")
	time.Sleep(time.Second * 1)
	flog.Info("hahahahahahah")
	flog.StopWait()
	flog.Done("done")
	flog.Error("error")
	flog.Fatal("fatal")
}
