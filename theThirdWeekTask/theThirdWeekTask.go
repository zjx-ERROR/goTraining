package thethirdweektask

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

var (
	eg  *errgroup.Group
	ctx context.Context
)

func Server() {
	eg, ctx = errgroup.WithContext(context.Background())

	server := &http.Server{Addr: "127.0.0.1:8080", Handler: nil}

	eg.Go(func() error {
		return server.ListenAndServe()
	})

	eg.Go(
		func() error {
			select {
			case <-ctx.Done():
				return server.Shutdown(ctx)
			}
		})

	eg.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		select {
		case <-ctx.Done():
			return errors.New("out")
		case <-sigs:
			return errors.New("out")
		}
	})
	if err := eg.Wait(); err != nil {
		log.Print(err)
	}
}
