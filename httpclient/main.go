package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

// func main() {
// 	Printfln("Starting HTTP Server")
// 	http.ListenAndServe(":5000", nil)
// }

// func main() {
// 	go http.ListenAndServe(":5000", nil)
// 	time.Sleep(time.Second)

// 	response, err := http.Get("http://localhost:5000/html")
// 	if err == nil {
// 		response.Write(os.Stdout)
// 	} else {
// 		Printfln("Error: %v", err.Error())
// 	}
// }

// func main() {
// 	go http.ListenAndServe(":5000", nil)
// 	time.Sleep(time.Second)

// 	response, err := http.Get("http://localhost:5000/html")
// 	if err == nil && response.StatusCode == http.StatusOK {
// 		data, err := io.ReadAll(response.Body)
// 		if err == nil {
// 			defer response.Body.Close()
// 			os.Stdout.Write(data)
// 		}
// 	} else {
// 		Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
// 	}
// }

// func main() {
// 	go http.ListenAndServe(":5000", nil)
// 	time.Sleep(time.Second)

// 	response, err := http.Get("http://localhost:5000/json")
// 	if err == nil && response.StatusCode == http.StatusOK {
// 		defer response.Body.Close()
// 		data := []Product{}
// 		err = json.NewDecoder(response.Body).Decode(&data)
// 		if err == nil {
// 			for _, p := range data {
// 				Printfln("Name: %v, Price: $%.2f", p.Name, p.Price)
// 			}
// 		} else {
// 			Printfln("Decode error: %v", err.Error())
// 		}
// 	} else {
// 		Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
// 	}
// }

/*Отправка пост запросов*/

// func main() {
// 	go http.ListenAndServe(":5000", nil)
// 	time.Sleep(time.Second)

// 	formData := map[string][]string {
// 		"name": {"Kayak"},
// 		"category": {"Waterspport"},
// 		"price": {"279"},
// 	}

// 	response, err := http.PostForm("http://localhost:5000/echo", formData)
// 	if err == nil && response.StatusCode == http.StatusOK {
// 		io.Copy(os.Stdout, response.Body)
// 		defer response.Body.Close()
// 	} else {
// 		Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
// 	}
// }

/* Публикация формы с помощью ридера*/

// func main() {
// 	go http.ListenAndServe(":5000", nil)
// 	time.Sleep(time.Second)

// 	var builder strings.Builder

// 	err := json.NewEncoder(&builder).Encode(Products[0])
// 	if err == nil {
// 		response, err := http.Post("http://localhost:5000/echo", "application/json", strings.NewReader(builder.String()))
// 		if err == nil && response.StatusCode == http.StatusOK {
// 			io.Copy(os.Stdout, response.Body)
// 			defer response.Body.Close()
// 		} else {
// 			Printfln("Error: %v", err.Error())
// 		}
// 	} else {
// 		Printfln("Error: %v", err.Error())
// 	}
// }

/*настройка запросов HTTP-клиента*/

// func main() {
// 	go http.ListenAndServe(":5000", nil)
// 	time.Sleep(time.Second)

// 	var builder strings.Builder
// 	err := json.NewEncoder(&builder).Encode(Products[0])
// 	if err == nil {
// 		reqURL, err := url.Parse("http://localhost:5000/echo")
// 		if err == nil {
// 			req := http.Request{
// 				Method: http.MethodPost,
// 				URL:    reqURL,
// 				Header: map[string][]string{
// 					"Content-Type": {"application.json"},
// 				},
// 				Body: io.NopCloser(strings.NewReader(builder.String())),
// 			}
// 			response, err := http.DefaultClient.Do(&req)
// 			if err == nil && response.StatusCode == http.StatusOK {
// 				io.Copy(os.Stdout, response.Body)
// 				defer response.Body.Close()
// 			} else {
// 				Printfln("Request Error: %v", err.Error())
// 			}
// 		} else {
// 			Printfln("Parse Error: %v", err.Error())
// 		}
// 	} else {
// 		Printfln("Encoder Error: %v", err.Error())
// 	}
// }

/*Использование удобных функциф для создания запроса*/

// func main() {
// 	go http.ListenAndServe(":5000", nil)
// 	time.Sleep(time.Second)

// 	var builder strings.Builder
// 	err := json.NewEncoder(&builder).Encode(Products[0])
// 	if err == nil {
// 		req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/echo", io.NopCloser(strings.NewReader(builder.String())))
// 		if err == nil {
// 			req.Header["Content-Type"] = []string{"application.json"}
// 			response, err := http.DefaultClient.Do(req)
// 			if err == nil && response.StatusCode == http.StatusOK {
// 				io.Copy(os.Stdout, response.Body)
// 				defer response.Body.Close()
// 			} else {
// 				Printfln("Request Error: %v", err.Error())
// 			}
// 		} else {
// 			Printfln("Request Init Error: %v", err.Error())
// 		}
// 	} else {
// 		Printfln("Encoder Error: %v", err.Error())
// 	}
// }

/*Работа с файлами cookie*/

// func main() {
// 	go http.ListenAndServe(":5000", nil)
// 	time.Sleep(time.Second)

// 	jar, err := cookiejar.New(nil)
// 	if err == nil {
// 		http.DefaultClient.Jar = jar
// 	}

// 	for i := 0; i < 3; i++ {
// 		req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/cookie", nil)
// 		if err == nil {
// 			response, err := http.DefaultClient.Do(req)
// 			if err == nil && response.StatusCode == http.StatusOK {
// 				io.Copy(os.Stdout, response.Body)
// 				defer response.Body.Close()
// 			} else {
// 				Printfln("Request Error: %v", err.Error())
// 			}
// 		} else {
// 			Printfln("Request Init Error: %v", err.Error())
// 		}
// 	}
// }

/*Создание отдельных клиентов и файлов cookie*/

// три отдельных значения Client
// func main() {
// 	go http.ListenAndServe(":5000", nil)
// 	time.Sleep(time.Second)

// 	clients := make([]http.Client, 3)
// 	for index, client := range clients {
// 		jar, err := cookiejar.New(nil)
// 		if err == nil {
// 			client.Jar = jar
// 		}

// 		for i := 0; i < 3; i++ {
// 			req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/cookie", nil)
// 			if err == nil {
// 				response, err := client.Do(req)
// 				if err == nil && response.StatusCode == http.StatusOK {
// 					fmt.Fprintf(os.Stdout, "Client %v: ", index)
// 					io.Copy(os.Stdout, response.Body)
// 					defer response.Body.Close()
// 				} else {
// 					Printfln("Request Error: %v", err.Error())
// 				}
// 			} else {
// 				Printfln("Request Init Error: %v", err.Error())
// 			}
// 		}
// 	}

// }

// с одним файлом CookieJar
// func main() {
// 	go http.ListenAndServe(":5000", nil)
// 	time.Sleep(time.Second)

// 	jar, err := cookiejar.New(nil)

// 	clients := make([]http.Client, 3)
// 	for index, client := range clients {
// 		if err == nil {
// 			client.Jar = jar
// 		}

// 		for i := 0; i < 3; i++ {
// 			req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/cookie", nil)
// 			if err == nil {
// 				response, err := client.Do(req)
// 				if err == nil && response.StatusCode == http.StatusOK {
// 					fmt.Fprintf(os.Stdout, "Client %v: ", index)
// 					io.Copy(os.Stdout, response.Body)
// 					defer response.Body.Close()
// 				} else {
// 					Printfln("Request Error: %v", err.Error())
// 				}
// 			} else {
// 				Printfln("Request Init Error: %v", err.Error())
// 			}
// 		}
// 	}
// }

/* Управление перенаправлениями */

// func main() {
// 	go http.ListenAndServe(":5000", nil)
// 	time.Sleep(time.Second)

// 	req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/redirect1", nil)
// 	if err == nil {
// 		var response *http.Response
// 		response, err := http.DefaultClient.Do(req)
// 		if err == nil {
// 			io.Copy(os.Stdout, response.Body)
// 		} else {
// 			Printfln("Request Error: %v", err.Error())
// 		}
// 	} else {
// 		Printfln("Error: %v", err.Error())
// 	}

// }

//  при отработке вышерасположенного кода мы получим Request Error: Get "/redirect1": stopped after 10 redirects

// func main() {
// 	go http.ListenAndServe(":5000", nil)
// 	time.Sleep(time.Second)

// 	http.DefaultClient.CheckRedirect = func(req *http.Request, previos []*http.Request) error {
// 		if len(previos) == 3 {
// 			url, _ := url.Parse("http://localhost:5000/html")
// 			req.URL = url
// 		}
// 		return nil
// 	}

// 	req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/redirect1", nil)
// 	if err == nil {
// 		var response *http.Response
// 		response, err := http.DefaultClient.Do(req)
// 		if err == nil {
// 			io.Copy(os.Stdout, response.Body)
// 		} else {
// 			Printfln("Request Error: %v", err.Error())
// 		}
// 	} else {
// 		Printfln("Error: %v", err.Error())
// 	}

// }
// в этом случае после переходов между redirect1 и redirect2 перейдем на страницу html

/* Создание составных форм */

func main() {
	go http.ListenAndServe(":5000", nil)
	time.Sleep(time.Second)

	var buffer bytes.Buffer
	formWriter := multipart.NewWriter(&buffer)

	fieldWriter, err := formWriter.CreateFormField("name")
	if err == nil {
		io.WriteString(fieldWriter, "Alise")
	}

	fildWriter, err := formWriter.CreateFormField("city")
	if err == nil {
		io.WriteString(fildWriter, "New York")
	}

	fileWriter, err := formWriter.CreateFormFile("CodeFile", "printer.go")
	if err == nil {
		fileData, err := os.ReadFile("./printer.go")
		if err == nil {
			fileWriter.Write(fileData)
		}
	}

	formWriter.Close()
	req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/form", &buffer)
	req.Header["Content-Type"] = []string{formWriter.FormDataContentType()}
	if err == nil {
		var response *http.Response
		response, err := http.DefaultClient.Do(req)
		if err == nil {
			io.Copy(os.Stdout, response.Body)
		} else {
			Printfln("Request Error: %v", err.Error())
		}
	} else {
		Printfln("Error: %v", err.Error())
	}

}
