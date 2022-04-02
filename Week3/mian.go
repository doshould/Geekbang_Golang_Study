package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
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
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	//单个服务错误退出
	serverOut := make(chan struct{})
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		serverOut <- struct{}{}
	})

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	//g1退出后, context 将不再阻塞, g2, g3都会随着退出, 最后go.Wait()退出, 所有协程都退出
	g.Go(func() error {
		//err := server.ListenAndServe()
		//if err != nil {
		//	log.Println("g1 will exit with error.", err.Error())
		//}
		//return err
		return server.ListenAndServe()
	})

	//g2退出时调用shutdown，g1退出，g2退出后不再阻塞，g3随着退出，最后go.Wait()退出，所有协程退出
	g.Go(func() error {
		select {
		case <-ctx.Done():
			log.Println("errgroup exit...")
		case <-serverOut:
			log.Println("server will out")
		}
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		log.Println("shutting down server...")
		return server.Shutdown(timeoutCtx)
	})

	//g3捕获os退出信号退出，g3退出后context不再阻塞
	//g2退出，g2退出时调用shutdown，g1退出
	//最后go.Wait()退出，所有协程退出
	g.Go(func() error {
		exit := make(chan os.Signal, 0)
		signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sigNal := <-exit:
			return errors.Errorf("get os sigNal: %v", sigNal)
		}
	})
	fmt.Printf("errgroup exiting: %+v\n", g.Wait())
}
