package main

import "reflect"

// Работа с типами функций
/*
func inspectFuncType(f interface{}) {
	funcType := reflect.TypeOf(f)
	if funcType.Kind() == reflect.Func {
		Printfln("Function parameters: %v", funcType.NumIn())
		for i := 0; i < funcType.NumIn(); i++ {
			paramType := funcType.In(i)
			if i < funcType.NumIn()-1 {
				Printfln("Parameter #%v, Type: %v", i, paramType)
			} else {
				Printfln("Parameter #%v, Type: %v, Variadic: %v", i, paramType, funcType.IsVariadic())
			}
		}
		Printfln("Function results: %v", funcType.NumOut())
		for i := 0; i < funcType.NumOut(); i++ {
			resultType := funcType.Out(i)
			Printfln("Result #%v, Type: %v", i, resultType)
		}
	}
}
func main() {
	inspectFuncType(Find)
}
*/
//Работа со значениями функций
//var 1 answer - Result #0: true
/*
func invokeFunction(f interface{}, params ...interface{}) {
	paramVals := []reflect.Value{}
	for _, p := range params {
		paramVals = append(paramVals, reflect.ValueOf(p))
	}
	funcVal := reflect.ValueOf(f)
	if funcVal.Kind() == reflect.Func {
		results := funcVal.Call(paramVals)
		for i, r := range results {
			Printfln("Result #%v: %v", i, r)
		}
	}
}
func main() {
	names := []string{"Alice", "Bob", "Charlie"}
	invokeFunction(Find, names, "London", "Bob")
}
*/
// var 2 answer - Results: [ALICE BOB CHARLIE]
/*
func mapSlice(slice interface{}, mapper interface{}) (mapped []interface{}) {
	sliceVal := reflect.ValueOf(slice)
	mapperVal := reflect.ValueOf(mapper)
	mapped = []interface{}{}
	if sliceVal.Kind() == reflect.Slice && mapperVal.Kind() == reflect.Func && mapperVal.Type().NumIn() == 1 && mapperVal.Type().In(0) == sliceVal.Type().Elem() {
		for i := 0; i < sliceVal.Len(); i++ {
			result := mapperVal.Call([]reflect.Value{
				sliceVal.Index(i)})
			for _, r := range result {
				mapped = append(mapped, r.Interface())
			}
		}
	}
	return
}
func main() {
	names := []string{"Alice", "Bob", "Charlie"}
	results := mapSlice(names, strings.ToUpper)
	Printfln("Results: %v", results)
}
*/

//Создание и вызов новых типов функций и значений
// var 1 answer - Results: [ALICE BOB CHARLIE]
/*
func mapSlice(slice interface{}, mapper interface{}) (mapped []interface{}) {
	sliceVal := reflect.ValueOf(slice)
	mapperVal := reflect.ValueOf(mapper)
	mapped = []interface{}{}
	if sliceVal.Kind() == reflect.Slice && mapperVal.Kind() == reflect.Func {
		paramTypes := []reflect.Type{sliceVal.Type().Elem()}
		resultTypes := []reflect.Type{}
		for i := 0; i < mapperVal.Type().NumOut(); i++ {
			resultTypes = append(resultTypes,
				mapperVal.Type().Out(i))
		}
		expectedFuncType := reflect.FuncOf(paramTypes,
			resultTypes, mapperVal.Type().IsVariadic())
		if mapperVal.Type() == expectedFuncType {
			for i := 0; i < sliceVal.Len(); i++ {
				result := mapperVal.Call([]reflect.Value{
					sliceVal.Index(i)})
				for _, r := range result {
					mapped = append(mapped, r.Interface())
				}
			}
		} else {
			Printfln("Function type not as expected")
		}
	}
	return
}
func main() {
	names := []string{"Alice", "Bob", "Charlie"}
	results := mapSlice(names, strings.ToUpper)
	Printfln("Results: %v", results)
}
*/
// var 2 answer -
//Lowercase Results: [alice bob charlie]
//Increment Results: [280 49.95 20.5]
//Price Results: [$279.00 $48.95 $19.50]
/*
func makeMapperFunc(mapper interface{}) interface{} {
	mapVal := reflect.ValueOf(mapper)
	if mapVal.Kind() == reflect.Func && mapVal.Type().NumIn() == 1 && mapVal.Type().NumOut() == 1 {
		inType := reflect.SliceOf(mapVal.Type().In(0))
		inTypeSlice := []reflect.Type{inType}
		outType := reflect.SliceOf(mapVal.Type().Out(0))
		outTypeSlice := []reflect.Type{outType}
		funcType := reflect.FuncOf(inTypeSlice, outTypeSlice, false)
		funcVal := reflect.MakeFunc(funcType,
			func(params []reflect.Value) (results []reflect.Value) {
				srcSliceVal := params[0]
				resultsSliceVal := reflect.MakeSlice(outType, srcSliceVal.Len(), 10)
				for i := 0; i < srcSliceVal.Len(); i++ {
					r := mapVal.Call([]reflect.Value{
						srcSliceVal.Index(i)})
					resultsSliceVal.Index(i).Set(r[0])
				}
				results = []reflect.Value{resultsSliceVal}
				return
			})
		return funcVal.Interface()
	}
	Printfln("Unexpected types")
	return nil
}

func main() {
	lowerStringMapper := makeMapperFunc(strings.ToLower).(func([]string) []string)
	names := []string{"Alice", "Bob", "Charlie"}
	results := lowerStringMapper(names)
	Printfln("Lowercase Results: %v", results)
	incrementFloatMapper := makeMapperFunc(func(val float64) float64 {
		return val + 1
	}).(func([]float64) []float64)
	prices := []float64{279, 48.95, 19.50}
	floatResults := incrementFloatMapper(prices)
	Printfln("Increment Results: %v", floatResults)
	floatToStringMapper := makeMapperFunc(func(val float64) string {
		return fmt.Sprintf("$%.2f", val)
	}).(func([]float64) []string)
	Printfln("Price Results: %v", floatToStringMapper(prices))
}
*/

// Работа с методами
/*
func inspectMethods(s interface{}) {
	sType := reflect.TypeOf(s)
	if sType.Kind() == reflect.Struct || (sType.Kind() ==
		reflect.Ptr &&
		sType.Elem().Kind() == reflect.Struct) {
		Printfln("Type: %v, Methods: %v", sType,
			sType.NumMethod())
		for i := 0; i < sType.NumMethod(); i++ {
			method := sType.Method(i)
			Printfln("Method name: %v, Type: %v",
				method.Name, method.Type)
		}
	}
}
func main() {
	inspectMethods(Purchase{})
	inspectMethods(&Purchase{})
}
*/

// Вызов методов
// var 1 answer - Type: *main.Product, Method: GetAmount, Results: [$279.00]
/*
func executeFirstVoidMethod(s interface{}) {
	sVal := reflect.ValueOf(s)
	for i := 0; i < sVal.NumMethod(); i++ {
		method := sVal.Type().Method(i)
		if method.Type.NumIn() == 1 {
			results := method.Func.Call([]reflect.Value{sVal})
			Printfln("Type: %v, Method: %v, Results: %v", sVal.Type(), method.Name, results)
			break
		} else {
			Printfln("Skipping method %v %v", method.Name,
				method.Type.NumIn())
		}
	}
}
func main() {
	executeFirstVoidMethod(&Product{Name: "Kayak", Price: 279})
}
*/
// var 2 answer Type: *main.Product, Method: GetAmount, Results: [$279.00]
/*
func executeFirstVoidMethod(s interface{}) {
	sVal := reflect.ValueOf(s)
	for i := 0; i < sVal.NumMethod(); i++ {
		method := sVal.Method(i)
		if method.Type().NumIn() == 0 {
			results := method.Call([]reflect.Value{})
			Printfln("Type: %v, Method: %v, Results: %v",
				sVal.Type(), sVal.Type().Method(i).Name, results)
			break
		} else {
			Printfln("Skipping method %v %v",
				sVal.Type().Method(i).Name,
				method.Type().NumIn())
		}
	}
}
func main() {
	executeFirstVoidMethod(&Product{Name: "Kayak", Price: 279})
}
*/

// Работы с интерфейсами
/*
func checkImplementation(check interface{}, targets ...interface{}) {
	checkType := reflect.TypeOf(check)
	if checkType.Kind() == reflect.Ptr && checkType.Elem().Kind() == reflect.Interface {
		checkType := checkType.Elem()
		for _, target := range targets {
			targetType := reflect.TypeOf(target)
			Printfln("Type %v implements %v: %v", targetType, checkType, targetType.Implements(checkType))
		}
	}
}
func main() {
	currencyItemType := (*CurrencyItem)(nil)
	checkImplementation(currencyItemType, &Product{}, &Purchase{})
}
*/

// Получение базовых значений из интерфейсов
/*
type Wrapper struct {
	NamedItem
}

func getUnderlying(item Wrapper, fieldName string) {
	itemVal := reflect.ValueOf(item)
	fieldVal := itemVal.FieldByName(fieldName)
	Printfln("Field Type: %v", fieldVal.Type())
	if fieldVal.Kind() == reflect.Interface {
		Printfln("Underlying	Type: %v", fieldVal.Elem().Type())
	}
}
func main() {
	getUnderlying(Wrapper{NamedItem: &Product{}}, "NamedItem")
}
*/
// Изучение методов интерфейса
/*
type Wrapper struct {
	NamedItem
}

func getUnderlying(item Wrapper, fieldName string) {
	itemVal := reflect.ValueOf(item)
	fieldVal := itemVal.FieldByName(fieldName)
	Printfln("Field Type: %v", fieldVal.Type())
	for i := 0; i < fieldVal.Type().NumMethod(); i++ {
		method := fieldVal.Type().Method(i)
		Printfln("Interface Method: %v, Exported: %v", method.Name, method.PkgPath == "")
	}
	Printfln("--------")
	if fieldVal.Kind() == reflect.Interface {
		Printfln("Underlying Type: %v", fieldVal.Elem().Type())
		for i := 0; i < fieldVal.Elem().Type().NumMethod(); i++ {
			method := fieldVal.Elem().Type().Method(i)
			Printfln("Underlying Method: %v", method.Name)
		}
	}
}
func main() {
	getUnderlying(Wrapper{NamedItem: &Product{}}, "NamedItem")
}
*/
//Работа с типами каналов
/*
func inspectChannel(channel interface{}) {
	channelType := reflect.TypeOf(channel)
	if channelType.Kind() == reflect.Chan {
		Printfln("Type %v, Direction: %v", channelType.Elem(), channelType.ChanDir())
	}
}
func main() {
	var c chan<- string
	inspectChannel(c)
}
*/
//Работа со значениями канала
/*
func sendOverChannel(channel interface{}, data interface{}) {
	channelVal := reflect.ValueOf(channel)
	dataVal := reflect.ValueOf(data)
	if channelVal.Kind() == reflect.Chan &&
		dataVal.Kind() == reflect.Slice &&
		channelVal.Type().Elem() ==
			dataVal.Type().Elem() {
		for i := 0; i < dataVal.Len(); i++ {
			val := dataVal.Index(i)
			channelVal.Send(val)
		}
		channelVal.Close()
	} else {
		Printfln("Unexpected types: %v, %v", channelVal.Type(), dataVal.Type())
	}
}
func main() {
	values := []string{"Alice", "Bob", "Charlie", "Dora"}
	channel := make(chan string)
	go sendOverChannel(channel, values)
	for {
		if val, open := <-channel; open {
			Printfln("Received value: %v", val)
		} else {
			break
		}
	}
}
*/

// Создание новых типов и значений каналов
/*
func createChannelAndSend(data interface{}) interface{} {
	dataVal := reflect.ValueOf(data)
	channelType := reflect.ChanOf(reflect.BothDir,
		dataVal.Type().Elem())
	channel := reflect.MakeChan(channelType, 1)
	go func() {
		for i := 0; i < dataVal.Len(); i++ {
			channel.Send(dataVal.Index(i))
		}
		channel.Close()
	}()
	return channel.Interface()
}

func main() {
	values := []string{"Alice", "Bob", "Charlie", "Dora"}
	channel := createChannelAndSend(values).(chan string)
	for {
		if val, open := <-channel; open {
			Printfln("Received value: %v", val)
		} else {
			break
		}
	}
}
*/
//Выбор из нескольких каналов
func createChannelAndSend(data interface{}) interface{} {
	dataVal := reflect.ValueOf(data)
	channelType := reflect.ChanOf(reflect.BothDir,
		dataVal.Type().Elem())
	channel := reflect.MakeChan(channelType, 1)
	go func() {
		for i := 0; i < dataVal.Len(); i++ {
			channel.Send(dataVal.Index(i))
		}
		channel.Close()
	}()
	return channel.Interface()
}
func readChannels(channels ...interface{}) {
	channelsVal := reflect.ValueOf(channels)
	cases := []reflect.SelectCase{}
	for i := 0; i < channelsVal.Len(); i++ {
		cases = append(cases, reflect.SelectCase{
			Chan: channelsVal.Index(i).Elem(),
			Dir:  reflect.SelectRecv,
		})
	}
	for {
		caseIndex, val, ok := reflect.Select(cases)
		if ok {
			Printfln("Value read: %v, Type: %v", val,
				val.Type())
		} else {
			if len(cases) == 1 {
				Printfln("All channels closed.")
				return
			}
			cases = append(cases[:caseIndex],
				cases[caseIndex+1:]...)
		}
	}
}
func main() {
	values := []string{"Alice", "Bob", "Charlie", "Dora"}
	channel := createChannelAndSend(values).(chan string)
	cities := []string{"London", "Rome", "Paris"}
	cityChannel := createChannelAndSend(cities).(chan string)
	prices := []float64{279, 48.95, 19.50}
	priceChannel := createChannelAndSend(prices).(chan float64)
	readChannels(channel, cityChannel, priceChannel)
}
