package main

// func main() {
// 	Printfln("Product: %v, Price: %v", Kayak.Name, Kayak.Price)
// }

/*
Понимание средств чтения и записи
пакет io
интерфейсы Reader и Writer
*/

/*Понимание средств чтения*/

// func processData(reader io.Reader) {
// 	b := make([]byte, 2)

// 	for {
// 		count, err := reader.Read(b)
// 		if count > 0 {
// 			Printfln("Read %v bytes: %v", count, string(b[0:count]))
// 		}
// 		if err == io.EOF {
// 			break
// 		}
// 	}
// }

// func main() {
// 	r := strings.NewReader("Kayak")
// 	processData(r)
// }

/*
Read 2 bytes: Ka
Read 2 bytes: ya
Read 1 bytes: k
*/

/*Понимание средств записи*/

// func processData(reader io.Reader, writer io.Writer) {
// 	b := make([]byte, 2)

// 	for {
// 		count, err := reader.Read(b)
// 		if count > 0 {
// 			writer.Write(b[0:count])
// 			Printfln("Read %v bytes: %v", count, string(b[0:count]))
// 		}
// 		if err == io.EOF {
// 			break
// 		}
// 	}
// }

// func main() {
// 	r := strings.NewReader("Kayak")
// 	var builder strings.Builder
// 	processData(r, &builder)
// 	Printfln("String builder contents: %s", builder.String())
// }

/*
Read 2 bytes: Ka
Read 2 bytes: ya
Read 1 bytes: k
String builder contents: Kayak
*/

/*Использование служебных функций для программ чтения и записи*/

// func processData(reader io.Reader, writer io.Writer) {
// 	count, err := io.Copy(writer, reader)
// 	if err == nil {
// 		Printfln("Read %v bytes", count)
// 	} else {
// 		Printfln("error: %v", err.Error())
// 	}
// }

// func main() {
// 	r := strings.NewReader("Kayak")
// 	var builder strings.Builder
// 	processData(r, &builder)
// 	Printfln("String builder contents: %s", builder.String())
// }
/*
Read 5 bytes
String builder contents: Kayak
*/

/*Использование специализированных средств чтения и записи*/

/*Использование пайпов*/

// func main() {
// 	pipeReader, pipeWriter := io.Pipe()

// 	go func() {
// 		GenerateData(pipeWriter)
// 		pipeWriter.Close()
// 	}()
// 	ConsumeData(pipeReader)
// }

/*
Улучшение
В функцию GenerateData добавлено
	if closer, ok := writer.(io.Closer); ok {
		closer.Close()
	}
*/

// func main() {
// 	pipeReader, pipeWriter := io.Pipe()
// 	go GenerateData(pipeWriter)
// 	ConsumeData(pipeReader)
// }

/*Объединение нескольких средств чтения*/

// func main() {
// 	r1 := strings.NewReader("Kayak")
// 	r2 := strings.NewReader("Lifejacket")
// 	r3 := strings.NewReader("Canoe")

// 	concatReader := io.MultiReader(r1, r2, r3)
// 	ConsumeData(concatReader)
// }

/*Объединение нескольких средств записи*/

// func main() {
// 	var w1 strings.Builder
// 	var w2 strings.Builder
// 	var w3 strings.Builder

// 	combinedWriter := io.MultiWriter(&w1, &w2, &w3)

// 	GenerateData(combinedWriter)

// 	Printfln("Writer №1: %v", w1.String())
// 	Printfln("Writer №2: %v", w2.String())
// 	Printfln("Writer №3: %v", w3.String())
// }

/*повторение данных чтения во Writer*/

// func main() {
// 	r1 := strings.NewReader("Kayak")
// 	r2 := strings.NewReader("Lifejacket")
// 	r3 := strings.NewReader("Canoe")

// 	concatReader := io.MultiReader(r1, r2, r3)

// 	var writer strings.Builder
// 	teeReader := io.TeeReader(concatReader, &writer)

// 	ConsumeData(teeReader)
// 	Printfln("Echo data: %v", writer.String())
// }

/*ограничение чтения данных*/

// func main() {
// 	r1 := strings.NewReader("Kayak")
// 	r2 := strings.NewReader("Lifejacket")
// 	r3 := strings.NewReader("Canoe")

// 	concatReader := io.MultiReader(r1, r2, r3)

// 	limited := io.LimitReader(concatReader, 5)
// 	ConsumeData(limited)
// }
/*
вывод с ограничением в 5 байтов
Read data: Ka
Read data: ya
Read data: k
Read data: Kayak
*/

/*Буферизация данных*/

// func main() {
// 	text := "It was a boat. A small boat."

// 	var reader io.Reader = NewCustumReader(strings.NewReader(text))
// 	var writer strings.Builder
// 	slice := make([]byte, 5)
// 	for {
// 		count, err := reader.Read(slice)
// 		if count > 0 {
// 			writer.Write(slice[0:count])
// 		}
// 		if err != nil {
// 			break
// 		}
// 	}
// 	Printfln("Read data: %v", writer.String())
// }

// func main() {
// 	text := "It was a boat. A small boat."

// 	var reader io.Reader = NewCustumReader(strings.NewReader(text))
// 	var writer strings.Builder
// 	slice := make([]byte, 5)

// 	reader = bufio.NewReader(reader)

// 	for {
// 		count, err := reader.Read(slice)
// 		if count > 0 {
// 			writer.Write(slice[0:count])
// 		}
// 		if err != nil {
// 			break
// 		}
// 	}
// 	Printfln("Read data: %v", writer.String())
// }

/*Использование дополнительных методов буферизованного чтения*/

// func main() {
// 	text := "It was a boat. A small boat."

// 	var reader io.Reader = NewCustumReader(strings.NewReader(text))
// 	var writer strings.Builder
// 	slice := make([]byte, 5)

// 	buffered := bufio.NewReader(reader)

// 	for {
// 		count, err := buffered.Read(slice)
// 		if count > 0 {
// 			Printfln("Buffer size: %v, buffered: %v", buffered.Size(), buffered.Buffered())
// 			writer.Write(slice[0:count])
// 		}
// 		if err != nil {
// 			break
// 		}
// 	}
// 	Printfln("Read data: %v", writer.String())
// }

/*Выполнение буферизированной записи*/
