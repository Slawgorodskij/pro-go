package main

//func printDetails(values ...Product) {
//	for _, elem := range values {
//		Printfln("Product: Name: %v, Category: %v, Price: %v", elem.Name, elem.Category, elem.Price)
//	}
//}

//для решения проблемы получения ошибки
/*
func printDetails(values ...interface{}) {
	for _, elem := range values {
		switch val := elem.(type) {
		case Product:
			Printfln("Product: Name: %v, Category: %v, Price: %v", val.Name, val.Category, val.Price)
		case Customer:
			Printfln("Customer: Name: %v, City: %v", val.Name, val.City)
		}
	}
}

func main() {
	product := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}

	//printDetails(product)
	// Product: Name: Kayak, Category: Watersports, Price: 279

	customer := Customer{Name: "Alice", City: "New York"}
	printDetails(product, customer)
	//  cannot use customer (variable of type Customer) as Product value in argument to printDetails

	//после изменения функции printDetails
	// Product: Name: Kayak, Category: Watersports, Price: 279
	// Customer: Name: Alice, City: New York
}
*/
// У выше описанного проблема все равно останется так как функции printDetails придется постоянно расширять для исключения
//в Go предусмотренна "Рефлексия" - которая позволяет приложению работать с типами, неизвестными при компиляции проекта

// Использование рефлексии
/*
func printDetails(values ...interface{}) {
	for _, elem := range values {
		fieldDetails := []string{}
		elemType := reflect.TypeOf(elem)
		elemValue := reflect.ValueOf(elem)
		if elemType.Kind() == reflect.Struct {
			for i := 0; i < elemType.NumField(); i++ {
				fieldName := elemType.Field(i).Name
				fieldVal := elemValue.Field(i)
				fieldDetails = append(fieldDetails, fmt.Sprintf("%v: %v", fieldName, fieldVal))
			}
			Printfln("%v: %v", elemType.Name(), strings.Join(fieldDetails, ", "))
		} else {
			Printfln("%v: %v", elemType.Name(), elemValue)
		}
	}
}

type Payment struct {
	Currency string
	Amount   float64
}

func main() {
	product := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	customer := Customer{Name: "Alice", City: "New York"}
	payment := Payment{Currency: "USD", Amount: 100.50}
	printDetails(product, customer, payment, 10, true)
}
*/
// Использование основных функций типа
/*
func getTypePath(t reflect.Type) (path string) {
	path = t.PkgPath()
	if path == "" {
		path = "(built-in)"
	}
	return
}

func printDetails(values ...interface{}) {
	for _, elem := range values {
		elemType := reflect.TypeOf(elem)

		Printfln("Name: %v, PkgPath: %v, Kind: %v", elemType.Name(), getTypePath(elemType), elemType.Kind())
	}
}

type Payment struct {
	Currency string
	Amount   float64
}

func main() {
	product := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	customer := Customer{Name: "Alice", City: "New York"}
	payment := Payment{Currency: "USD", Amount: 100.50}
	printDetails(product, customer, payment, 10, true)
}
*/

// Использование базовых возможностей Value
/*
func printDetails(values ...interface{}) {
	for _, elem := range values {
		elemValue := reflect.ValueOf(elem)
		switch elemValue.Kind() {
		case reflect.Bool:
			var val bool = elemValue.Bool()
			Printfln("Bool: %v", val)
		case reflect.Int:
			var val int64 = elemValue.Int()
			Printfln("Int: %v", val)
		case reflect.Float32, reflect.Float64:
			var val float64 = elemValue.Float()
			Printfln("Float: %v", val)
		case reflect.String:
			var val string = elemValue.String()
			Printfln("String: %v", val)
		case reflect.Ptr:
			var val reflect.Value = elemValue.Elem()
			if val.Kind() == reflect.Int {
				Printfln("Pointer to Int: %v", val.Int())
			}
		default:
			Printfln("Other: %v", elemValue.String())
		}
	}
}

func main() {
	product := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	number := 100
	printDetails(true, 10, 23.30, "Alice", &number, product)
}
*/
//Определение типов
/*
var intPtrType = reflect.TypeOf((*int)(nil))

func printDetails(values ...interface{}) {
	for _, elem := range values {
		elemValue := reflect.ValueOf(elem)
		elemType := reflect.TypeOf(elem)
		if elemType == intPtrType {
			Printfln("Pointer to Int: %v",
				elemValue.Elem().Int())
		} else {
			switch elemValue.Kind() {
			case reflect.Bool:
				var val bool = elemValue.Bool()
				Printfln("Bool: %v", val)
			case reflect.Int:
				var val int64 = elemValue.Int()
				Printfln("Int: %v", val)
			case reflect.Float32, reflect.Float64:
				var val float64 = elemValue.Float()
				Printfln("Float: %v", val)
			case reflect.String:
				var val string = elemValue.String()
				Printfln("String: %v", val)
			default:
				Printfln("Other: %v", elemValue.String())
			}
		}
	}
}

func main() {
	product := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	number := 100
	printDetails(true, 10, 23.30, "Alice", &number, product)
}
*/

//Идентификация байтовых срезов
/*
var intPtrType = reflect.TypeOf((*int)(nil))
var byteSliceType = reflect.TypeOf([]byte(nil))

func printDetails(values ...interface{}) {
	for _, elem := range values {
		elemValue := reflect.ValueOf(elem)
		elemType := reflect.TypeOf(elem)
		if elemType == intPtrType {
			Printfln("Pointer to Int: %v",
				elemValue.Elem().Int())
		} else if elemType == byteSliceType {
			Printfln("Byte slice: %v", elemValue.Bytes())
		} else {
			switch elemValue.Kind() {
			case reflect.Bool:
				var val bool = elemValue.Bool()
				Printfln("Bool: %v", val)
			case reflect.Int:
				var val int64 = elemValue.Int()
				Printfln("Int: %v", val)
			case reflect.Float32, reflect.Float64:
				var val float64 = elemValue.Float()
				Printfln("Float: %v", val)
			case reflect.String:
				var val string = elemValue.String()
				Printfln("String: %v", val)
			default:
				Printfln("Other: %v", elemValue.String())
			}
		}
	}
}

func main() {
	product := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	number := 100
	slice := []byte("Alice")
	printDetails(true, 10, 23.30, "Alice", &number, product, slice)
}
*/
// Получение базовых значений
/*
func selectValue(data interface{}, index int) (result interface{}) {
	dataVal := reflect.ValueOf(data)
	if dataVal.Kind() == reflect.Slice {
		result = dataVal.Index(index).Interface()
	}
	return
}

func main() {
	names := []string{"Alice", "Bob", "Charlie"}
	val := selectValue(names, 1).(string)
	Printfln("Selected: %v", val)
}
*/

// Установка Value с использованием рефлексии
/*
func incrementOrUpper(values ...interface{}) {
	for _, elem := range values {
		elemValue := reflect.ValueOf(elem)
		if elemValue.Kind() == reflect.Ptr {
			elemValue = elemValue.Elem()
		}
		if elemValue.CanSet() {
			switch elemValue.Kind() {
			case reflect.Int:
				elemValue.SetInt(elemValue.Int() + 1)
			case reflect.String:
				elemValue.SetString(strings.ToUpper(elemValue.String()))
			}
			Printfln("Modified Value: %v", elemValue)
		} else {
			Printfln("Cannot set %v: %v", elemValue.Kind(), elemValue)
		}
	}
}
func main() {
	name := "Alice"
	price := 279
	city := "London"
	incrementOrUpper(&name, &price, &city)
	for _, val := range []interface{}{name, price, city} {
		Printfln("Value: %v", val)
	}
}
*/

// Установка одного Value с помощью другого
/*
func setAll(src interface{}, targets ...interface{}) {
	srcVal := reflect.ValueOf(src)
	for _, target := range targets {
		targetVal := reflect.ValueOf(target)
		if targetVal.Kind() == reflect.Ptr && targetVal.Elem().Type() == srcVal.Type() && targetVal.Elem().CanSet() {
			targetVal.Elem().Set(srcVal)
		}
	}
}
func main() {
	name := "Alice"
	price := 279
	city := "London"
	setAll("New String", &name, &price, &city)
	setAll(10, &name, &price, &city)
	for _, val := range []interface{}{name, price, city} {
		Printfln("Value: %v", val)
	}
}
*/

// Сравнение Value
/*
При использовании этого варианта выдается ошибка:
Found #1: true
panic: runtime error: comparing uncomparable type []string
goroutine 1 [running]:
main.contains({0x4894c0?, 0xc000010030?}, {0x489a00, 0xc00010cf08})
        /home/vladimir/project/pro-go/reflection/main(part_1).go:310 +0x12f
main.main()
        /home/vladimir/project/pro-go/reflection/main(part_1).go:322 +0x23c
exit status 2

func contains(slice interface{}, target interface{}) (found bool) {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() == reflect.Slice {
		for i := 0; i < sliceVal.Len(); i++ {
			if sliceVal.Index(i).Interface() == target {
				found = true
			}
		}
	}
	return
}
*/
/*
//После внесенных изменений получаем ожидаемый результат
func contains(slice interface{}, target interface{}) (found bool) {
	sliceVal := reflect.ValueOf(slice)
	targetType := reflect.TypeOf(target)
	if sliceVal.Kind() == reflect.Slice && sliceVal.Type().Elem().Comparable() && targetType.Comparable() {
		for i := 0; i < sliceVal.Len(); i++ {
			if sliceVal.Index(i).Interface() == target {
				found = true
			}
		}
	}
	return
}

func main() {
	city := "London"
	citiesSlice := []string{"Paris", "Rome", "London"}
	Printfln("Found #1: %v", contains(citiesSlice, city))
	sliceOfSlices := [][]string{citiesSlice, {"First", "Second", "Third"}}
	Printfln("Found #2: %v", contains(sliceOfSlices, citiesSlice))
}
*/

//Использование удобной функции сравнения https://pkg.go.dev/reflect@go1.17.1#DeepEqual,
/*
func contains(slice interface{}, target interface{}) (found bool) {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() == reflect.Slice {
		for i := 0; i < sliceVal.Len(); i++ {
			if reflect.DeepEqual(sliceVal.Index(i).Interface(), target) {
				found = true
			}
		}
	}
	return
}
func main() {
	city := "London"
	citiesSlice := []string{"Paris", "Rome", "London"}
	Printfln("Found #1: %v", contains(citiesSlice, city))
	sliceOfSlices := [][]string{citiesSlice, {"First", "Second", "Third"}}
	Printfln("Found	#2: %v", contains(sliceOfSlices, citiesSlice))
}
*/

// Преобразование значений
/*
func convert(src, target interface{}) (result interface{}, assigned bool) {
	srcVal := reflect.ValueOf(src)
	targetVal := reflect.ValueOf(target)
	if srcVal.Type().ConvertibleTo(targetVal.Type()) {
		result = srcVal.Convert(targetVal.Type()).Interface()
		assigned = true
	} else {
		result = src
	}
	return
}
func main() {
	name := "Alice"
	price := 279
	newVal, ok := convert(price, 100.00)
	Printfln("Converted %v: %v, %T", ok, newVal, newVal)
	newVal, ok = convert(name, 100.00)
	Printfln("Converted %v: %v, %T", ok, newVal, newVal)
}
*/

// Преобразование числовых типов
/*
func IsInt(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}
	return false
}

func IsFloat(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

func convert(src, target interface{}) (result interface{}, assigned bool) {
	srcVal := reflect.ValueOf(src)
	targetVal := reflect.ValueOf(target)
	if srcVal.Type().ConvertibleTo(targetVal.Type()) {
		if (IsInt(targetVal) && IsInt(srcVal)) && targetVal.OverflowInt(srcVal.Int()) {
			Printfln("Int overflow")
			return src, false
		} else if IsFloat(targetVal) && IsFloat(srcVal) && targetVal.OverflowFloat(srcVal.Float()) {
			Printfln("Float overflow")
			return src, false
		}
		result = srcVal.Convert(targetVal.Type()).Interface()
		assigned = true
	} else {
		result = src
	}
	return
}

func main() {
	name := "Alice"
	price := 279

	newVal, ok := convert(price, 100.00)
	Printfln("Converted %v: %v, %T", ok, newVal, newVal)
	newVal, ok = convert(name, 100.00)
	Printfln("Converted %v: %v, %T", ok, newVal, newVal)
	newVal, ok = convert(5000, int8(100))
	Printfln("Converted %v: %v, %T", ok, newVal, newVal)
}
*/

// Создание новых значений
/*
func swap(first interface{}, second interface{}) {
	firstValue, secondValue := reflect.ValueOf(first),
		reflect.ValueOf(second)
	if firstValue.Type() == secondValue.Type() && firstValue.Kind() == reflect.Ptr && firstValue.Elem().CanSet() && secondValue.Elem().CanSet() {
		temp := reflect.New(firstValue.Elem().Type())
		temp.Elem().Set(firstValue.Elem())
		firstValue.Elem().Set(secondValue.Elem())
		secondValue.Elem().Set(temp.Elem())
	}
}
func main() {
	name := "Alice"
	price := 279
	city := "London"
	swap(&name, &city)
	for _, val := range []interface{}{name, price, city} {
		Printfln("Value: %v", val)
	}
}
*/
