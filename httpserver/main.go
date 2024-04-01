package main

import (
	"io"
	"net/http"
	"strings"
)

// func main(){
// 	for _, p := range Products {
// 		Printfln("Product: %v, Category: %v, Price: $%.2f", p.Name, p.Category, p.Price)
// 	}
// }

/*Создание простого HTTP-сервера*/

// type StringHandler struct {
// 	message string
// }

// func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
// 	io.WriteString(writer, sh.message)
// }

// func main() {
// 	err := http.ListenAndServe(":5001", StringHandler{message: "Hello, Word"})
// 	if err != nil {
// 		Printfln("Error: %v", err.Error())
// 	}
// }

/* Проврека запроса*/

// type StringHandler struct {
// 	message string
// }

// func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
// 	Printfln("Method: %v", request.Method)
// 	Printfln("URL: %v", request.URL)
// 	Printfln("HTTP Version: %v", request.Proto)
// 	Printfln("Host: %v", request.Host)
// 	for name, val := range request.Header {
// 		Printfln("Header: %v, Value: %v", name, val)
// 	}
// 	Printfln("---")
// 	io.WriteString(writer, sh.message)
// }

// func main() {
// 	err := http.ListenAndServe(":5001", StringHandler{message: "Hello, Word"})
// 	if err != nil {
// 		Printfln("Error: %v", err.Error())
// 	}
// }

/*Фильтрация запросов и генерация ответов*/

// type StringHandler struct {
// 	message string
// }

// func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
// 	if request.URL.Path == "/favicon.ico" {
// 		Printfln("Request for icon detected -returning 400")
// 		writer.WriteHeader(http.StatusNotFound)
// 		return
// 	}
// 	Printfln("Request for %v", request.URL.Path)
// 	io.WriteString(writer, sh.message)
// }

// func main() {
// 	err := http.ListenAndServe(":5001", StringHandler{message: "Hello, Word"})
// 	if err != nil {
// 		Printfln("Error: %v", err.Error())
// 	}
// }

/*Использование удобных функций ответа*/

// type StringHandler struct {
// 	message string
// }

// func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
// 	Printfln("Request for %v", request.URL.Path)
// 	switch request.URL.Path {
// 	case "/favicon.ico":
// 		http.NotFound(writer, request)
// 	case "/message":
// 		io.WriteString(writer, sh.message)
// 	default:
// 		http.Redirect(writer, request, "/message", http.StatusTemporaryRedirect)
// 	}
// 	io.WriteString(writer, sh.message)
// }

// func main() {
// 	err := http.ListenAndServe(":5001", StringHandler{message: "Hello, Word"})
// 	if err != nil {
// 		Printfln("Error: %v", err.Error())
// 	}
// }

/*Использование обработцича удобной маршрутизации*/

// type StringHandler struct {
// 	message string
// }

// func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
// 	Printfln("Request for %v", request.URL.Path)
// 	io.WriteString(writer, sh.message)
// }

// func main() {
// 	http.Handle("/message", StringHandler{"Hello, Word"})
// 	http.Handle("/favicon.ico", http.NotFoundHandler())
// 	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
// 	err := http.ListenAndServe(":5001", nil)
// 	if err != nil {
// 		Printfln("Error: %v", err.Error())
// 	}
// }

/*Поддержка HTTPS-запросов*/

// type StringHandler struct {
// 	message string
// }

// func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
// 	Printfln("Request for %v", request.URL.Path)
// 	io.WriteString(writer, sh.message)
// }

// func main() {
// 	http.Handle("/message", StringHandler{"Hello, Word"})
// 	http.Handle("/favicon.ico", http.NotFoundHandler())
// 	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
// 	go func() {
// 		err := http.ListenAndServeTLS(":5500", "certificate.cer", "certificate.pkey", nil)
// 		if err != nil {
// 			Printfln("HTTPS Error: %v", err.Error())
// 		}
// 	}()
// 	err := http.ListenAndServe(":5001", nil)
// 	if err != nil {
// 		Printfln("Error: %v", err.Error())
// 	}
// }

/*пренаправление  HTTP-запросов на HTTPS*/

// type StringHandler struct {
// 	message string
// }

// func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
// 	Printfln("Request for %v", request.URL.Path)
// 	io.WriteString(writer, sh.message)
// }

// func HTTPSRedirect(writer http.ResponseWriter, request *http.Request) {
// 	host := strings.Split(request.Host, ":")[0]
// 	target := "https://" + host + ":5500" + request.URL.Path
// 	if len(request.URL.RawQuery) > 0 {
// 		target += "?" + request.URL.RawQuery
// 	}
// 	http.Redirect(writer, request, target, http.StatusTemporaryRedirect)
// }

// func main() {
// 	http.Handle("/message", StringHandler{"Hello, Word"})
// 	http.Handle("/favicon.ico", http.NotFoundHandler())
// 	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
// 	go func() {
// 		err := http.ListenAndServeTLS(":5500", "certificate.cer", "certificate.pkey", nil)
// 		if err != nil {
// 			Printfln("HTTPS Error: %v", err.Error())
// 		}
// 	}()
// 	err := http.ListenAndServe(":5001", http.HandlerFunc(HTTPSRedirect))
// 	if err != nil {
// 		Printfln("Error: %v", err.Error())
// 	}
// }

/*Создание статического HTTP-сервера*/

type StringHandler struct {
	message string
}

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	Printfln("Request for %v", request.URL.Path)
	io.WriteString(writer, sh.message)
}

func HTTPSRedirect(writer http.ResponseWriter, request *http.Request) {
	host := strings.Split(request.Host, ":")[0]
	target := "https://" + host + ":5500" + request.URL.Path
	if len(request.URL.RawQuery) > 0 {
		target += "?" + request.URL.RawQuery
	}
	http.Redirect(writer, request, target, http.StatusTemporaryRedirect)
}

func main() {
	http.Handle("/message", StringHandler{"Hello, Word"})
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))

	fsHandler := http.FileServer(http.Dir("./static"))
	http.Handle("/files/", http.StripPrefix("/files", fsHandler))

	go func() {
		err := http.ListenAndServeTLS(":5500", "certificate.cer", "certificate.pkey", nil)
		if err != nil {
			Printfln("HTTPS Error: %v", err.Error())
		}
	}()
	err := http.ListenAndServe(":5001", http.HandlerFunc(HTTPSRedirect))
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}
