package di

import (
	"biohouse/internal/service"

	"git.huoys.com/middle-end/library/pkg/net/comet"
	"github.com/bilibili/kratos/pkg/log"
)

//go:generate kratos tool wire
type App struct {
	svc *service.Service
	// http *gin.Engine
	tcp *comet.Server
}

func NewApp(svc *service.Service /*h *gin.Engine,*/, c *comet.Server) (app *App, closeFunc func(), err error) {
	app = &App{
		svc: svc,
		// http: h,
		tcp: c,
	}
	closeFunc = func() {
		// ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
		// if err := kgin.Shutdown(h, ctx); err != nil {
		// 	log.Error("httpSrv.Shutdown error(%v)", err)
		// }

		// cancel()
		svc.Close()
		log.Info("httpSrv.Shutdown")
	}
	return
}
