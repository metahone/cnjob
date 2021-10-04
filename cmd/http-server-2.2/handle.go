package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	for k, v := range headers {
		logrus.Println(k, v)
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "ok") // fmt.Fprintf(w, "ok")
}

func final(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "success")
}
