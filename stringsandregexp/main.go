package main

import "fmt"

/*  СРАВНЕНИЕ СТРОК  */

//func main() {
//	product := "Kayak"
//
//	fmt.Println("Product:", product)
//
//	/* Contains(s, substr) - возвращует tru если в строке s содержится substr и false если нет*/
//	fmt.Println("Contains", strings.Contains(product, "yak"))
//	fmt.Println("Contains", strings.Contains(product, "yk"))
//
//	/* ContainsAny(s, substr) - возвращует tru если в строке s содержится любой из символов substr*/
//	fmt.Println("ContainsAny", strings.ContainsAny(product, "abc"))
//	fmt.Println("ContainsAny", strings.ContainsAny(product, "tbc"))
//
//	/* ContainsRune(s, rune) - возвращует tru если в строке s содержится определенная руна*/
//	fmt.Println("ContainsRune", strings.ContainsRune(product, 'k'))
//	fmt.Println("ContainsRune", strings.ContainsRune(product, 't'))
//
//	/* EqualFold(s1, s2) - производит сравнение без учета регистра*/
//	fmt.Println("EqualFold", strings.EqualFold(product, "KAYAK"))
//
//	/* HasPrefix(s1, prefix) - возвращает true если строка начинается с указанного префикса*/
//	fmt.Println("HasPrefix", strings.HasPrefix(product, "Ka"))
//
//	/* HasSuffix(s1, suffix) - возвращает true если строка заканчивается указанным суффиксом*/
//	fmt.Println("HasSuffix", strings.HasSuffix(product, "yak"))
//}

/*  ПРЕОБРАЗОВАНИЕ РЕГИСТРА СТРОК  */

//func main() {
//	description := "A bout for sailing"
//	descriptionRus := "Готовлюсь к отплытию"
//
//	fmt.Println("Original", description)
//	fmt.Println("Оригинал", descriptionRus)
//
//	/* ToLower(str) возвращает новую строку символы преобразует в нижний регистр */
//	fmt.Println("ToLower", strings.ToLower(description))
//	fmt.Println("ToLower", strings.ToLower(descriptionRus))
//
//	/* ToUpper(str) возвращает новую строку символы преобразует в верхний регистр */
//	fmt.Println("ToUpper", strings.ToUpper(description))
//	fmt.Println("ToUpper", strings.ToUpper(descriptionRus))
//
//	/* Title(str) возвращает новую строку первые символы каждого слова преобразованы в верхний регистр */
//	fmt.Println("Title", strings.Title(description))
//	fmt.Println("Title", strings.Title(descriptionRus))
//
//	/* ToTitle(str) возвращает новую строку символы преобразует в верхний регистр */
//	fmt.Println("ToTitle", strings.ToTitle(description))
//	fmt.Println("ToTitle", strings.ToTitle(descriptionRus))
//
//	/* НЮАНСЫ */
//
//	specialChar := "\u01c9"
//	fmt.Println("Original", specialChar, []byte(specialChar))
//
//	upperChar := strings.ToUpper(specialChar)
//	fmt.Println("Upper", upperChar, []byte(upperChar))
//
//	titleChar := strings.ToTitle(specialChar)
//	fmt.Println("Title", titleChar, []byte(titleChar))
//}

/*  РАБОТА С РЕГИСТРОМ СИМВОЛОВ  */
func main() {
	product := "Kayak"
	fmt.Println(product)
}
