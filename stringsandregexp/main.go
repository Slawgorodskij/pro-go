package main

import (
	"fmt"
	"strings"
)

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

/*  РАБОТА С РЕГИСТРОМ СИМВОЛОВ  (пакет unicode)  */
//func main() {
//	product := "Kayak"
//	/*IsLower(rune) возвращает tru если указанная руна в нижнем регистре*/
//	for key, char := range product {
//		fmt.Println(string(char), key, "IsLower():", unicode.IsLower(char))
//	}
//
//	/*ToLower(rune) возвращает строчную руну*/
//	for key, char := range product {
//		fmt.Println(string(char), key, "ToLower():", string(unicode.ToLower(char)))
//	}
//
//	/*IsUpper(rune) возвращает tru если указанная руна в верхнем регистре*/
//	/*ToUpper(rune) возвращает руну в верхнем регистре*/
//
//	/*IsTitle(rune) возвращает tru если указанная руна является заглавной*/
//	for _, char := range product {
//		fmt.Println(string(char), "IsTitle():", unicode.IsTitle(char))
//	}
//
//	/*ToTitle(rune) возвращает руну в заглавном регистре*/
//	for _, char := range product {
//		fmt.Println(string(char), "ToTitle():", string(unicode.ToTitle(char)))
//	}
//}

/* ПРОВЕРКА СТРОК */
//func main() {
//	description := "A boat for one person"
//
//	/* Count(s, sub) функция возвращает int, который сообщает, сколько раз указанная подстрока встречается в строке s */
//	fmt.Println("Count:", strings.Count(description, "o"))
//
//	/* Index(s, sub) функция возвращает индекс первого вхождения указанной подстроки в строке s. -1 если вхождения нет */
//	fmt.Println("Index:", strings.Index(description, "o"))
//
//	/* LastIndex(s, sub) функция возвращает индекс последнего вхождения указанной подстроки в строке s. -1 если вхождения нет */
//	fmt.Println("LastIndex:", strings.LastIndex(description, "o"))
//
//	/* IndexAny(s, sub) функция возвращает индекс первого вхождения любого символа укзанной подстроки в строке s. -1 если вхождения нет */
//	fmt.Println("IndexAny:", strings.IndexAny(description, "o"))
//
//	/* LastIndexAny(s, sub) функция возвращает индекс последнего вхождения любого символа указанной подстроки в строке s. -1 если вхождения нет */
//	fmt.Println("LastIndexAny:", strings.LastIndexAny(description, "o"))
//
//	/* IndexByte(s, b) функция возвращает индекс первого вхождения указанного byte в строке s. -1 если вхождения нет */
//	/* LastIndexByte(s, b) функция возвращает индекс последнего вхождения указанного byte в строке s. -1 если вхождения нет */
//
//	/* Проверка строк с помощью пользовательской функций */
//	isLetterB := func(r rune) bool {
//		return r == 'B' || r == 'b'
//	}
//
//	/* IndexFunc(s, func) функция возвращает индекс первого вхождения символа в строку s, для которого указанная функция возвращает true */
//	fmt.Println("IndexFunc:", strings.IndexFunc(description, isLetterB))
//
//	/* LastIndexFunc(s, func) функция возвращает индекс последнего вхождения символа в строку s, для которого указанная функция возвращает true */
//	fmt.Println("LastIndexFunc:", strings.LastIndexFunc(description, isLetterB))
//}

/*  МАНИПУЛИРОВАНИЕ СТРОКАМИ  */

/* Разделение строк */

func main() {
	description := "A boat for one person"
	/* Fields(s) разбивает строку на пробельные символы и возвращает срез, содержащий непробельные разделы строки s */
	/* FieldsFunc(s, func) разбивает строку на символы, для которых пользовательская функция возвращает tru, и возвращает срез, содержащий оставшиеся части строки s */

	/* Split(s, sub) разбивает строку s на каждое вхождение указанной подстроки, возвращая срез. Если разделителем является пустая строка, то срез будет содержать строки для каждого символа */
	splits := strings.Split(description, " ")
	for _, x := range splits {
		fmt.Println("Split >>" + x + "<<")
	}
	/* SplitN(s, sub, max) похожа на Split, но принимает дополнительный аргумент типа int указывающий максимальное количество возвращаемых строк. Последняя подстрока будет содержать не разделенную часть исходной строки */

	/* SplitAfter(s, sub) похожа на Split, но включает строку используемую в результатах */
	splitsAfter := strings.SplitAfter(description, " ")
	for _, x := range splitsAfter {
		fmt.Println("SplitAfter >>" + x + "<<")
	}
	/* SplitAfterN(s, sub, max) похожа на SplitAfter, но принимает дополнительный аргумент типа int, указывающий максимальное количество возвращаемых строк */
}
