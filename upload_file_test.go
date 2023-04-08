package golang_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UplpadForm(w http.ResponseWriter, r *http.Request) {
	onceInitTemplate.Do(func() {
		templateMutex.Lock()
		defer templateMutex.Unlock()
		myTemplates = initTemplates()
	})

	err := myTemplates.ExecuteTemplate(w, "upload_form", map[string]interface{}{
		"Title":  "Template Upload File",
		"Action": "Upload File",
	})

	if err != nil {
		panic(err)
	}
}

func Upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) //Set max upload file size jadi 10 mb
	if err != nil {
		panic(err)
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	fileDestinasion, err := os.Create("./upload/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileDestinasion, file)
	if err != nil {
		panic(err)
	}

	// sampai disini itu d=sudah cukup untuk ngambil file dari hasil upload dan masukkan ke penyimpanan
	// berikut adalah contoh jika ingin ngambil value lainnya

	name := r.PostFormValue("name")

	onceInitTemplate.Do(func() {
		templateMutex.Lock()
		defer templateMutex.Unlock()
		myTemplates = initTemplates()
	})
	err = myTemplates.ExecuteTemplate(w, "upload_success", map[string]interface{}{
		"Title":  "Template Upload File",
		"Action": "Upload File Sukses",
		"Name":   name,
		"File":   "/static/" + fileHeader.Filename,
	})
	if err != nil {
		panic(err)
	}
}

func TestUploadForm(t *testing.T) { // ini cuma akan run di server, bukan unit test asli
	mux := http.NewServeMux()
	mux.HandleFunc("/", UplpadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./upload"))))

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed upload/Golang.png
var fileUpload []byte

func TestUploadFile(t *testing.T) { // ini unit test yang benar
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	err := writer.WriteField("name", "Willi Buli")
	if err != nil {
		panic(err)
	}

	file, err := writer.CreateFormFile("file", "Golang-test.png")
	if err != nil {
		panic(err)
	}
	_, err = file.Write(fileUpload)
	if err != nil {
		panic(err)
	}
	err = writer.Close()
	if err != nil {
		panic(err)
	}

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	response := recorder.Result()
	bodyResponse, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bodyResponse))
}
