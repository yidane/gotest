package main

import (
	"errors"
	"flag"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/yidane/log4go"
)

func main() {
	logPath := flag.String("logpath", "", "you do not input log file path")

	flag.Parse()

	var logger log4go.Logger
	if *logPath == "" {
		flag.PrintDefaults()
		fmt.Println("log will be outputed by console")
		logger = log4go.NewDefaultLogger(log4go.FINE)
	} else {
		filePath := strings.ToLower(strings.TrimSpace(*logPath))
		logger = make(log4go.Logger)
		logWriter := log4go.NewFileLogWriter(filePath, false)
		if runtime.GOOS != "windows" {
			logWriter.SetFormat("[%D %T] [%L] (%S) %M \r\n")
		}
		logger.AddFilter("file", log4go.FINE, logWriter)
		defer logger.Close()
	}

	logger.Error(errors.New("yidane error"))
	logger.Log(log4go.ERROR, "w32tm", "yidane error")

	defer time.Sleep(time.Second * 3)

	if runtime.GOOS != "windows" {
		logger.Info("current os dot not be supported")
		return
	}

	c := exec.Command("cmd", "/c", "w32tm /resync")
	err := c.Run()
	if err != nil {
		logger.Error(err)
	} else {
		msg, err := c.CombinedOutput()
		if err != nil {
			logger.Error(err)
		} else {
			logger.Log(log4go.INFO, "w32tm", string(msg))
		}
	}
}
