package rest

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// RunServer runs HTTP/REST gateway
func RunServer(ctx context.Context, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	r := mux.NewRouter()

	r.HandleFunc("/api/status", status).Methods("GET")

	//addr := ":" + app.Config.Port
	//log.Println("listen on", addr)
	//if err := http.ListenAndServe(addr, r); err != nil {
	//	log.Fatal(err)
	//}

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: r,
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			logger.Log.Warn("shutting down HTTP/REST gateway...")
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	logger.Log.Info("starting HTTP/REST gateway...")
	return srv.ListenAndServe()
}

type statusResponse struct {
	Status string `bson:"status" json:"status"`
}

func status(w http.ResponseWriter, r *http.Request) {
	var data = statusResponse{"ok"}
	response.RespondWithJson(w, http.StatusOK, data)
}
