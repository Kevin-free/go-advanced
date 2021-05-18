package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"biohouse/internal/di"
	GameConfig "biohouse/internal/game/config"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/conf/paladin/apollo"
	"github.com/bilibili/kratos/pkg/log"
)

func main() {
	flag.Parse()
	var err error
	if err = paladin.Init(apollo.PaladinDriverApollo); err != nil {
		panic(err)
	}
	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("biohouse start")

	// 配置初始化
	if !GameConfig.Init() {
		return
	}

	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Info("biohouse exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
