package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// func main() {
// 	for _, p := range Products {
// 		Printfln("Product: %v, Category: %v, Price: $%.2f", p.Name, p.Category, p.Price)
// 	}
// }

/*Использование функции удобства записи*/

// func main() {
// 	total := 0.0
// 	for _, p := range Products {
// 		total += p.Price
// 	}

// 	dataStr := fmt.Sprintf("Time: %v, Total: $%.2f\n", time.Now().Format("Mon 15:04:05"), total)

// 	file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

// 	if err == nil {
// 		defer file.Close()
// 		file.WriteString(dataStr)
// 	} else {
// 		Printfln("error: %v", err.Error())
// 	}
// }

/*Запись данных JSON в файл*/
// func main() {
// 	cheapProducts := []Product{}
// 	for _, p := range Products {
// 		if p.Price < 100 {
// 			cheapProducts = append(cheapProducts, p)
// 		}
// 	}

// 	file, err := os.OpenFile("cheap.json", os.O_WRONLY|os.O_CREATE, 0666)
// 	if err == nil {
// 		defer file.Close()
// 		encoder := json.NewEncoder(file)
// 		encoder.Encode(cheapProducts)
// 	} else {
// 		Printfln("Error: %v", err.Error())
// 	}
// }

/*Использование удобных функции для создания новых файлов*/
// func main() {
// 	cheapProducts := []Product{}
// 	for _, p := range Products {
// 		if p.Price < 100 {
// 			cheapProducts = append(cheapProducts, p)
// 		}
// 	}

// 	file, err := os.CreateTemp(".", "tempfile-*.json")
// 	if err == nil {
// 		defer file.Close()
// 		encoder := json.NewEncoder(file)
// 		encoder.Encode(cheapProducts)
// 	} else {
// 		Printfln("Error: %v", err.Error())
// 	}
// }

/*Работа с путями к файлам*/

// func main() {
// path, err := os.UserHomeDir()
// if err == nil {
// 	path = filepath.Join(path, "myApp", "MyTempFile.json")
// }

// 	Printfln("Full path: %v", path)
// 	Printfln("Volume name: %v", filepath.VolumeName(path))
// 	Printfln("Dir component: %v", filepath.Dir(path))
// 	Printfln("File component: %v", filepath.Base(path))
// 	Printfln("File extension: %v", filepath.Ext(path))
// }

/*Управление файлами и каталогами*/

func main() {
	path, err := os.UserHomeDir()
	if err == nil {
		path = filepath.Join(path, "project/pro-go/files/myApp", "MyTempFile.json")
	}
	Printfln("Full path: %v", path)

	err = os.MkdirAll(filepath.Dir(path), 0766)

	if err == nil {
		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			defer file.Close()
			encoder := json.NewEncoder(file)
			encoder.Encode(Products)
		}
	}
	if err != nil {
		Printfln("error %v", err.Error())
	}
}
