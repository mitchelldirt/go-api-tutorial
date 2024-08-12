package main

import (
	"fmt"
	"go-api-tutorial/internal/handlers"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service on localhost:8080")

	fmt.Println(`
      ___           ___                    ___           ___
     /\  \         /\  \                  /\  \         /\  \          ___
    /::\  \       /::\  \                /::\  \       /::\  \        /\  \
   /:/\:\  \     /:/\:\  \              /:/\:\  \     /:/\:\  \       \:\  \
  /:/  \:\  \   /:/  \:\  \            /::\~\:\  \   /::\~\:\  \      /::\__\
 /:/__/_\:\__\ /:/__/ \:\__\          /:/\:\ \:\__\ /:/\:\ \:\__\  __/:/\/__/
 \:\  /\ \/__/ \:\  \ /:/  /          \/__\:\/:/  / \/__\:\/:/  / /\/:/  /
  \:\ \:\__\    \:\  /:/  /                \::/  /       \::/  /  \::/__/
   \:\/:/  /     \:\/:/  /                 /:/  /         \/__/    \:\__\
    \::/  /       \::/  /                 /:/  /                    \/__/
     \/__/         \/__/                  \/__/
`)

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Error(err)
	}
}
