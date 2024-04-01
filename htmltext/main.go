package main

import (
	"os"
	"strings"
	"text/template"
)

// func main() {
// 	for _, p := range Products {
// 		Printfln("Product: %v, Category: %v, Price: $%.2f", p.Name, p.Category, p.Price)
// 	}
// }

// func main() {
// 	t, err := template.ParseFiles("templates/template.html")
// 	if err == nil {
// 		t.Execute(os.Stdout, &Kayak)
// 	} else {
// 		Printfln("Error: %v", err.Error())
// 	}
// }

/* Загрузка нескольких шаблонов*/

// func main() {
// 	t1, err1 := template.ParseFiles("templates/template.html")
// 	t2, err2 := template.ParseFiles("templates/extras.html")
// 	if err1 == nil && err2 == nil {
// 		t1.Execute(os.Stdout, &Kayak)
// 		os.Stdout.WriteString("\n")
// 		t2.Execute(os.Stdout, &Kayak)
// 	} else {
// 		Printfln("Error: %v %v", err1.Error(), err2.Error())
// 	}
// }

/*Альтернатива*/

// func main() {
// 	allTemplates, err := template.ParseFiles("templates/template.html", "templates/extras.html")
// 	if err == nil {
// 		allTemplates.ExecuteTemplate(os.Stdout, "template.html", &Kayak)
// 		os.Stdout.WriteString("\n")
// 		allTemplates.ExecuteTemplate(os.Stdout, "extras.html", &Kayak)
// 	} else {
// 		Printfln("Error: %v %v", err.Error())
// 	}
// }

/*пречисление загруженных шаблонов*/

// func main() {
// 	allTemplates, err := template.ParseGlob("templates/*.html")
// 	if err == nil {
// 		for _, t := range allTemplates.Templates() {
// 			Printfln("template name:%v", t.Name())
// 		}
// 	} else {
// 		Printfln("Error: %v %v", err.Error())
// 	}
// }

/* Поиск определенного шаблона*/

// func Exec(t *template.Template) error {
// 	return t.Execute(os.Stdout, &Kayak)
// }

// func main() {
// 	allTemplates, err := template.ParseGlob("templates/*.html")
// 	if err == nil {
// 		selectedTemplated := allTemplates.Lookup("template.html")
// 		err = Exec(selectedTemplated)
// 	} else {
// 		Printfln("Error: %v %v", err.Error())
// 	}
// }

// func Exec(t *template.Template) error {
// 	return t.Execute(os.Stdout, Products)
// }

// func main() {
// 	allTemplates, err := template.ParseGlob("templates/*.html")
// 	if err == nil {
// 		selectedTemplated := allTemplates.Lookup("template.html")
// 		err = Exec(selectedTemplated)
// 	} else {
// 		Printfln("Error: %v %v", err.Error())
// 	}
// }

/*создание именнованых вложенных шаблонов*/
// func Exec(t *template.Template) error {
// 	return t.Execute(os.Stdout, Products)
// }

// func main() {
// 	allTemplates, err := template.ParseGlob("templates/*.html")
// 	if err == nil {
// 		selectedTemplated := allTemplates.Lookup("mainTemplate")
// 		err = Exec(selectedTemplated)
// 	} else {
// 		Printfln("Error: %v %v", err.Error())
// 	}
// }

/* определение блоков шаблона*/

// func Exec(t *template.Template) error {
// 	return t.Execute(os.Stdout, Products)
// }

// func main() {
// 	allTemplates, err := template.ParseFiles("templates/template.html", "templates/list.html")
// 	if err == nil {
// 		selectedTemplated := allTemplates.Lookup("mainTemplate")
// 		err = Exec(selectedTemplated)
// 	} else {
// 		Printfln("Error: %v %v", err.Error())
// 	}
// }

/*Определение функций шаблона*/

// func GetCategories(products []Product) (categories []string) {
// 	catMap := map[string]string{}
// 	for _, p := range products {
// 		if catMap[p.Category] == "" {
// 			catMap[p.Category] = p.Category
// 			categories = append(categories, p.Category)
// 		}
// 	}
// 	return
// }

// func Exec(t *template.Template) error {
// 	return t.Execute(os.Stdout, Products)

// }

// func main() {
// 	allTemplates := template.New("allTemplates")
// 	allTemplates.Funcs(map[string]interface{}{
// 		"getCats": GetCategories,
// 	})
// 	allTemplates, err := allTemplates.ParseGlob("templates/*.html")
// 	if err == nil {
// 		selectedTemplated := allTemplates.Lookup("mainTemplate")
// 		err = Exec(selectedTemplated)
// 	} else {
// 		Printfln("Error: %v %v", err.Error())
// 	}
// }

/*Отключение кодирования результата функции*/

// func GetCategories(products []Product) (categories []template.HTML) {
// 	catMap := map[string]string{}
// 	for _, p := range products {
// 		if catMap[p.Category] == "" {
// 			catMap[p.Category] = p.Category
// 			categories = append(categories, "<b>p.Category</b>")
// 		}
// 	}
// 	return
// }

// func Exec(t *template.Template) error {
// 	return t.Execute(os.Stdout, Products)

// }

// func main() {
// 	allTemplates := template.New("allTemplates")
// 	allTemplates.Funcs(map[string]interface{}{
// 		"getCats": GetCategories,
// 	})
// 	allTemplates, err := allTemplates.ParseGlob("templates/*.html")
// 	if err == nil {
// 		selectedTemplated := allTemplates.Lookup("mainTemplate")
// 		err = Exec(selectedTemplated)
// 	} else {
// 		Printfln("Error: %v %v", err.Error())
// 	}
// }

/*Представление доступа к функциям стандартной библиотеки*/

// func GetCategories(products []Product) (categories []string) {
// 	catMap := map[string]string{}
// 	for _, p := range products {
// 		if catMap[p.Category] == "" {
// 			catMap[p.Category] = p.Category
// 			categories = append(categories, p.Category)
// 		}
// 	}
// 	return
// }

// func Exec(t *template.Template) error {
// 	return t.Execute(os.Stdout, Products)

// }

// func main() {
// 	allTemplates := template.New("allTemplates")
// 	allTemplates.Funcs(map[string]interface{}{
// 		"getCats": GetCategories,
// 		"lower": strings.ToLower,
// 	})
// 	allTemplates, err := allTemplates.ParseGlob("templates/*.html")
// 	if err == nil {
// 		selectedTemplated := allTemplates.Lookup("mainTemplate")
// 		err = Exec(selectedTemplated)
// 	} else {
// 		Printfln("Error: %v %v", err.Error())
// 	}
// }

/*Использование переменных шаблона в действиях диапазона*/

// func GetCategories(products []Product) (categories []string) {
// 	catMap := map[string]string{}
// 	for _, p := range products {
// 		if catMap[p.Category] == "" {
// 			catMap[p.Category] = p.Category
// 			categories = append(categories, p.Category)
// 		}
// 	}
// 	return
// }

// func Exec(t *template.Template) error {
// 	productMap := map[string]Product{}
// 	for _, p := range Products {
// 		productMap[p.Name] = p
// 	}
// 	return t.Execute(os.Stdout, &productMap)

// }

// func main() {
// 	allTemplates := template.New("allTemplates")
// 	allTemplates.Funcs(map[string]interface{}{
// 		"getCats": GetCategories,
// 		"lower":   strings.ToLower,
// 	})
// 	allTemplates, err := allTemplates.ParseGlob("templates/*.html")
// 	if err == nil {
// 		selectedTemplated := allTemplates.Lookup("mainTemplate")
// 		err = Exec(selectedTemplated)
// 	} else {
// 		Printfln("Error: %v %v", err.Error())
// 	}
// }

/*Создание текстовых шаблонов*/

func GetCategories(products []Product) (categories []string) {
	catMap := map[string]string{}
	for _, p := range products {
		if catMap[p.Category] == "" {
			catMap[p.Category] = p.Category
			categories = append(categories, p.Category)
		}
	}
	return
}

func Exec(t *template.Template) error {
	productMap := map[string]Product{}
	for _, p := range Products {
		productMap[p.Name] = p
	}
	return t.Execute(os.Stdout, &productMap)

}

func main() {
	allTemplates := template.New("allTemplates")
	allTemplates.Funcs(map[string]interface{}{
		"getCats": GetCategories,
		"lower":   strings.ToLower,
	})
	allTemplates, err := allTemplates.ParseGlob("templates/*.txt")
	if err == nil {
		selectedTemplated := allTemplates.Lookup("mainTemplate")
		err = Exec(selectedTemplated)
	} else {
		Printfln("Error: %v %v", err.Error())
	}
}
