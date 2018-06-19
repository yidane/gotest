package main

import (
	"flag"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
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
		if runtime.GOOS != "windows" {
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

	execCommand("net start w32time")
	execCommand("w32tm /resync")
}

func execCommand(c string) {
	startService := exec.Command("cmd", "/c", c)
	msg, err := startService.CombinedOutput()

	if ee, ok := err.(*exec.ExitError); ok {
		fmt.Println(ee.Success())
	} else {

	}

	if err != nil {
		logger.Error(err)
	} else {
		logger.Log(log4go.INFO, "w32tm", "succeed")
	}
	logger.Log(log4go.INFO, "w32tm", string(msg))
}
