package server

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/mmarci96/go-ws-traffic-controller/go-reverse-proxy/internal/configs"
)

func Run() error {
	config, err := configs.NewConfiguration()
	if err != nil {
		return fmt.Errorf("could not load configuration: %v", err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping)

	for _, resource := range config.Resources {
		url, _ := url.Parse(resource.Destination_URL)
		proxy := NewProxy(url)
		mux.HandleFunc(resource.Endpoint, ProxyRequestHandler(proxy, url, resource.Endpoint))
	}
	addr := config.Server.Host + ":" + config.Server.Listen_port
	fmt.Printf("[ TinyRP ] Started server on %s\n", addr)

	errorPageDir := "../html/error"
	mux.Handle("/404", http.FileServer(http.Dir(errorPageDir)))
	if err := http.ListenAndServe(config.Server.Host+":"+config.Server.Listen_port, mux); err != nil {
		return fmt.Errorf("could not start the server: %v", err)
	}
	fmt.Printf("[ TinyRP ] Running on %s\n", addr)

	return nil
}
