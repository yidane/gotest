package main

import (
	"flag"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/yidane/log4go"
)

var logger log4go.Logger

func main() {
	logPath := flag.String("logpath", "", "you do not input log file path")

	flag.Parse()

	if *logPath == "" {
		flag.PrintDefaults()
		fmt.Println("log will be outputed by console")
		logger = log4go.NewDefaultLogger(log4go.FINE)
	} else {
		filePath := strings.ToLower(strings.TrimSpace(*logPath))
		logger = make(log4go.Logger)
		logWriter := log4go.NewFileLogWriter(filePath, false)
		if runtime.GOOS == "windows" {
			logWriter.SetFormat("[%D %T] [%L] (%S) %M \r\n")
		}
		logger.AddFilter("file", log4go.FINE, logWriter)
		defer logger.Close()
	}

	defer time.Sleep(time.Second * 1)

	if runtime.GOOS != "windows" {
		logger.Info("current os dot not be supported")
		return
	}

	execStartNet()
	execResync()
}

func execStartNet() {
	startService := exec.Command("cmd", "/c", "net start w32time")
	msg, err := startService.CombinedOutput()

	if err != nil {
		ee, ok := err.(*exec.ExitError)
		if !ok {
			logger.Error(err)
			return
		}
		waitStatus, ok := ee.ProcessState.Sys().(syscall.WaitStatus)
		if !ok {
			logger.Error(err)
			return
		}
		if waitStatus.ExitStatus() == 2 {
			logger.Log(log4go.WARNING, "w32tm start ", "w32time is running")
			return
		}
		logger.Log(log4go.ERROR, "w32tm start ", string(msg))
	} else {
		logger.Log(log4go.INFO, "w32tm start ", string(msg))
	}
}

func execResync() {
	startService := exec.Command("cmd", "/c", "w32tm /resync")
	msg, err := startService.CombinedOutput()

	if err != nil {
		logger.Error(err)
	} else {
		logger.Log(log4go.INFO, "w32tm resync", string(msg))
	}
}
