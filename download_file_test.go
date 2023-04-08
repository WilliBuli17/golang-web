package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")

	if file == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
		return
	}

	//code di line 19 ini hanya opsional, gunanya jika user mengakses file, maka file tidak akan di render di web, tapi langsug di download
	w.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"")

	http.ServeFile(w, r, "./upload/"+file)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
