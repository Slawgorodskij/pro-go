package main

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

// func main() {
// 	//description := "A boat for one person"
// 	/* Fields(s) разбивает строку на пробельные символы и возвращает срез, содержащий непробельные разделы строки s */
// 	/* FieldsFunc(s, func) разбивает строку на символы, для которых пользовательская функция возвращает tru, и возвращает срез, содержащий оставшиеся части строки s */

// 	/* Split(s, sub) разбивает строку s на каждое вхождение указанной подстроки, возвращая срез. Если разделителем является пустая строка, то срез будет содержать строки для каждого символа */
// 	// splits := strings.Split(description, " ")
// 	// for _, x := range splits {
// 	// 	fmt.Println("Split >>" + x + "<<")
// 	// }
// 	/* SplitN(s, sub, max) похожа на Split, но принимает дополнительный аргумент типа int указывающий максимальное количество возвращаемых строк. Последняя подстрока будет содержать не разделенную часть исходной строки */

// 	/* SplitAfter(s, sub) похожа на Split, но включает строку используемую в результатах */
// 	// splitsAfter := strings.SplitAfter(description, " ")
// 	// for _, x := range splitsAfter {
// 	// 	fmt.Println("SplitAfter >>" + x + "<<")
// 	// }
// 	/* SplitAfterN(s, sub, max) похожа на SplitAfter, но принимает дополнительный аргумент типа int, указывающий максимальное количество возвращаемых строк */
// 	// splitsAfterN := strings.SplitAfterN(description, " ", 3)
// 	// for _, x := range splitsAfterN {
// 	// 	fmt.Println("Split >>" + x + "<<")
// 	// }

// 	/* Split, SplitN, SplitAfter, SplitAfterN не работают с повторяющимися последовательностями символов*/
// 	//description := "This  is  double  spaced"

// 	// splits := strings.SplitN(description, " ", 3)
// 	// for _, x := range splits {
// 	// 	fmt.Println("Split >>" + x + "<<")
// 	// }

// 	/* для обработки повторяющихся пробельных символов используется функция  Fields*/
// 	// splits := strings.Fields(description)
// 	// for _, x := range splits {
// 	// 	fmt.Println("Split >>" + x + "<<")
// 	// }
// 	// /* но она не поддерживает ограничение на количество результатов*/

// }

/* РАЗДЕЛЕНИЕ  с использованием пользовательской функции для разделния строк*/

// func main() {
// 	description := "This  is  double  spaced"
// 	splitter := func(r rune) bool {
// 		return r == ' '
// 	}
// 	splits := strings.FieldsFunc(description, splitter)
// 	for _, x := range splits {
// 		fmt.Println("Field >>" + x + "<<")
// 	}
// }

/*Обрезка строк*/

/*Обрезка пробелов - TrimSpace*/
// func main() {
// 	username := " Alice"
// 	trimmed := strings.TrimSpace(username)
// 	fmt.Println("Trimmed:", ">>" + trimmed +  "<<")
// }

/*Обрезка наборов символов - Trim, TrimLeft, TrimRight*/

// func main() {
// 	description := "A boat for one person"
// 	trimmed := strings.Trim(description, "Asno ")
// 	fmt.Println("Trimed:", trimmed)
// }

/*Обрезка подстрок - TrimPrefix, TrimSuffix*/

// func main() {
// 	description := "A boat for one person"
// 	prefixTrimmed := strings.TrimPrefix(description, "A boat")
// 	wrongPrefix := strings.TrimPrefix(description, "A hat ")
// 	fmt.Println("Trimed:", prefixTrimmed)
// 	fmt.Println("Not trimed:", wrongPrefix)
// }

/*Обрезка наборов символов - TrimFunc, TrimLeftFunc, TrimRightFunc */

// func main() {
// 	description := "A boat for one person"

// 	trimmer := func (r rune) bool {
// 		return r == 'A' || r == 'n'
// 	}

// 	trimmed := strings.TrimFunc(description, trimmer)
// 	fmt.Println("Trimed:", trimmed)
// }

/*Изменение строк Replace, ReplaceAll*/

// func main() {
// 	text := "It was a boat. A small boat."

// 	replace := strings.Replace(text, "boat", "canoe", 1)
// 	replaceAll := strings.ReplaceAll(text, "boat", "truck")

// 	fmt.Println("Replace:", replace)
// 	fmt.Println("ReplaceAll:", replaceAll)
// }

/*Изменение строк с помощью функции карты Map*/

// func main() {
// 	text := "It was a boat. A small boat."

// 	mapper := func(r rune) rune {
// 		if r == 'b' {
// 			return 'c'
// 		}
// 		return r
// 	}

// 	mapped := strings.Map(mapper, text)

// 	fmt.Println("Mapped:", mapped)
// }

/*Использование заменителя строк Replacer*/

// func main() {
// 	text := "It was a boat. A small boat."

// 	replacer := strings.NewReplacer("boat", "kayak", "small", "huge")
// 	replaced := replacer.Replace(text)

// 	fmt.Println("Replaced:", replaced)
// }

/*Построение и генерация строк Join, Repeat*/

// func main() {
// 	text := "It was a boat. A small boat."

// 	elements := strings.Fields(text)

// 	joined := strings.Join(elements, "--")

// 	fmt.Println("Joined:", joined)
// }

/* Строительные строки
WriteString(s) - добавляет строку s
WriteRune(r) - добавляет символ r
WriteByte(b) - добавляет байт b
String() - Возвращает строку созданную компоновщиком
Reset() - сбрасывает строку созданную построителем
Len() - Возвращает количество байтов, используемых для хранения
Cap() - Возвращает количество байтов, выделенных компоновщиком
Grow(size) - Увеличивает количество байтов, используемых для хранения
*/

// func main() {
// 	text := "It was a boat. A small boat."

// 	var builder strings.Builder

// 	for _, sub := range strings.Fields(text) {
// 		if sub == "small" {
// 			builder.WriteString("very ")
// 		}
// 		builder.WriteString(sub)
// 		builder.WriteRune(' ')
// 	}

// 	fmt.Println("String:", builder.String())
// }

/*
Использование регулярных выражений
Match(pattern, b) - возвращает bool - соответствует ли шаблон байтовому срезу b
MatchString(pattern, s) - возвращает bool - соответствует ли шаблон строке s
*/

// func main() {
// 	description := "A boat for one person."

// 	match, err := regexp.MatchString("[A-z]oat", description)
// 	if err == nil {
// 		fmt.Println("Match:", match)
// 	} else {
// 		fmt.Println("Error:", err)
// 	}

// }

/*
Использование регулярных выражений
Компиляция и повторное использование
Compile(pattern) - возвращает RegExp, который можно использоваать для повторного сопоставления с указанным шаблоном
MustCompile(pattern) - аналогичен предыдущему, но вызывает панику если указанный шаблон не может быть скомпилирован
*/

// func main() {
// 	pattern, compileErr := regexp.Compile("[A-z]oat")
// 	description := "A boat for one person."
// 	question := "Is that a goat?"
// 	preference := "I like oats"

// 	if compileErr == nil {
// 		fmt.Println("Description:", pattern.MatchString(description))
// 		fmt.Println("Question:", pattern.MatchString(question))
// 		fmt.Println("Preference:", pattern.MatchString(preference))
// 	} else {
// 		fmt.Println("Error:", compileErr)
// 	}

// }
/*
Description: true
Question: true
Preference: false
*/

// func getSubstring(s string, indices []int) string {
// 	return string(s[indices[0]:indices[1]])
// }

// func main() {
// 	pattern := regexp.MustCompile("K[a-z]{4}|[A-z]oat")
// 	description := "Kayak. A boat for one person."

// 	firstIndex := pattern.FindStringIndex(description)
// 	allIndices := pattern.FindAllStringIndex(description, -1)

// 	fmt.Println("first index", firstIndex[0], "-", firstIndex[1], "=", getSubstring(description, firstIndex))

// 	for i, idx := range allIndices {
// 		fmt.Println("Index", i, "=", idx[0], "-", idx[1], "=", getSubstring(description, idx))
// 	}

// }
/*
first index 0 - 5 = Kayak
Index 0 = 0 - 5 = Kayak
Index 1 = 9 - 13 = boat
*/

// func main() {
// 	pattern := regexp.MustCompile("K[a-z]{4}|[A-z]oat")
// 	description := "Kayak. A boat for one person."

// 	firstMatch := pattern.FindString(description)
// 	allMatches := pattern.FindAllString(description, -1)

// 	fmt.Println("First match", firstMatch)

// 	for i, m := range allMatches {
// 		fmt.Println("Match", i, "=", m)
// 	}

// }
/*
First match Kayak
Match 0 = Kayak
Match 1 = boat
*/

/*Разделение строк с помощью регулярного выражения*/

// func main() {
// 	pattern := regexp.MustCompile(" |boat|one")

// 	description := "Kayak. A boat for one person."

// 	split := pattern.Split(description, -1)

// 	for _, s := range split {
// 		if s != "" {
// 			fmt.Println("Substring:", s)
// 		}
// 	}
// }

/*Использование под выражений*/

// func main() {
// 	pattern := regexp.MustCompile("A [A-z]* for [A-z]* person")

// 	description := "Kayak. A boat for one person."

// 	str := pattern.FindString(description)

// 	fmt.Println("Match:", str)
// }
/*Match: A boat for one person*/

// func main() {
// 	pattern := regexp.MustCompile("A ([A-z]*) for ([A-z]*) person")

// 	description := "Kayak. A boat for one person."

// 	sub := pattern.FindStringSubmatch(description)

// 	for _, s := range sub {
// 		fmt.Println("Match:", s)
// 	}
// }
/*
Match: A boat for one person
Match: boat
Match: one
*/

/*Использование именнованых подвыражений*/

// func main() {
// 	pattern := regexp.MustCompile("A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")

// 	description := "Kayak. A boat for one person."

// 	subs := pattern.FindStringSubmatch(description)

// 	for _, name := range []string{"type", "capacity"} {
// 		fmt.Println(name, "=", subs[pattern.SubexpIndex(name)])
// 	}
// }
/*
type = boat
capacity = one
*/

/*Подмена строк с помощью регулярного выражения*/

// func main() {
// 	pattern := regexp.MustCompile("A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")

// 	description := "Kayak. A boat for one person."

// 	template := "(type: ${type}, capacity: ${capacity})"

// 	replaced := pattern.ReplaceAllString(description, template)

// 	fmt.Println("ReplaceAllString: ", replaced)
// }

/*Kayak. (type: boat, capacity: one).*/

/*Замена совпадающего контента функцией*/

// func main() {
// 	pattern := regexp.MustCompile("A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")

// 	description := "Kayak. A boat for one person."

// 	replaced := pattern.ReplaceAllStringFunc(description, func(s string) string {
// 		return "This is the replacement content"
// 	})

// 	fmt.Println(replaced)
// }
/*Kayak. This is the replacement content.*/
