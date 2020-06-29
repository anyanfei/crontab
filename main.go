package main

import (
	"crontab/master"
	"flag"
	"fmt"
	"runtime"
)

var(
	confFile string
)

func initArgs(){
	// main -config ./main.json
	flag.StringVar(&confFile,"config","./main.json","指定main.json")
	flag.Parse()
}

func initEnv(){
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var(
		err error
	)
	//初始化命令行参数
	initArgs()
	//设置线程数
	initEnv()
	//加载配置
	if err = master.InitConfig(confFile);err !=nil{
		goto ERR
	}
	//启动任务管理器
	if err = master.InitJobMgr();err !=nil{
		goto ERR
	}
	//启动http服务
	if err = master.InitApiServer();err!=nil{
		goto ERR
	}
	return
	ERR:
		fmt.Println(err)
}

