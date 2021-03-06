package master

import (
	"net"
	"net/http"
	"strconv"
	"time"
)

type ApiServer struct {
	httpServer *http.Server
}

var(
	G_apiServer *ApiServer
)

func handleJobSave(resp http.ResponseWriter,req *http.Request){

}

func InitApiServer()(err error){
	var (
		mux *http.ServeMux
		listener net.Listener
		httpServer *http.Server
	)

	//路由配置
	mux = http.NewServeMux()
	mux.HandleFunc("/job/save",handleJobSave)

	//监听端口
	if listener , err = net.Listen("tcp",":"+ strconv.Itoa(G_config.ApiPort));err !=nil{
		return
	}

	//创建http服务
	httpServer = &http.Server{
		ReadTimeout:time.Duration(G_config.ApiReadTimeout)*time.Millisecond,
		WriteTimeout:time.Duration(G_config.ApiWriteTimeout)*time.Millisecond,
		Handler:mux,
	}

	G_apiServer = &ApiServer{
		httpServer:httpServer,
	}

	go httpServer.Serve(listener)

	return
}