package main

import (
	"time"
)

// func main() {
// 	Printfln("Hello, Dates and Times")
// }

/*Представление дат и времени*/

// func PrintTime(label string, t *time.Time) {
// 	Printfln("%s: Day: %v: Month: %v Year: %v", label, t.Day(), t.Month(), t.Year())
// }

// func main() {
// 	current := time.Now()
// 	specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
// 	unix := time.Unix(1433228090, 0)

// 	PrintTime("Current", &current)
// 	PrintTime("Specific", &specific)
// 	PrintTime("UNIX", &unix)
// }

/*
Current: Day: 29: Month: January Year: 2024
Specific: Day: 9: Month: June Year: 1995
UNIX: Day: 2: Month: June Year: 2015
*/

/*Форматирование времени как строк*/

// func PrintTime(label string, t *time.Time) {
// 	layout := "Day: 02 Month: Jan Year: 2006"
// 	fmt.Println(label, t.Format(layout))

// 	fmt.Println(label, t.Format(time.RFC822Z))
// }

// func main() {
// 	current := time.Now()
// 	specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
// 	unix := time.Unix(1433228090, 0)

// 	PrintTime("Current", &current)
// 	PrintTime("Specific", &specific)
// 	PrintTime("UNIX", &unix)
// }
/*
Current Day: 29 Month: Jan Year: 2024
Specific Day: 09 Month: Jun Year: 1995
UNIX Day: 02 Month: Jun Year: 2015

Current Day: 29 Month: Jan Year: 2024
Current 29 Jan 24 17:44 +0700
Specific Day: 09 Month: Jun Year: 1995
Specific 09 Jun 95 00:00 +0800
UNIX Day: 02 Month: Jun Year: 2015
UNIX 02 Jun 15 13:54 +0700
*/

/*Разбор значений времени из строк*/

// func PrintTime(label string, t *time.Time) {
// 	fmt.Println(label, t.Format(time.RFC822Z))
// }

// func main() {
// 	layout := "2006-Jan-02"
// 	dates := []string{
// 		"1995-Jun-09",
// 		"2015-Jun-02",
// 	}
// 	for _, d := range dates {
// 		time, err := time.Parse(layout, d)
// 		if err == nil {
// 			PrintTime("parsed", &time)
// 		} else {
// 			Printfln("Error: %s", err.Error())
// 		}
// 	}
// }

/*
parsed 09 Jun 95 00:00 +0000
parsed 02 Jun 15 00:00 +0000
*/

/*Использование предопределенных макетов дат*/

// func PrintTime(label string, t *time.Time) {
// 	fmt.Println(label, t.Format(time.RFC822Z))
// }

// func main() {
// 	dates := []string{
// 		"09 Jun 95 00:00 GMT",
// 		"02 Jun 15 00:00 GMT",
// 	}
// 	for _, d := range dates {
// 		time, err := time.Parse(time.RFC822, d)
// 		if err == nil {
// 			PrintTime("parsed", &time)
// 		} else {
// 			Printfln("Error: %s", err.Error())
// 		}
// 	}
// }
/*
parsed 09 Jun 95 00:00 +0000
parsed 02 Jun 15 00:00 +000
*/

/*Указание разбора местоположения*/

// func PrintTime(label string, t *time.Time) {
// 	fmt.Println(label, t.Format(time.RFC822Z))
// }

// func main() {
// 	layout := "02 Jan 06 15:04"
// 	date := "09 Jun 95 19:30"

// 	london, lonerr := time.LoadLocation("Europe/London")
// 	newyork, nycerr := time.LoadLocation("America/New_York")

// 	if lonerr == nil && nycerr == nil {
// 		nolocation, _ := time.Parse(layout, date)
// 		londonTime, _ := time.ParseInLocation(layout, date, london)
// 		newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
// 		PrintTime("No location:", &nolocation)
// 		PrintTime("London:", &londonTime)
// 		PrintTime("New York:", &newyorkTime)
// 	} else {
// 		fmt.Println(lonerr.Error(), nycerr.Error())
// 	}
// }
/*
No location: 09 Jun 95 19:30 +0000
London: 09 Jun 95 19:30 +0100
New York: 09 Jun 95 19:30 -0400
*/

/*Использование локального местоположения*/

// func PrintTime(label string, t *time.Time) {
// 	fmt.Println(label, t.Format(time.RFC822Z))
// }

// func main() {
// 	layout := "02 Jan 06 15:04"
// 	date := "09 Jun 95 19:30"

// 	london, lonerr := time.LoadLocation("Europe/London")
// 	newyork, nycerr := time.LoadLocation("America/New_York")
// 	local, _ := time.LoadLocation("Local")

// 	if lonerr == nil && nycerr == nil {
// 		nolocation, _ := time.Parse(layout, date)
// 		londonTime, _ := time.ParseInLocation(layout, date, london)
// 		newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
// 		localTime, _ := time.ParseInLocation(layout, date, local)
// 		PrintTime("No location:", &nolocation)
// 		PrintTime("London:", &londonTime)
// 		PrintTime("New York:", &newyorkTime)
// 		PrintTime("Local:", &localTime)
// 	} else {
// 		fmt.Println(lonerr.Error(), nycerr.Error())
// 	}
// }

/*
No location: 09 Jun 95 19:30 +0000
London: 09 Jun 95 19:30 +0100
New York: 09 Jun 95 19:30 -0400
Local: 09 Jun 95 19:30 +0800
*/

/*непосредственное указание часовых поясов*/

// func PrintTime(label string, t *time.Time) {
// 	fmt.Println(label, t.Format(time.RFC822Z))
// }

// func main() {
// 	layout := "02 Jan 06 15:04"
// 	date := "09 Jun 95 19:30"

// 	london := time.FixedZone("BST", 1*60*60)
// 	newyork := time.FixedZone("EDT", -4*60*60)
// 	local := time.FixedZone("Local", 0)

// 	nolocation, _ := time.Parse(layout, date)
// 	londonTime, _ := time.ParseInLocation(layout, date, london)
// 	newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
// 	localTime, _ := time.ParseInLocation(layout, date, local)
// 	PrintTime("No location:", &nolocation)
// 	PrintTime("London:", &londonTime)
// 	PrintTime("New York:", &newyorkTime)
// 	PrintTime("Local:", &localTime)
// }

/*
No location: 09 Jun 95 19:30 +0000
London: 09 Jun 95 19:30 +0100
New York: 09 Jun 95 19:30 -0400
Local: 09 Jun 95 19:30 +0000
*/

/*Управление значениями времени*/

// func main() {
// 	t, err := time.Parse(time.RFC822, "09 Jun 95 04:59 BST")
// 	if err == nil {
// 		Printfln("After: %v", t.After(time.Now()))
// 		Printfln("Round: %v", t.Round(time.Hour))
// 		Printfln("Truncate: %v", t.Truncate(time.Hour))
// 	} else {
// 		fmt.Println((err.Error()))
// 	}
// }
/*
After: false
Round: 1995-06-09 05:00:00 +0000 BST
Truncate: 1995-06-09 04:00:00 +0000 BST
*/

// func main() {
// 	t1, _ := time.Parse(time.RFC822Z, "09 Jun 95 04:59 +0100")
// 	t2, _ := time.Parse(time.RFC822Z, "08 Jun 95 23:59 -0400")

// 	Printfln("Equal Method: %v", t1.Equal(t2))
// 	Printfln("Equality Operation: %v", t1 == t2)
// }

/*
Equal Method: true
Equality Operation: false
*/

/*Представление продолжительности*/

// func main() {

// 	var d time.Duration = time.Hour + (30 * time.Minute)

// 	Printfln("Hours: %v", d.Hours())
// 	Printfln("Mins: %v", d.Minutes())
// 	Printfln("Seconds: %v", d.Seconds())
// 	Printfln("Milliseconds: %v", d.Milliseconds())

// 	rounded := d.Round(time.Hour)
// 	Printfln("Rounded Hours: %v", rounded.Hours())
// 	Printfln("Rounded Mins: %v", rounded.Minutes())

// 	trunc := d.Truncate(time.Hour)
// 	Printfln("Truncated Hours: %v", trunc.Hours())
// 	Printfln("Truncated Mins: %v", trunc.Minutes())
// }
/*
Hours: 1.5
Mins: 90
Seconds: 5400
Milliseconds: 5400000
Rounded Hours: 2
Rounded Mins: 120
Truncated Hours: 1
Truncated Mins: 60
*/

/*Создание продолжительности относительно времени */

// func main() {
// 	toYears := func(d time.Duration) int {
// 		return int(d.Hours() / (24 * 365))
// 	}

// 	future := time.Date(2051, 0, 0, 0, 0, 0, 0, time.Local)
// 	past := time.Date(1965, 0, 0, 0, 0, 0, 0, time.Local)

// 	Printfln("Future: %v", toYears(time.Until(future)))
// 	Printfln("Past: %v", toYears(time.Since(past)))
// }
/*
Future: 26
Past: 59
*/

/*Создание длительности из строк*/

// func main() {
// 	d, err := time.ParseDuration("1h30m")
// 	if err == nil {
// 		Printfln("Hours: %v", d.Hours())
// 		Printfln("Mins: %v", d.Minutes())
// 		Printfln("Seconds: %v", d.Seconds())
// 		Printfln("Milliseconds: %v", d.Milliseconds())
// 	} else {
// 		fmt.Println(err.Error())
// 	}
// }
/*
Hours: 1.5
Mins: 90
Seconds: 5400
Milliseconds: 5400000
*/

/*Исппользование функций времени для горутин и каналов*/

/*Перевод горутины в сон Sleep(duration) */

// func writeToChannel(channel chan<- string) {
// 	names := []string{"Alice", "Bob", "Charlie", "Dora"}
// 	for _, name := range names {
// 		channel <- name
// 		time.Sleep(time.Second * 1)
// 	}
// 	close(channel)
// }

// func main() {
// 	naneChannel := make(chan string)

// 	go writeToChannel(naneChannel)

// 	for name := range naneChannel {
// 		Printfln("Read name: %v", name)
// 	}
// }

/*Отсрочка выполнения функции*/

// func writeToChannel(channel chan<- string) {
// 	names := []string{"Alice", "Bob", "Charlie", "Dora"}
// 	for _, name := range names {
// 		channel <- name
// 	}
// 	close(channel)
// }

// func main() {
// 	naneChannel := make(chan string)

// 	time.AfterFunc(time.Second*5, func() {
// 		writeToChannel(naneChannel)
// 	})

// 	for name := range naneChannel {
// 		Printfln("Read name: %v", name)
// 	}
// }

/*получение уведомлений по времени*/

// func writeToChannel(channel chan<- string) {
// 	Printfln("Waiting for initial duration...")
// 	_ = <-time.After(time.Second * 2)
// 	Printfln("Initial duration elapsed.")

// 	names := []string{"Alice", "Bob", "Charlie", "Dora"}
// 	for _, name := range names {
// 		channel <- name
// 		time.Sleep(time.Second * 1)
// 	}
// 	close(channel)
// }

// func main() {
// 	naneChannel := make(chan string)

// 	go writeToChannel(naneChannel)

// 	for name := range naneChannel {
// 		Printfln("Read name: %v", name)
// 	}
// }

/*Использование уведомлений в качестве тайм-футов в опреаторах Select*/

// func writeToChannel(channel chan<- string) {
// 	Printfln("Waiting for initial duration...")
// 	_ = <-time.After(time.Second * 2)
// 	Printfln("Initial duration elapsed.")

// 	names := []string{"Alice", "Bob", "Charlie", "Dora"}
// 	for _, name := range names {
// 		channel <- name
// 		time.Sleep(time.Second * 3)
// 	}
// 	close(channel)
// }

// func main() {
// 	nameChannel := make(chan string)

// 	go writeToChannel(nameChannel)

// 	channelOpen := true

// 	for channelOpen {
// 		Printfln("Starting channel read")
// 		select {
// 		case name, ok := <-nameChannel:
// 			if !ok {
// 				channelOpen = false
// 				break
// 			} else {
// 				Printfln("read name: %v", name)
// 			}
// 		case <-time.After(time.Second * 2):
// 			Printfln("Timeout")
// 		}
// 	}
// }

/*остановка и сброс таймеров*/

// func writeToChannel(channel chan<- string) {
// 	timer := time.NewTimer(time.Minute * 10)
// 	go func() {
// 		time.Sleep(time.Second * 2)
// 		Printfln("Resseting timer")
// 		timer.Reset(time.Second)
// 	}()
// 	Printfln("Waiting for initial duration...")
// 	_ = <-timer.C
// 	Printfln("Initial duration elapsed.")

// 	names := []string{"Alice", "Bob", "Charlie", "Dora"}
// 	for _, name := range names {
// 		channel <- name
// 	}
// 	close(channel)
// }

// func main() {
// 	nameChannel := make(chan string)

// 	go writeToChannel(nameChannel)
// 	for name := range nameChannel {
// 		Printfln("Read name: %v", name)
// 	}
// }

/*
Timer создается с продолжительностью 10 минут. Горутина спит две секунды, затем сбрасывает таймер и начинает выполнение
*/

/*Получение повторяющихся уведомлений*/

// func writeToChannel(nameChannel chan<- string) {
// 	names := []string{"Alice", "Bob", "Charlie", "Dora"}
// 	tickChannel := time.Tick(time.Second)
// 	index := 0

// 	for {
// 		<-tickChannel
// 		nameChannel <- names[index]
// 		index++
// 		if index == len(names) {
// 			index = 0
// 		}
// 	}
// }

func writeToChannel(nameChannel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	ticker := time.NewTicker(time.Second/10)
	index := 0

	for {
		<-ticker.C
		nameChannel <- names[index]
		index++
		if index == len(names) {
			ticker.Stop()
			close(nameChannel)
			break
		}
	}
}

func main() {
	nameChannel := make(chan string)

	go writeToChannel(nameChannel)
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}
