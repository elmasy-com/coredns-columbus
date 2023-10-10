package columbus

import (
	"runtime"
	"strconv"
	"sync/atomic"

	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() { plugin.Register("columbus", setup) }

func setup(c *caddy.Controller) error {
	c.Next()

	var err error

	if DomainsChan == nil {

		buffSize := 10000

		if c.NextArg() {
			buffSize, err = strconv.Atoi(c.Val())
			if err != nil {
				log.Fatalf("Failed to convert columbus.buff %s to int: %s\n", c.Val(), err)
			}
		}

		DomainsChan = make(chan *string, buffSize)
	}

	if InsertWorkers == nil || !InsertWorkers.Load() {

		InsertWorkers = new(atomic.Bool)

		nWorkers := runtime.NumCPU()

		if c.NextArg() {
			nWorkers, err = strconv.Atoi(c.Val())
			if err != nil {
				log.Fatalf("Failed to convert columbus.workers %s to int: %s\n", c.Val(), err)
			}
		}

		for i := 0; i < nWorkers; i++ {
			go insertWorker()
		}

		InsertWorkers.Store(true)
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Columbus{Next: next}
	})

	return nil
}
