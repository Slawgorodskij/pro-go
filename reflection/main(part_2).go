package main

/*
func createPointerType(t reflect.Type) reflect.Type {
	return reflect.PtrTo(t)
}

func followPointerType(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}

func main() {
	name := "Alice"
	t := reflect.TypeOf(name)
	Printfln("Original Type: %v", t)
	pt := createPointerType(t)
	Printfln("Pointer Type: %v", pt)
	Printfln("Follow pointer type: %v", followPointerType(pt))
}
*/

// Работа со значениями указателя
/*
var stringPtrType = reflect.TypeOf((*string)(nil))

func transformString(val interface{}) {
	elemValue := reflect.ValueOf(val)
	if elemValue.Type() == stringPtrType {
		upperStr := strings.ToUpper(elemValue.Elem().String())
		if elemValue.Elem().CanSet() {
			elemValue.Elem().SetString(upperStr)
		}
	}
}
func main() {
	name := "Alice"
	transformString(&name)
	Printfln("Follow pointer value: %v", name)
}
*/

//Работа с типами массивов и срезов
/*
func checkElemType(val interface{}, arrOrSlice interface{}) bool {
	elemType := reflect.TypeOf(val)
	arrOrSliceType := reflect.TypeOf(arrOrSlice)
	return (arrOrSliceType.Kind() == reflect.Array ||
		arrOrSliceType.Kind() == reflect.Slice) &&
		arrOrSliceType.Elem() == elemType
}
func main() {
	name := "Alice"
	city := "London"
	hobby := "Running"
	slice := []string{name, city, hobby}
	array := [3]string{name, city, hobby}
	Printfln("Slice (string): %v", checkElemType("testString",
		slice))
	Printfln("Array (string): %v", checkElemType("testString",
		array))
	Printfln("Array (int): %v", checkElemType(10, array))
}
*/
//Работа со значениями массива и среза
/*
func setValue(arrayOrSlice interface{}, index int, replacement interface{}) {
	arrayOrSliceVal := reflect.ValueOf(arrayOrSlice)
	replacementVal := reflect.ValueOf(replacement)
	if arrayOrSliceVal.Kind() == reflect.Slice {
		elemVal := arrayOrSliceVal.Index(index)
		if elemVal.CanSet() {
			elemVal.Set(replacementVal)
		}
	} else if arrayOrSliceVal.Kind() == reflect.Ptr &&
		arrayOrSliceVal.Elem().Kind() == reflect.Array &&
		arrayOrSliceVal.Elem().CanSet() {
		arrayOrSliceVal.Elem().Index(index).Set(replacementVal)
	}
}
func main() {
	name := "Alice"
	city := "London"
	hobby := "Running"
	slice := []string{name, city, hobby}
	array := [3]string{name, city, hobby}
	Printfln("Original slice: %v", slice)
	newCity := "Paris"
	setValue(slice, 1, newCity)
	Printfln("Modified slice: %v", slice)
	Printfln("Original slice: %v", array)
	newCity = "Rome"
	setValue(&array, 1, newCity)
	Printfln("Modified slice: %v", array)
}
*/
//Перечисление срезов и массивов
/*
func enumerateStrings(arrayOrSlice interface{}) {
	arrayOrSliceVal := reflect.ValueOf(arrayOrSlice)
	if (arrayOrSliceVal.Kind() == reflect.Array ||
		arrayOrSliceVal.Kind() == reflect.Slice) &&
		arrayOrSliceVal.Type().Elem().Kind() == reflect.String {
		for i := 0; i < arrayOrSliceVal.Len(); i++ {
			Printfln("Element: %v, 	Value: %v", i, arrayOrSliceVal.Index(i).String())
		}
	}
}

func main() {
	name := "Alice"
	city := "London"
	hobby := "Running"
	slice := []string{name, city, hobby}
	array := [3]string{name, city, hobby}
	enumerateStrings(slice)
	enumerateStrings(array)
	enumerateStrings(slice)
	enumerateStrings(array)
}
*/
//Создание новых срезов из существующих срезов
/*
func findAndSplit(slice interface{}, target interface{}) interface{} {
	sliceVal := reflect.ValueOf(slice)
	targetType := reflect.TypeOf(target)
	if sliceVal.Kind() == reflect.Slice && sliceVal.Type().Elem() == targetType {
		for i := 0; i < sliceVal.Len(); i++ {
			if sliceVal.Index(i).Interface() == target {
				return sliceVal.Slice(0, i+1)
			}
		}
	}
	return slice
}
func main() {
	name := "Alice"
	city := "London"
	hobby := "Running"
	slice := []string{name, city, hobby}
	Printfln("Strings: %v", findAndSplit(slice, "London"))
	numbers := []int{1, 3, 4, 5, 7}
	Printfln("Numbers: %v", findAndSplit(numbers, 4))
}
*/
//Создание, копирование и добавление элементов в срезы
/*
func pickValues(slice interface{}, indices ...int) interface{} {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() == reflect.Slice {
		newSlice := reflect.MakeSlice(sliceVal.Type(), 0, 10)
		for _, index := range indices {
			newSlice = reflect.Append(newSlice,
				sliceVal.Index(index))
		}
		return newSlice
	}
	return nil
}
func main() {
	name := "Alice"
	city := "London"
	hobby := "Running"
	slice := []string{name, city, hobby, "Bob", "Paris",
		"Soccer"}
	picked := pickValues(slice, 0, 3, 5)
	Printfln("Picked values: %v", picked)
}
*/

// Работа с типами карт
/*
func describeMap(m interface{}) {
	mapType := reflect.TypeOf(m)
	if mapType.Kind() == reflect.Map {
		Printfln("Key type: %v, Val type: %v", mapType.Key(),
			mapType.Elem())
	} else {
		Printfln("Not a map")
	}
}
func main() {
	pricesMap := map[string]float64{"Kayak": 279, "Lifejacket": 48.95, "Soccer Ball": 19.50}
	describeMap(pricesMap)
}
*/
//Работа со значениями карты
/*
func printMapContents(m interface{}) {
	mapValue := reflect.ValueOf(m)
	if mapValue.Kind() == reflect.Map {
		for _, keyVal := range mapValue.MapKeys() {
			reflectedVal := mapValue.MapIndex(keyVal)
			Printfln("Map Key: %v, Value: %v", keyVal,
				reflectedVal)
		}
	} else {
		Printfln("Not a map")
	}
}
func main() {
	pricesMap := map[string]float64{
		"Kayak": 279, "Lifejacket": 48.95, "Soccer Ball": 19.50,
	}
	printMapContents(pricesMap)
}
*/
//Установка и удаление значений карты
/*
func setMap(m interface{}, key interface{}, val interface{}) {
	mapValue := reflect.ValueOf(m)
	keyValue := reflect.ValueOf(key)
	valValue := reflect.ValueOf(val)
	if mapValue.Kind() == reflect.Map &&
		mapValue.Type().Key() == keyValue.Type() &&
		mapValue.Type().Elem() == valValue.Type() {
		mapValue.SetMapIndex(keyValue, valValue)
	} else {
		Printfln("Not a map or mismatched types")
	}
}
func removeFromMap(m interface{}, key interface{}) {
	mapValue := reflect.ValueOf(m)
	keyValue := reflect.ValueOf(key)
	if mapValue.Kind() == reflect.Map &&
		mapValue.Type().Key() == keyValue.Type() {
		mapValue.SetMapIndex(keyValue, reflect.Value{})
	}
}
func main() {
	pricesMap := map[string]float64{
		"Kayak":       279,
		"Lifejacket":  48.95,
		"Soccer Ball": 19.50,
	}
	setMap(pricesMap, "Kayak", 100.00)
	setMap(pricesMap, "Hat", 10.00)
	removeFromMap(pricesMap, "Lifejacket")
	for k, v := range pricesMap {
		Printfln("Key: %v, Value: %v", k, v)
	}
}
*/

//Создание новых карт
/*
func createMap(slice interface{}, op func(interface{}) interface{}) interface{} {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() == reflect.Slice {
		mapType := reflect.MapOf(sliceVal.Type().Elem(),
			sliceVal.Type().Elem())
		mapVal := reflect.MakeMap(mapType)
		for i := 0; i < sliceVal.Len(); i++ {
			elemVal := sliceVal.Index(i)
			mapVal.SetMapIndex(elemVal,
				reflect.ValueOf(op(elemVal.Interface())))
		}
		return mapVal.Interface()
	}
	return nil
}
func main() {
	names := []string{"Alice", "Bob", "Charlie"}
	reverse := func(val interface{}) interface{} {
		if str, ok := val.(string); ok {
			return strings.ToUpper(str)
		}
		return val
	}
	namesMap := createMap(names, reverse).(map[string]string)
	for k, v := range namesMap {
		Printfln("Key: %v, Value:%v", k, v)
	}
}
*/
//Работа с типами структур
/*
func inspectStructs(structs ...interface{}) {
	for _, s := range structs {
		structType := reflect.TypeOf(s)
		if structType.Kind() == reflect.Struct {
			inspectStructType(structType)
		}
	}
}
func inspectStructType(structType reflect.Type) {
	Printfln("--- Struct Type: %v", structType)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		Printfln("Field %v: Name: %v, Type: %v, Exported: %v", field.Index, field.Name, field.Type, field.PkgPath == "")
	}
	Printfln("--- End Struct Type: %v", structType)
}
func main() {
	inspectStructs(Purchase{})
}
*/

// Обработка вложенных полей
/*
func inspectStructs(structs ...interface{}) {
	for _, s := range structs {
		structType := reflect.TypeOf(s)
		if structType.Kind() == reflect.Struct {
			inspectStructType([]int{}, structType)
		}
	}
}
func inspectStructType(baseIndex []int, structType reflect.Type) {
	Printfln("--- Struct Type: %v", structType)
	for i := 0; i < structType.NumField(); i++ {
		fieldIndex := append(baseIndex, i)
		field := structType.Field(i)
		Printfln("Field %v: Name: %v, Type: %v, Exported: %v",
			fieldIndex, field.Name, field.Type, field.PkgPath ==
				"")
		if field.Type.Kind() == reflect.Struct {
			field := structType.FieldByIndex(fieldIndex)
			inspectStructType(fieldIndex, field.Type)
		}
	}
	Printfln("--- End Struct Type: %v", structType)
}
func main() {
	inspectStructs(Purchase{})
}
*/
//Поиск поля по имени
/*
func describeField(s interface{}, fieldName string) {
	structType := reflect.TypeOf(s)
	field, found := structType.FieldByName(fieldName)
	if found {
		Printfln("Found: %v, Type: %v, Index: %v",
			field.Name, field.Type, field.Index)
		index := field.Index
		for len(index) > 1 {
			index = index[0 : len(index)-1]
			field = structType.FieldByIndex(index)
			Printfln("Parent : %v, Type: %v, Index: %v",
				field.Name, field.Type, field.Index)
		}
		Printfln("Top-Level Type: %v", structType)
	} else {
		Printfln("Field %v not found", fieldName)
	}
}
func main() {
	describeField(Purchase{}, "Price")
}
*/

// Проверка тегов структуры
/*
func inspectTags(s interface{}, tagName string) {
	structType := reflect.TypeOf(s)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		tag := field.Tag
		valGet := tag.Get(tagName)
		valLookup, ok := tag.Lookup(tagName)
		Printfln("Field: %v, Tag %v: %v", field.Name, tagName,
			valGet)
		Printfln("Field: %v, Tag %v: %v, Set: %v",
			field.Name, tagName, valLookup, ok)
	}
}

type Person struct {
	Name    string `alias:"id"`
	City    string `alias:""`
	Country string
}

func main() {
	inspectTags(Person{}, "alias")
}
*/

// Создание типов структур
/*
func inspectTags(s interface{}, tagName string) {
	structType := reflect.TypeOf(s)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		tag := field.Tag
		valGet := tag.Get(tagName)
		valLookup, ok := tag.Lookup(tagName)
		Printfln("Field: %v, Tag %v: %v", field.Name, tagName, valGet)
		Printfln("Field: %v, Tag %v: %v, Set: %v", field.Name, tagName, valLookup, ok)
	}
}
func main() {
	stringType := reflect.TypeOf("this is a string")
	structType := reflect.StructOf([]reflect.StructField{
		{Name: "Name", Type: stringType, Tag: `alias:"id"`},
		{Name: "City", Type: stringType, Tag: `alias:""`},
		{Name: "Country", Type: stringType},
	})
	inspectTags(reflect.New(structType), "alias")
}
*/
//Работа со структурными значениями
/*
func getFieldValues(s interface{}) {
	structValue := reflect.ValueOf(s)
	if structValue.Kind() == reflect.Struct {
		for i := 0; i < structValue.NumField(); i++ {
			fieldType := structValue.Type().Field(i)
			fieldVal := structValue.Field(i)
			Printfln("Name: %v, Type: %v, Value: %v", fieldType.Name, fieldType.Type, fieldVal)
		}
	} else {
		Printfln("Not a struct")
	}
}
func main() {
	product := Product{Name: "Kayak", Category: "Watersports", Price: 279}
	customer := Customer{Name: "Acme", City: "Chicago"}
	purchase := Purchase{Customer: customer, Product: product, Total: 279, taxRate: 10}
	getFieldValues(purchase)
}
*/
//Установка значений поля структуры
/*
func setFieldValue(s interface{}, newVals map[string]interface{}) {
	structValue := reflect.ValueOf(s)
	if structValue.Kind() == reflect.Ptr &&
		structValue.Elem().Kind() == reflect.Struct {
		for name, newValue := range newVals {
			fieldVal := structValue.Elem().FieldByName(name)
			if fieldVal.CanSet() {
				fieldVal.Set(reflect.ValueOf(newValue))
			} else if fieldVal.CanAddr() {
				ptr := fieldVal.Addr()
				if ptr.CanSet() {
					ptr.Set(reflect.ValueOf(newValue))
				} else {
					Printfln("Cannot set field via pointer")
				}
			} else {
				Printfln("Cannot set field")
			}
		}
	} else {
		Printfln("Not a pointer to a struct")
	}
}
func getFieldValues(s interface{}) {
	structValue := reflect.ValueOf(s)
	if structValue.Kind() == reflect.Struct {
		for i := 0; i < structValue.NumField(); i++ {
			fieldType := structValue.Type().Field(i)
			fieldVal := structValue.Field(i)
			Printfln("Name: %v, Type: %v, Value: %v",
				fieldType.Name, fieldType.Type, fieldVal)
		}
	} else {
		Printfln("Not a struct")
	}
}
func main() {
	product := Product{Name: "Kayak", Category: "Watersports",
		Price: 279}
	customer := Customer{Name: "Acme", City: "Chicago"}
	purchase := Purchase{Customer: customer, Product: product,
		Total:   279,
		taxRate: 10}
	setFieldValue(&purchase, map[string]interface{}{
		"City": "London", "Category": "Boats", "Total": 100.50,
	})
	getFieldValues(purchase)
}
*/
