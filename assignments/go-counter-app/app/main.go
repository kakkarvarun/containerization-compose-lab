package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var version = "dev" // set via -ldflags

func main() {
	port := getenv("APP_PORT", "8080")
	greeting := getenv("APP_GREETING", "Hello from Go")

	dataDir := getenv("DATA_DIR", "/app/data")
	counterFile := filepath.Join(dataDir, "counter.txt")

	// ensure data dir exists
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		log.Fatalf("failed to create data dir: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// read counter
		n := readCounter(counterFile)
		n++
		if err := os.WriteFile(counterFile, []byte(strconv.Itoa(n)), 0o644); err != nil {
			http.Error(w, "failed to write counter", 500)
			return
		}
		fmt.Fprintf(w, "%s 👋\n", greeting)
		fmt.Fprintf(w, "Hits: %d\n", n)
		fmt.Fprintf(w, "Version: %s\n", version)
	})

	log.Printf("listening on :%s …", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func readCounter(path string) int {
	b, err := os.ReadFile(path)
	if err != nil {
		return 0
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return 0
	}
	return i
}
