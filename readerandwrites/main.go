package main

/*кодирование данных JSON*/

// func main() {
// 	var b bool = true
// 	var str string = "Hello"
// 	var fval float64 = 99.99
// 	var ival int = 200
// 	var pointer *int = &ival

// 	var writer strings.Builder
// 	encoder := json.NewEncoder(&writer)

// 	for _, val := range []interface{}{b, str, fval, ival, pointer} {
// 		encoder.Encode(val)
// 	}
// 	fmt.Print(writer.String())
// }

/*
true
"Hello"
99.99
200
200
*/

/*Кодирование массивов и срезов*/

// func main() {
// 	names := []string{"Kayak", "Lifejaket", "Soccer Ball"}
// 	numbers := [3]int{12, 20, 30}

// 	var byteArray [5]byte
// 	copy(byteArray[0:], []byte(names[0]))
// 	byteSlice := []byte(names[0])

// 	var writer strings.Builder
// 	encoder := json.NewEncoder(&writer)

// 	encoder.Encode(names)
// 	encoder.Encode(numbers)
// 	encoder.Encode(byteArray)
// 	encoder.Encode(byteSlice)

// 	fmt.Print(writer.String())
// }

/*
["Kayak","Lifejaket","Soccer Ball"]
[12,20,30]
байтовые массивы и срезы отличаются
[75,97,121,97,107]
"S2F5YWs="
*/

/*Кодирование карт*/

// func main() {
// 	m := map[string]float64{
// 		"Kayak":      279,
// 		"Lifejacket": 49.95,
// 	}

// 	var writer strings.Builder
// 	encoder := json.NewEncoder(&writer)

// 	encoder.Encode(m)

// 	fmt.Print(writer.String())
// }

/*{"Kayak":279,"Lifejacket":49.95}*/

/*Кодирование структур*/

// func main() {
// 	var writer strings.Builder
// 	encoder := json.NewEncoder(&writer)

// 	encoder.Encode(Kayak)

// 	fmt.Print(writer.String())
// }

/*{"Name":"Kayak","Category":"Watersports","Price":279}*/

/*Понимание еффекта продвижения в JSON при кодировании*/

// func main() {
// 	var writer strings.Builder
// 	encoder := json.NewEncoder(&writer)

// 	dp := DiscountedProduct{
// 		Product:  &Kayak,
// 		Discount: 10.50,
// 	}

// 	encoder.Encode(dp)

// 	fmt.Print(writer.String())
// }
/*{"Name":"Kayak","Category":"Watersports","Price":279,"Discount":10.5}*/

/*
настройка JSON-кодирования структур

в discoun.go внесли изменения, добавлено `json:"product"`

вывод изменится
{"product":{"Name":"Kayak","Category":"Watersports","Price":279},"Discount":10.5}
использование тэга предотвратло продвижение поля
*/

/*
пропуск поля

в discoun.go внесли изменения, добавлено `json:"-"`

в выводе поле Discount к которому добавлен тэг пропущено
{"product":{"Name":"Kayak","Category":"Watersports","Price":279}}
*/

/*пропуск не назначенных полей*/

// func main() {
// 	var writer strings.Builder
// 	encoder := json.NewEncoder(&writer)

// 	dp := DiscountedProduct{
// 		Product:  &Kayak,
// 		Discount: 10.50,
// 	}

// 	encoder.Encode(dp)

// 	db2 := DiscountedProduct{Discount: 10.50}
// 	encoder.Encode(&db2)

// 	fmt.Print(writer.String())
// }

/*
{"product":{"Name":"Kayak","Category":"Watersports","Price":279}}
{"product":null}

применение omitempty
{"product":{"Name":"Kayak","Category":"Watersports","Price":279}}
{}

использование omitempty без имени (`json:",omitempty"`)
{"Name":"Kayak","Category":"Watersports","Price":279}
{}

*/

/*
Принудительное кодирование полей как строк

тэг- `json:",string"`

до
{"Name":"Kayak","Category":"Watersports","Price":279,"Discount":10.5}
{"Discount":10.5}

после
{"Name":"Kayak","Category":"Watersports","Price":279,"Discount":"10.5"}
{"Discount":"10.5"}
*/

/*Интерфейсы кодирования добавлен файл interface.go*/
// func main() {
// 	var writer strings.Builder
// 	encoder := json.NewEncoder(&writer)

// 	dp := DiscountedProduct{
// 		Product:  &Kayak,
// 		Discount: 10.50,
// 	}

// 	namedItems := []Named{&dp, &Person{PersonName: "Alice"}}
// 	encoder.Encode(namedItems)

// 	fmt.Print(writer.String())
// }
/*[{"Name":"Kayak","Category":"Watersports","Price":279,"Discount":"10.5"},{"PersonName":"Alice"}]*/

/*Создание полностью настраиваемых кодировок JSON*/

// func main() {
// 	var writer strings.Builder
// 	encoder := json.NewEncoder(&writer)

// 	dp := DiscountedProduct{
// 		Product:  &Kayak,
// 		Discount: 10.50,
// 	}

// 	namedItems := []Named{&dp, &Person{PersonName: "Alice"}}
// 	encoder.Encode(namedItems)

// 	fmt.Print(writer.String())
// }

/*
В файл discount.go добавлена функция - func (dp *DiscountedProduct) MarshalJSON() (jsn []byte, err error)
[{"cost":268.5,"product":"Kayak"},{"PersonName":"Alice"}]
*/

/*Декодирование данных JSON*/

// func main() {
// 	reader := strings.NewReader(`true "Hello" 99.99 200`)
// 	vals := []interface{}{}

// 	decoder := json.NewDecoder(reader)

// 	for {
// 		var decoderVal interface{}
// 		err := decoder.Decode(&decoderVal)
// 		if err != nil {
// 			if err != io.EOF {
// 				Printfln("Error: %v", err.Error())
// 			}
// 			break
// 		}
// 		vals = append(vals, decoderVal)
// 	}

// 	for _, val := range vals {
// 		Printfln("Decoded (%T): %v", val, val)
// 	}
// }

/*
Decoded (bool): true
Decoded (string): Hello
Decoded (float64): 99.99
Decoded (float64): 200
*/

/*Расшифровка числовых значений*/

// func main() {
// 	reader := strings.NewReader(`true "Hello" 99.99 200`)
// 	vals := []interface{}{}

// 	decoder := json.NewDecoder(reader)
// 	decoder.UseNumber()

// 	for {
// 		var decoderVal interface{}
// 		err := decoder.Decode(&decoderVal)
// 		if err != nil {
// 			if err != io.EOF {
// 				Printfln("Error: %v", err.Error())
// 			}
// 			break
// 		}
// 		vals = append(vals, decoderVal)
// 	}

// 	for _, val := range vals {
// 		if num, ok := val.(json.Number); ok {
// 			if ival, err := num.Int64(); err == nil {
// 				Printfln("Decoded Integer: %v", ival)
// 			} else if fpval, err := num.Float64(); err == nil {
// 				Printfln("Decoded Floating Point: %v", fpval)
// 			} else {
// 				Printfln("Decoded String: %v", num.String())
// 			}
// 		} else {
// 			Printfln("Decoded (%T): %v", val, val)
// 		}
// 	}
// }
/*
Decoded (bool): true
Decoded (string): Hello
Decoded Floating Point: 99.99
Decoded Integer: 200
*/

/*Указание типов для декодирования*/
