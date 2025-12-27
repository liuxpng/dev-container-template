// Package main provides a simple HTTP server example
// to verify your Go development environment is working.
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"runtime"
)

// ServerInfo contains information about the server environment
type ServerInfo struct {
	Status     string   `json:"status"`
	GoVersion  string   `json:"go_version"`
	GOOS       string   `json:"goos"`
	GOARCH     string   `json:"goarch"`
	NumCPU     int      `json:"num_cpu"`
	GoRoot     string   `json:"goroot"`
	Tools      []string `json:"tools"`
}

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go Dev Container</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            max-width: 800px;
            margin: 50px auto;
            padding: 20px;
            background: #f5f5f5;
        }
        .container {
            background: white;
            padding: 30px;
            border-radius: 10px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
        }
        h1 { color: #00ADD8; }
        .success { color: green; font-size: 18px; }
        table { border-collapse: collapse; width: 100%; margin: 20px 0; }
        td { padding: 10px; border: 1px solid #ddd; }
        tr:nth-child(even) { background: #f9f9f9; }
        code { background: #e0e0e0; padding: 2px 6px; border-radius: 4px; }
        .tools { display: flex; flex-wrap: wrap; gap: 10px; margin: 20px 0; }
        .tool { background: #00ADD8; color: white; padding: 5px 15px; border-radius: 20px; font-size: 14px; }
    </style>
</head>
<body>
    <div class="container">
        <h1>Go Dev Container</h1>
        <p class="success">Your Go development environment is ready!</p>

        <h2>Environment Information</h2>
        <table>
            <tr><td><strong>Go Version</strong></td><td>{{.GoVersion}}</td></tr>
            <tr><td><strong>OS</strong></td><td>{{.GOOS}}</td></tr>
            <tr><td><strong>Architecture</strong></td><td>{{.GOARCH}}</td></tr>
            <tr><td><strong>CPUs</strong></td><td>{{.NumCPU}}</td></tr>
            <tr><td><strong>GOROOT</strong></td><td>{{.GoRoot}}</td></tr>
        </table>

        <h2>Installed Tools</h2>
        <div class="tools">
            {{range .Tools}}
            <span class="tool">{{.}}</span>
            {{end}}
        </div>

        <h2>Quick Links</h2>
        <ul>
            <li><a href="/api/info">View JSON API Info</a></li>
        </ul>

        <h2>Next Steps</h2>
        <ol>
            <li>Edit files in the <code>src/</code> directory</li>
            <li>Use <code>go mod init</code> to initialize a module</li>
            <li>Run <code>air</code> for hot reload during development</li>
            <li>Use <code>golangci-lint run</code> for code linting</li>
            <li>Debug with <code>dlv</code> (Delve)</li>
        </ol>
    </div>
</body>
</html>
`

func getServerInfo() ServerInfo {
	return ServerInfo{
		Status:    "ok",
		GoVersion: runtime.Version(),
		GOOS:      runtime.GOOS,
		GOARCH:    runtime.GOARCH,
		NumCPU:    runtime.NumCPU(),
		GoRoot:    runtime.GOROOT(),
		Tools: []string{
			"golangci-lint",
			"delve",
			"gopls",
			"air",
			"mockgen",
			"swag",
			"goimports",
		},
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.New("home").Parse(htmlTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	info := getServerInfo()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, info); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	info := getServerInfo()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(info); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/info", apiHandler)

	port := 8080
	addr := fmt.Sprintf(":%d", port)

	fmt.Println("Go Dev Container Server")
	fmt.Printf("Serving at http://localhost:%d\n", port)
	fmt.Println("Press Ctrl+C to stop")

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
