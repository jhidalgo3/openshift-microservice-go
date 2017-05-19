package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	mux              = http.NewServeMux()
	listenAddr       string
	closeConnections bool
)

type (
	Content struct {
		App             string
		Namespace       string
		Version         string
		Hostname        string
		RefreshInterval string
		ExtraInfo       string
		SkipErrors      bool
		ShowVersion     bool
		Name            string
		Ip              string
	}

	Ping struct {
		Instance string `json:"instance"`
		Version  string `json:"version"`
	}
)

func init() {
	flag.StringVar(&listenAddr, "listen", ":8080", "listen address")
	flag.BoolVar(&closeConnections, "close-conn", false, "send Connection:close as a response header for all connections")
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return hostname
}

func getVersion() string {
	ver := os.Getenv("VERSION")
	if ver == "" {
		ver = "0.1"
	}

	return ver
}

func loadTemplate(filename string) (*template.Template, error) {
	return template.ParseFiles(filename)
}

func index(w http.ResponseWriter, r *http.Request) {
	remote := r.RemoteAddr

	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		remote = forwarded
	}

	log.Printf("request from %s\n", remote)

	tpl := template.Must(template.New("out").Parse(html))
	/*t, err := loadTemplate("templates/index.html.tmpl")
	if err != nil {
		fmt.Printf("error loading template: %s\n", err)
		return
	}*/

	name := os.Getenv("NAME")
	ip := os.Getenv("IP")
	app := os.Getenv("APP")
	namespace := os.Getenv("NAMESPACE")
	extraInfo := os.Getenv("EXTRA_INFO")

	hostname := getHostname()
	refreshInterval := os.Getenv("REFRESH_INTERVAL")
	if refreshInterval == "" {
		refreshInterval = "1000"
	}

	cnt := &Content{
		App:             app,
		Namespace:       namespace,
		Name:            name,
		Version:         getVersion(),
		Hostname:        hostname,
		RefreshInterval: refreshInterval,
		ExtraInfo:       extraInfo,
		SkipErrors:      os.Getenv("SKIP_ERRORS") != "",
		ShowVersion:     os.Getenv("SHOW_VERSION") != "",
		Ip:              ip,
	}

	tpl.Execute(w, cnt)
}

func ping(w http.ResponseWriter, r *http.Request) {
	if closeConnections {
		w.Header().Set("Connection", "close")
	}

	hostname := getHostname()
	p := Ping{
		Instance: hostname,
		Version:  getVersion(),
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	flag.Parse()

	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/", index)

	hostname := getHostname()

	log.Printf("instance: %s\n", hostname)
	log.Printf("close connections: %v", closeConnections)
	log.Printf("listening on %s\n", listenAddr)

	if err := http.ListenAndServe(listenAddr, mux); err != nil {
		log.Fatalf("error serving: %s", err)
	}
}
