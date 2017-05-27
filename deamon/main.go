package main

//使用golang构建本地服务
//service will install / un-install, start / stop, and run a program as a service (daemon).
//Currently supports Windows XP+, Linux/Upstart, and OSX/Launchd

//使用nssm发布服务
//http://nssm.cc/

//在nssm目录下输入命令 nssm install，将会启动GUI界面。按照 http://nssm.cc/usage 所说明操作即可完成服务的安装

//TODO:服务安装成功后，无法启动，没明白

import (
	"log"

	"github.com/kardianos/service"
)

var logger service.Logger

type program struct{}

func (p *program) run() {
	// Do work here
}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "golangService",
		DisplayName: "golangServiceDisplayName",
		Description: "测试golang写服务",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
