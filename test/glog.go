package main

import (
	"flag"
	"github.com/golang/glog"
	// "github.com/spf13/pflag"
)

func main() {
	//初始化命令行参数
	flag.Set("alsologtostderr", "true")
	flag.Parse()

	glog.Info("hello, glog")
	glog.Warning("warning glog")
	glog.Error("error glog")

	glog.V(3).Infoln("info with v 3")
	glog.V(2).Infoln("info with v 2")
	glog.V(1).Infoln("info with v 1")
	glog.V(0).Infoln("info with v 0")

	// 退出时调用，确保日志写入文件中
	glog.Flush()
}
