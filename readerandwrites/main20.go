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

//без буферизации

// func main() {
// 	text := "It was a boat. A small boat."

// 	var builder strings.Builder
// 	var writer = NewCustomWriter(&builder)
// 	for i := 0; true; {
// 		end := i + 5
// 		if end >= len(text) {
// 			writer.Write([]byte(text[i:]))
// 			break
// 		}
// 		writer.Write([]byte(text[i:end]))
// 		i = end
// 	}
// 	Printfln("Written data: %v", builder.String())
// }

// с буферизацией

// func main() {
// 	text := "It was a boat. A small boat."

// 	var builder strings.Builder
// 	var writer = bufio.NewWriterSize(NewCustomWriter(&builder), 20)
// 	for i := 0; true; {
// 		end := i + 5
// 		if end >= len(text) {
// 			writer.Write([]byte(text[i:]))
// 			writer.Flush()
// 			break
// 		}
// 		writer.Write([]byte(text[i:end]))
// 		i = end
// 	}
// 	Printfln("Written data: %v", builder.String())
// }

/*Форматирование и сканирование с помощью средств чтения и записи*/

/*Сканирование значений из считывателя*/

// func scanFromReader(reader io.Reader, template string, vals ...interface{}) (int, error) {
// 	return fmt.Fscanf(reader, template, vals...)
// }

// func main() {
// 	reader := strings.NewReader("Kayak Watersports $279.00")

// 	var name, category string
// 	var price float64
// 	scanTemplate := "%s %s $%f"

// 	_, err := scanFromReader(reader, scanTemplate, &name, &category, &price)
// 	if err != nil {
// 		Printfln("Error: %v", err.Error())
// 	} else {
// 		Printfln("Name %v", name)
// 		Printfln("Category %v", category)
// 		Printfln("Price %.2f", price)
// 	}
// }

// func scanSingle(reader io.Reader, val interface{}) (int, error) {
// 	return fmt.Fscan(reader, val)
// }

// func main() {
// 	reader := strings.NewReader("Kayak Watersports $279.00")

// 	for {
// 		var str string
// 		_, err := scanSingle(reader, &str)
// 		if err != nil {
// 			if err != io.EOF {
// 				Printfln("Error: %v", err.Error())
// 			}
// 			break
// 		}
// 		Printfln("Value: %v", str)
// 	}
// }

/*Запись отформатированных строк в Writer*/

// func writeFormatted(writer io.Writer, template string, vals ...interface{}) {
// 	fmt.Fprintf(writer, template, vals...)
// }

// func main() {
// 	var writer strings.Builder
// 	template := "Name: %s, Category: %s, Price: %.2f"
// 	writeFormatted(&writer, template, "Kayak", "Watersports", float64(279))
// 	fmt.Println(writer.String())
// }

/*Использование Replace с Writer*/

// func writeReplaced(writer io.Writer, str string, subs ...string) {
// 	replacer := strings.NewReplacer(subs...)
// 	replacer.WriteString(writer, str)
// }

// func main() {
// 	text := "It was a boat. A small boat."
// 	subs := []string{"boat", "kayak", "small", "huge"}

// 	var writer strings.Builder
// 	writeReplaced(&writer, text, subs...)
// 	fmt.Println(writer.String())
// }
