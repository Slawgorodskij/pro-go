package main

import (
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

// func main() {
// 	path, err := os.UserHomeDir()
// 	if err == nil {
// 		path = filepath.Join(path, "project/pro-go/files/myApp", "MyTempFile.json")
// 	}
// 	Printfln("Full path: %v", path)

// 	err = os.MkdirAll(filepath.Dir(path), 0766)

// 	if err == nil {
// 		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
// 		if err == nil {
// 			defer file.Close()
// 			encoder := json.NewEncoder(file)
// 			encoder.Encode(Products)
// 		}
// 	}
// 	if err != nil {
// 		Printfln("error %v", err.Error())
// 	}
// }

/*Изучение файловой системы*/

// func main() {
// 	path, err := os.Getwd()
// 	if err == nil {
// 		dirEntries, err := os.ReadDir(path)
// 		if err == nil {
// 			for _, dentry := range dirEntries {
// 				Printfln("Entry name: %v, IsDir: %v", dentry.Name(), dentry.IsDir())
// 			}
// 		}
// 	}
// 	if err != nil {
// 		Printfln("error %v", err.Error())
// 	}
// }

/*определение существования файла*/

// func main() {
// 	targetFiles := []string {"no_such_file.text", "config.json"}

// 	for _, name := range targetFiles {
// 		info, err := os.Stat(name)
// 		if os.IsNotExist(err) {
// 			Printfln("File does not exits: %v", name)
// 		} else if err != nil {
// 			Printfln("Other error: %v", err.Error())
// 		} else {
// 			Printfln("File: %v, Size: %v", info.Name(), info.Size())
// 		}
// 	}
// }

/*поиск файлов с помощью шаблона*/

// func main() {
// 	path, err := os.Getwd()
// 	if err == nil {
// 		matches, err := filepath.Glob(filepath.Join(path, "*.json"))
// 		if err == nil {
// 			for _, m := range matches {
// 				Printfln("Match: %v", m)
// 			}
// 		}
// 	}
// 	if err != nil {
// 		Printfln("Error %v", err.Error())
// 	}
// }

/*обработкка всех файлов в каталоге*/

func callback(path string, dir os.DirEntry, dirErr error) (err error) {
	info, _ := dir.Info()
	Printfln("Path %v, Size: %v", path, info.Size())
	return
}

func main() {
	path, err := os.Getwd()
	if err == nil {
		err = filepath.WalkDir(path, callback)
	} else {
		Printfln("Error %v", err.Error())
	}
}
