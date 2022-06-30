// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package main

import (
	"time"

	"github.com/ergoapi/log"
	"github.com/sirupsen/logrus"
)

func dolog() {
	logfile := log.GetFileLogger("/tmp", "dofile1.log")
	logfile.SetLevel(logrus.DebugLevel)
	logfile.Debug("debug level")
}

func dolog2() {
	logfile := log.GetFileLogger("/tmp", "dofile2.log")
	logfile.SetLevel(logrus.DebugLevel)
	logfile.Debug("debug level")
}

func main() {
	dolog()
	dolog2()
	flog := log.GetInstance()
	flog.SetLevel(logrus.DebugLevel)
	log.StartFileLogging("/tmp", "test.log")
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
