package columbus

import (
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() {

	buffSize := os.Getenv("COREDNS_COLUMBUS_BUFFSIZE")
	if buffSize == "" {
		buffSize = "100000"
	}

	buffSizeInt, err := strconv.Atoi(buffSize)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse COREDNS_COLUMBUS_BUFFSIZE: %s", err))
	}

	numWorker := os.Getenv("COREDNS_COLUMBUS_WORKER")
	if numWorker == "" {
		numWorker = strconv.Itoa(runtime.NumCPU())
	}

	numWorkerInt, err := strconv.Atoi(numWorker)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse COREDNS_COLUMBUS_WORKER: %s", err))
	}

	DomainsChan = make(chan *string, buffSizeInt)

	for i := 0; i < numWorkerInt; i++ {
		go domainInserter()
	}

	plugin.Register("columbus", setup)
}

func setup(c *caddy.Controller) error {
	c.Next()
	if c.NextArg() {
		return plugin.Error("columbus", c.ArgErr())
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Columbus{Next: next}
	})

	return nil
}
