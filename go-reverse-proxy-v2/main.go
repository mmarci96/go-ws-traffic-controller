package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host       string `yaml:"host"`
		ListenPort string `yaml:"listen_port"`
	} `yaml:"server"`
	Static struct {
		Dir string `yaml:"dir"`
	} `yaml:"static"`
	Resources []struct {
		Endpoint       string `yaml:"endpoint"`
		DestinationURL string `yaml:"destination_url"`
	} `yaml:"resources"`
}

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

func main() {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}

	// Validate static directory
	if _, err := os.Stat(config.Static.Dir); os.IsNotExist(err) {
		log.Fatalf("Static directory %s not found", config.Static.Dir)
	}

	mux := http.NewServeMux()

	// Setup reverse proxies
	for _, res := range config.Resources {
		dest, err := url.Parse(res.DestinationURL)
		if err != nil {
			log.Fatalf("Invalid destination URL: %v", err)
		}

		proxy := httputil.NewSingleHostReverseProxy(dest)
		prefix := res.Endpoint + "/"

		mux.Handle(prefix, http.StripPrefix(res.Endpoint, proxy))
		mux.HandleFunc(res.Endpoint, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, prefix, http.StatusPermanentRedirect)
		})
	}

	// Setup SPA static file server
	spa := spaHandler{staticDir: config.Static.Dir}
	mux.Handle("/", spa)

	// Start server
	addr := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.ListenPort)
	log.Printf("Server listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
