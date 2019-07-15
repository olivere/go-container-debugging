package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := runMain(); err != nil {
		os.Exit(1)
	}
}

func runMain() error {
	var (
		addr = flag.String("addr", envString("ADDR", ":8080"), "")
	)
	flag.Parse()

	http.HandleFunc("/", root)

	fmt.Printf("Listening on %s\n", *addr)
	srv := &http.Server{
		Addr: *addr,
	}
	return srv.ListenAndServe()
}

func envString(name, defaultValue string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return defaultValue
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Current time is: %s\n", time.Now())
}
