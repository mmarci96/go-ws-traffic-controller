package server

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/mmarci96/go-ws-traffic-controller/go-reverse-proxy/internal/configs"
)

type spaHandler struct {
	staticDir string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fs := http.Dir(h.staticDir)
	path := r.URL.Path

	// Check if requested file exists
	if _, err := fs.Open(path); os.IsNotExist(err) {
		// Serve index.html for SPA routing
		r.URL.Path = "/"
	}

	http.FileServer(fs).ServeHTTP(w, r)
}

// Run starts server and listens on defined port
func Run() error {
	// load configurations from config file
	config, err := configs.NewConfiguration()
	if err != nil {
		return fmt.Errorf("could not load configuration: %v", err)
	}
	// Creates a new router
	mux := http.NewServeMux()
	// Iterating through the configuration resource and registering them
	// into the router.
	for _, resource := range config.Resources {
		url, _ := url.Parse(resource.Destination_URL)
		proxy := NewProxy(url)
		mux.HandleFunc(resource.Endpoint, ProxyRequestHandler(proxy, url, resource.Endpoint))
	}
	fmt.Printf("[ TinyRP ] Server running on http://%s:%s\n", config.Server.Host, config.Server.Listen_port)
	// Registering the healthcheck endpoint
	mux.HandleFunc("/ping", ping)

	// Setup SPA static file server
	spa := spaHandler{staticDir: config.Static.Dir}
	mux.Handle("/", spa)

	// Running proxy server
	if err := http.ListenAndServe(config.Server.Host+":"+config.Server.Listen_port, mux); err != nil {
		return fmt.Errorf("could not start the server: %v", err)
	}
	return nil
}
