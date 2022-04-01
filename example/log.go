// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package main

import (
	"time"

	"github.com/ergoapi/log"
	"github.com/sirupsen/logrus"
)

func main() {
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
