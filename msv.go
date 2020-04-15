package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"net/http"
	"os"
	"os/signal"
	"syscall"

   "github.com/endormi/msv/tools/endpoint"
   "github.com/endormi/msv/tools/lang"
   "github.com/endormi/msv/tools/req"
   "github.com/endormi/msv/tools/serv"
   "github.com/endormi/msv/tools/server"
   "github.com/endormi/msv/tools/service"
)

func main() {
	var httpAddr = flag.String("http", httpPort, "http address")
	var log log.Logger
	{
		log = log.NewLogfmtLogger(os.Stderr)
		log = log.NewSyncLogger(logger)
		log = log.With(logger,
			"service", "tools",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("message", "Started!")
	defer level.Info(logger).Log("message", "Ended!")

	flag.Parse()
	ctx := context.Background()
	var srv tools.Service
	{
		srv = tools.NewServ(repp, log)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <- c)
	}()

	endpoint := tools.CreateEndpoint(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := account.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <- errs)
}
