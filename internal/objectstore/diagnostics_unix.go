package objectstore

import (
	"github.com/kubenext/kubefun/internal/log"
	"os"
	"os/signal"
	"syscall"
)

func initStatusCheck(stopCh <-chan struct{}, logger log.Logger, factories *factoriesCache) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGUSR2)

	done := false
	for !done {
		select {
		case <-stopCh:
			done = true
		case <-sigCh:
			logger.With("factory-count", len(factories.factories)).Debugf("dynamic cache status")
		}
	}

	logger.Debugf("dynamic cache status exiting")
}
