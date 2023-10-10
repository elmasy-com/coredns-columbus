package columbus

import (
	"context"
	"net/http"
	"sync/atomic"

	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/elmasy-com/elnet/validator"

	"github.com/miekg/dns"

	eldns "github.com/elmasy-com/elnet/dns"
)

var (
	log                        = clog.NewWithPlugin("columbus")
	DomainsChan   chan *string = nil
	InsertWorkers *atomic.Bool = nil // Indicate whether insertWokers started
)

type Columbus struct {
	Next plugin.Handler
}

func (e Columbus) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {

	di := NewDomainInserter(w)

	return plugin.NextOrFailure(e.Name(), e.Next, ctx, di, r)
}

func (e Columbus) Name() string { return "columbus" }

type DomainInserter struct {
	dns.ResponseWriter
}

func NewDomainInserter(w dns.ResponseWriter) *DomainInserter {
	return &DomainInserter{ResponseWriter: w}
}

func (r *DomainInserter) WriteMsg(res *dns.Msg) error {

	if len(DomainsChan) == cap(DomainsChan) {
		log.Warningf("DomainChannel is full!")
	}

	if res.Rcode == 0 && len(res.Answer) > 0 && len(res.Question) > 0 {
		DomainsChan <- &res.Question[0].Name

	}

	return r.ResponseWriter.WriteMsg(res)
}

func insertWorker() {

	for d := range DomainsChan {

		if !validator.Domain(*d) {
			continue
		}

		req, err := http.NewRequest("PUT", "https://columbus.elmasy.com/api/insert/"+eldns.Clean(*d), nil)
		if err != nil {
			log.Errorf("Failed to create request for %s: %s", *d, err)
			continue
		}

		req.Header.Add("User-Agent", "coredns-columbus")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Errorf("Failed to PUT for %s: %s", *d, err)
			continue
		}

		resp.Body.Close()
	}
}
