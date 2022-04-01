# log

> 参考 devspace v5.x版本

[devspace](https://github.com/loft-sh/devspace/tree/master/pkg/util/log)

## usage

```go
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

```
