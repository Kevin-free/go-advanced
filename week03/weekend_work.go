package week03

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("pong"))
	})

	// 模拟单个服务错误退出
	serverOut := make(chan struct{})
	mux.HandleFunc("/shutdown", func(writer http.ResponseWriter, request *http.Request) {
		serverOut<- struct{}{}
	})

	server := http.Server{
		Handler:mux,
		Addr:":8080",
	}

	// g1
	g.Go(func() error {
		return server.ListenAndServe()
	})

	// g2
	g.Go(func() error {
		select {
		case <- ctx.Done():
			log.Println("errgroup exit")
		case <- serverOut:
			log.Println("server out")
		}

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		log.Println("shutting down server")
		return server.Shutdown(timeoutCtx)
	})

	// g3
	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit,syscall.SIGINT, syscall.SIGTERM)

		select {
		case <- ctx.Done():
			return ctx.Err()
		case <- quit:
			return errors.New("quit")
		}
	})

	fmt.Printf("errgroup exiting: %+v\n", g.Wait())
}
