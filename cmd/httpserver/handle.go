package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	ip, err := GetIP(r)
	if err != nil {
		log.Warningf("GetIP error: %s", err)
	}
	log.Infof("recv %s request route: healthz", ip)

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "ok") // fmt.Fprintf(w, "ok")
	log.Infof("response %s http code %d", ip, http.StatusOK)
}

func root(w http.ResponseWriter, r *http.Request) {
	ip, err := GetIP(r)
	if err != nil {
		log.Warningf("GetIP error: %s", err)
	}
	log.Infof("recv %s request route: root", ip)

	headers := r.Header
	for k, v := range headers {
		log.Println(k, v)
	}

	w.WriteHeader(http.StatusOK)
	for k, v := range headers {
		// Because Header is a map[string][]string, two loops are required to access all headers.
		for _, item := range v {
			w.Header().Add(k, item)
			io.WriteString(w, fmt.Sprintf("%s: %s\n", k, v))
		}
	}

	if len(Version) > 0 { // os.Getenv("VERSION")
		w.Header().Add("Version", Version)
		io.WriteString(w, fmt.Sprintf("Version: %s\n", Version))
	}
	log.Infof("response %s http code %d", ip, http.StatusOK)
}

// GetIP returns request real ip.
func GetIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	ip = r.Header.Get("X-Forward-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i, nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	return "", errors.New("no valid ip found")
}
