package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dokyan1989/g1/app/web1/config"
)

var cfg *config.Config

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	var err error

	port := 8080
	cfg, err = config.Load()
	if err != nil {
		return err
	}

	r, err := newRouter(cfg)
	if err != nil {
		return err
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		return err
	}

	return nil
}
