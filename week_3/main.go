package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type myServer struct{}

func (receiver myServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.String() {
	case "/":
		fmt.Fprintf(w, "this is deafult")
	case "/test":
		fmt.Fprintf(w, "this is test")
	default:
		fmt.Fprintf(w, "unknow http")
	}
}

func signalDeal() error {
	signCh := make(chan os.Signal)

	signal.Notify(signCh)
	s := <-signCh
	fmt.Println("catch system signal", s)
	switch s {
	case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
		return errors.New("find return signal,exit")
	default:
		// 使用 ctrl + z 也退出了，不是我所预期的
		fmt.Println("other signal")
	}

	return nil
}

func main() {
	// use errCtx group with context
	group, errCtx := errgroup.WithContext(context.Background())
	var s myServer
	se := http.Server{
		Handler: s,
		Addr:    ":8080",
	}
	http.Handle("/", s)
	group.Go(func() error {
		defer fmt.Println("goroutine 1 return")
		return se.ListenAndServe()
	})

	// 发送停止http服务信号
	group.Go(func() error {
		select {
		case <-errCtx.Done():
			fmt.Println("goroutine 2 return")
			return se.Shutdown(errCtx)
		}
		return nil
	})

	// 添加处理 linux 信号
	group.Go(func() error {
		err := signalDeal()
		if err != nil {
			fmt.Println("goroutine 3 return")
			return err
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Println("all goroutine are dead get errors:", err)
	}
}
