package handlers

import (
	"crypto/tls"
	"fmt"
	"github.com/statping/statping/utils"
	"net/http"
)

func startServer(host string) {
	httpError = make(chan error)
	httpServer = &http.Server{
		Addr:         host,
		WriteTimeout: timeout,
		ReadTimeout:  timeout,
		IdleTimeout:  timeout,
		Handler:      router,
	}
	httpServer.SetKeepAlivesEnabled(false)
	if err := httpServer.ListenAndServe(); err != nil {
		httpError <- err
	}
}

func startSSLServer(ip string) {
	httpError = make(chan error)
	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	srv := &http.Server{
		Addr:         fmt.Sprintf("%v:%v", ip, 443),
		Handler:      router,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		WriteTimeout: timeout,
		ReadTimeout:  timeout,
		IdleTimeout:  timeout,
	}
	if err := srv.ListenAndServeTLS(utils.Directory+"/server.crt", utils.Directory+"/server.key"); err != nil {
		httpError <- err
	}
}
