package main

import (
	"fmt"
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

/*Разбор значений времени и строк*/

func PrintTime(label string, t *time.Time) {
	layout := "Day: 02 Month: Jan Year: 2006"
	fmt.Println(label, t.Format(layout))

	fmt.Println(label, t.Format(time.RFC822Z))
}

func main() {
	current := time.Now()
	specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
	unix := time.Unix(1433228090, 0)

	PrintTime("Current", &current)
	PrintTime("Specific", &specific)
	PrintTime("UNIX", &unix)
}
