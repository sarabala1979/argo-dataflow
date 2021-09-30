package http

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/argoproj-labs/argo-dataflow/runner/sidecar/source"
)

type httpSource struct {
	ready bool
}

func New(sourceName, authorization string, process source.Process) source.Interface {
	h := &httpSource{ready: true}
	http.HandleFunc("/sources/"+sourceName, func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != authorization {
			w.WriteHeader(403)
			return
		}
		if !h.ready { // if we are not ready, we cannot serve requests
			w.WriteHeader(503)
			_, _ = w.Write([]byte("not ready"))
			return
		}
		receviedTime := time.Now()
		msg, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		if err := process(r.Context(), msg, receviedTime.UTC()); err != nil {
			w.WriteHeader(500)
			_, _ = w.Write([]byte(err.Error()))
		} else {
			w.WriteHeader(204)
		}
	})
	return h
}

func (s *httpSource) Close() error {
	s.ready = false
	return nil
}
