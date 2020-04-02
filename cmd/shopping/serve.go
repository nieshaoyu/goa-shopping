package shopping

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"sync"

	"goa-shopping/config"
	controller "goa-shopping/controller"
	"goa-shopping/gen/auth"
	log "goa-shopping/gen/log"
	"goa-shopping/gen/shop"
)

func RunServer(cfg *config.Config) {

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New("goa-shopping", !cfg.Debug)
	}

	// Initialize the services.
	var (
		userSvc auth.Service
		shopSvc shop.Service
	)
	{
		userSvc = controller.NewAuth(logger)
		shopSvc = controller.NewShop(logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		userEndpoints *auth.Endpoints
		shopEndpoints *shop.Endpoints
	)
	{
		userEndpoints = auth.NewEndpoints(userSvc)
		shopEndpoints = shop.NewEndpoints(shopSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	addr := fmt.Sprintf("http://%s:%s", cfg.Server.Host, cfg.Server.HTTPPort)
	u, _ := url.Parse(addr)
	handleHTTPServer(ctx, u.Host, shopEndpoints, userEndpoints, &wg, errc, logger, cfg.Debug)

	// Wait for signal.
	logger.Infof("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Info("exited")
}
