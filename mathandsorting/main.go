package main

// func main() {
// 	Printfln("Hello, Math and Sorting")
// }

// func main() {
// 	val1 := 279.00
// 	val2 := 48.95

// 	//Abc(val) возвращает абсолютное значение
// 	Printfln("Abc: %v", math.Abs(val1))

// 	//Ceil(val0) - возвращает наименьшее целое число, равное или превышающее указанное значение  (результат float64)
// 	Printfln("Ceil: %v", math.Ceil(val2))

// 	//Copysign(x, y) возвращает абсолютное значение х со знаком у
// 	Printfln("Copysign: %v", math.Copysign(val1, -5))

// 	//Floor(val) возвращает наибольшее целое число, равное или превышающее указанное значение  (результат float64)
// 	Printfln("Floor: %v", math.Floor(val2))

// 	//Max(x, y) возвращает самое большое
// 	Printfln("Max: %v", math.Max(val1, val2))

// 	//Min(x, y) возвращает нименьшее
// 	Printfln("Min: %v", math.Min(val1, val2))

// 	//Mod(x, y) возвращает остаток от x/y
// 	Printfln("Mod: %v", math.Mod(val1, val2))

// 	//Pow(x, y) возвращает значение х возведеное в степень у
// 	Printfln("Pow: %v", math.Pow(val1, 2))

// 	//Round(val) округляет до ближайшего целого числа, округляя половинные значения в большую сторону (результат float64)
// 	Printfln("Round: %v", math.Round(val2))

// 	//RoundToEven(val) округляет до ближайшего целого числа, округляя половинные значения до ближайшего четного (результат float64)
// 	Printfln("RoundToEven: %v", math.RoundToEven(val2))
// }

/*
Abc: 279
Ceil: 49
Copysign: -279
Floor: 48
Max: 279
Min: 48.95
Mod: 34.249999999999986
Pow: 77841
Round: 49
RoundToEven: 49
*/

/*
Генерация случайных (псевдослучайных) чисел
Seed(s) - устанавливает начальное значение, используя указанное значение int64
Float32() - генерирует случайное значение float32 в диапазоне от 0 до 1
Float64() - генерирует случайное значение float64 в диапазоне от 0 до 1
Int() - генерирует случайное значение int
Intn(max) - генерирует случайное значение int меньше указанного
Uint32() - генерирует случайное значение uint32
Uint64() -генерирует случайное значение uint64
Shuffle(count, func) - используется для рандомизации порядка элементов
*/

// func main() {
// 	for i := 0; i < 5; i++ {
// 		Printfln("Value %v : %v", i, rand.Int())
// 	}
// }

/*
вывод всегда одинаковый
Value 0 : 5577006791947779410
Value 1 : 8674665223082153551
Value 2 : 6129484611666145821
Value 3 : 4037200794235010051
Value 4 : 3916589616287113937
*/

/*
Now - текущее время
UnixNano - целочисленное значение int64
*/

// func main() {

// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < 5; i++ {
// 		Printfln("Value %v : %v", i, rand.Int())
// 	}
// }

/*
вывод всегда отличается
Value 0 : 163819611726067811
Value 1 : 1021180900046278721
Value 2 : 2994518148756790924
Value 3 : 6309115064855227061
Value 4 : 795408947462603097
*/

/*Генерация случайного числа в определенном диапазоне*/

// func main() {

// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < 5; i++ {
// 		Printfln("Value %v : %v", i, rand.Intn(10))
// 	}
// }

/*
Value 0 : 1
Value 1 : 4
Value 2 : 1
Value 3 : 3
Value 4 : 3
*/

/*функции для указания минимального значения нет, но:*/

// func intRange(min, max int) int {
// 	return rand.Intn(max - min) + min
// }
// func main() {

// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < 5; i++ {
// 		Printfln("Value %v : %v", i, intRange(10, 20))
// 	}
// }

/*
Value 0 : 12
Value 1 : 17
Value 2 : 19
Value 3 : 18
Value 4 : 18
*/

/*Перетасовка элементов*/

// var names = []string{"Alice", "Bob", "Charlie", "Dora", "Edith"}

// func main() {
// 	rand.Seed(time.Now().UnixNano())
// 	rand.Shuffle(len(names), func(first, second int) {
// 		names[first], names[second] = names[second], names[first]
// 	})

// 	for i, name := range names {
// 		Printfln("Index %v : Name: %v", i, name)
// 	}
// }
/*
Index 0 : Name: Dora
Index 1 : Name: Alice
Index 2 : Name: Charlie
Index 3 : Name: Bob
Index 4 : Name: Edith
*/

/*Сортировка данных*/

/*Сортировка числовых и строковых срезов*/

/*
Float64s(slice) - сортирует срез значения float64. Элементы сортируются на месте.
Float64sAreSorted(slice) - true если элементы упорядочены.
Ints(slice) - сортирует срез значения int. Элементы сортируются на месте.
IntsAreSorted(slice) - true если элементы упорядочены.
Strings(slice) - сортирует срез значения string. Элементы сортируются на месте.
StringsAreSorted(slice) - true если элементы упорядочены.
*/

// func main() {
// 	ints := []int{9, 4, 2, -1, 10}
// 	Printfln("ints: %v", ints)
// 	sort.Ints(ints)
// 	Printfln("ints sorted: %v", ints)

// 	floats := []float64{279, 48.95, 19.50}
// 	Printfln("Floats: %v", floats)
// 	sort.Float64s(floats)
// 	Printfln("Floats sorted: %v", floats)

// 	strings := []string{"Kayak", "Lifejacket", "Stadium"}
// 	Printfln("Strings: %v", strings)
// 	if !sort.StringsAreSorted(strings) {
// 		sort.Strings(strings)
// 		Printfln("String Sorted: %v", strings)
// 	} else {
// 		Printfln("Strings Already Sorted: %v", strings)
// 	}
// }

/*
ints: [9 4 2 -1 10]
ints sorted: [-1 2 4 9 10]
Floats: [279 48.95 19.5]
Floats sorted: [19.5 48.95 279]
Strings: [Kayak Lifejacket Stadium]
Strings Already Sorted: [Kayak Lifejacket Stadium]
*/

/*для создания отсортированного среза необходимо использовать встроенные функции make, copy*/
// func main() {
// 	ints := []int{9, 4, 2, -1, 10}

// 	sortedInst := make([]int, len(ints))
// 	copy(sortedInst, ints)

// 	sort.Ints(sortedInst)
// 	Printfln("ints: %v", ints)
// 	Printfln("ints sorted: %v", sortedInst)
// }

/*
ints: [9 4 2 -1 10]
ints sorted: [-1 2 4 9 10]
*/

/*
Поиск отсортированных данных
SearchInts(slice, val)
SearchFloat64(slice, val)
SearchStrings(slice, val)
Search(count, testFunc)
*/

// func main() {
// 	ints := []int{9, 4, 2, -1, 10}

// 	sortedInst := make([]int, len(ints))
// 	copy(sortedInst, ints)

// 	sort.Ints(sortedInst)
// 	Printfln("ints: %v", ints)
// 	Printfln("ints sorted: %v", sortedInst)

// 	indexOf4 := sort.SearchInts(sortedInst, 4)
// 	indexOf3 := sort.SearchInts(sortedInst, 3)

// 	Printfln("index of 4: %v", indexOf4)
// 	Printfln("index of 3: %v", indexOf3)

// 	Printfln("index of 4: %v (present: %v)", indexOf4, sortedInst[indexOf4] == 4)
// 	Printfln("index of 3: %v (present: %v)", indexOf3, sortedInst[indexOf3] == 3)
// }

/*
ints: [9 4 2 -1 10]
ints sorted: [-1 2 4 9 10]
index of 4: 2
index of 3: 2

index of 4: 2 (present: true)
index of 3: 2 (present: false)
*/

/*
Сортировка пользовательских типов данных
len() - возвращает количество элементов, которые будут отсортированы
less(i, j) - возвращает true если i должен быть перед j
Swap(i, j) - меняет местами элементы по указаным индексам

Sort(data)
Stable(data)
IsSorted(data)
Reverse(data)
*/

// func main() {
// 	products := []Product{
// 		{"Kayak", 279},
// 		{"Lifejacket", 49.995},
// 		{"Soccer Ball", 19.50},
// 	}
// 	ProductSlices(products)

// 	for _, p := range products {
// 		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
// 	}
// }

/*
Name: Soccer Ball, Price: 19.50
Name: Lifejacket, Price: 49.99
Name: Kayak, Price: 279.00
*/

/*Сортировка с использованием разных полей*/

// func main() {
// 	products := []Product{
// 		{"Kayak", 279},
// 		{"Lifejacket", 49.995},
// 		{"Soccer Ball", 19.50},
// 	}
// 	ProductSlicesByName(products)

// 	for _, p := range products {
// 		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
// 	}
// }

/*
Name: Kayak, Price: 279.00
Name: Lifejacket, Price: 49.99
Name: Soccer Ball, Price: 19.50
*/

/*Определение функции сравнения*/

func main() {
	products := []Product{
		{"Kayak", 279},
		{"Lifejacket", 49.995},
		{"Soccer Ball", 19.50},
	}
	SortWith(products, func(p1, p2 Product) bool {
		return p1.Name < p2.Name
	})

	for _, p := range products {
		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
	}
}

/*
Name: Kayak, Price: 279.00
Name: Lifejacket, Price: 49.99
Name: Soccer Ball, Price: 19.50
*/
