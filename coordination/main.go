package main

import (
	"context"
	"sync"
	"time"
)

/*
func doSum(count int, val *int) {
	for i := 0; i < count; i++ {
		*val++
	}
}

func main() {
	counter := 0
	doSum(5000, &counter)
	Printfln("Total: %v", counter)
}
*/
// Использование групп ожидания
//проблема:
/*
func doSum(count int, val *int) {
	for i := 0; i < count; i++ {
		*val++
	}
}

func main() {
	counter := 0
	go doSum(5000, &counter)
	Printfln("Total: %v", counter)
}

// answer - Total: 0
*/
// решение
/*
var waitGroup = sync.WaitGroup{}

func doSum(count int, val *int) {
	for i := 0; i < count; i++ {
		*val++
	}
	waitGroup.Done()
}

func main() {
	counter := 0
	waitGroup.Add(1)
	go doSum(5000, &counter)
	waitGroup.Wait()
	Printfln("Total: %v", counter)
}

//answer - Total: 5000
*/

// Использование взаимного исключения
/*
//проблема:на некоторых компьютерах не выведется правильный ответ
var waitGroup = sync.WaitGroup{}

func doSum(count int, val *int) {
	time.Sleep(time.Second)
	for i := 0; i < count; i++ {
		*val++
	}
	waitGroup.Done()
}
func main() {
	counter := 0
	numRoutines := 3
	waitGroup.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go doSum(5000, &counter)
	}
	waitGroup.Wait()
	Printfln("Total: %v", counter)
}
// решение

var waitGroup = sync.WaitGroup{}
var mutex = sync.Mutex{}

func doSum(count int, val *int) {
	time.Sleep(time.Second)
	for i := 0; i < count; i++ {
		mutex.Lock()
		*val++
		mutex.Unlock()
	}
	waitGroup.Done()
}
func main() {
	counter := 0
	numRoutines := 3
	waitGroup.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go doSum(5000, &counter)
	}
	waitGroup.Wait()
	Printfln("Total: %v", counter)
}

*/
//Использование мьютекса чтения-записи
/*
var waitGroup = sync.WaitGroup{}
var rwmutex = sync.RWMutex{}
var squares = map[int]int{}

func calculateSquares(max, iterations int) {
	for i := 0; i < iterations; i++ {
		val := rand.Intn(max)
		rwmutex.RLock()
		square, ok := squares[val]
		rwmutex.RUnlock()
		if ok {
			Printfln("Cached value: %v = %v", val, square)
		} else {
			rwmutex.Lock()
			if _, ok := squares[val]; !ok {
				squares[val] = int(math.Pow(float64(val), 2))
				Printfln("Added value: %v = %v", val,
					squares[val])
			}
			rwmutex.Unlock()
		}
	}
	waitGroup.Done()
}
func main() {
	rand.Seed(time.Now().UnixNano())
	numRoutines := 3
	waitGroup.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go calculateSquares(10, 5)
	}
	waitGroup.Wait()
	Printfln("Cached values: %v", len(squares))
}
*/

//Использование условий для координации горутин
/*
var waitGroup = sync.WaitGroup{}
var rwmutex = sync.RWMutex{}

// var readyCond = sync.NewCond(rwmutex.RLocker()) - вызовы чередуются
var readyCond = sync.NewCond(&rwmutex) // - вызовы не чередуются
var squares = map[int]int{}

func generateSquares(max int) {
	rwmutex.Lock()
	Printfln("Generating data...")
	for val := 0; val < max; val++ {
		squares[val] = int(math.Pow(float64(val), 2))
	}
	rwmutex.Unlock()
	Printfln("Broadcasting condition")
	readyCond.Broadcast()
	waitGroup.Done()
}
func readSquares(id, max, iterations int) {
	readyCond.L.Lock()
	for len(squares) == 0 {
		readyCond.Wait()
	}
	for i := 0; i < iterations; i++ {
		key := rand.Intn(max)
		Printfln("#%v Read value: %v = %v", id, key,
			squares[key])
		time.Sleep(time.Millisecond * 100)
	}
	readyCond.L.Unlock()
	waitGroup.Done()
}
func main() {
	rand.Seed(time.Now().UnixNano())
	numRoutines := 2
	waitGroup.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go readSquares(i, 10, 5)
	}
	waitGroup.Add(1)
	go generateSquares(10)
	waitGroup.Wait()
	Printfln("Cached values: %v", len(squares))
}
*/

//Обеспечение однократного выполнения функции
/*
var waitGroup = sync.WaitGroup{}
var once = sync.Once{}
var squares = map[int]int{}

func generateSquares(max int) {
	Printfln("Generating data...")
	for val := 0; val < max; val++ {
		squares[val] = int(math.Pow(float64(val), 2))
	}
}
func readSquares(id, max, iterations int) {
	once.Do(func() { generateSquares(max) })
	for i := 0; i < iterations; i++ {
		key := rand.Intn(max)
		Printfln("#%v Read value: %v = %v", id, key,
			squares[key])
		time.Sleep(time.Millisecond * 100)
	}
	waitGroup.Done()
}
func main() {
	rand.Seed(time.Now().UnixNano())
	numRoutines := 2
	waitGroup.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go readSquares(i, 10, 5)
	}
	waitGroup.Wait()
	Printfln("Cached values: %v", len(squares))
}
*/

// Использование контекстов
/*
func processRequest(wg *sync.WaitGroup, count int) {
	total := 0
	for i := 0; i < count; i++ {
		Printfln("Processing request: %v", total)
		total++
		time.Sleep(time.Millisecond * 250)
	}
	Printfln("Request processed...%v", total)
	wg.Done()
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	Printfln("Request dispatched...")
	go processRequest(&waitGroup, 10)
	waitGroup.Wait()
}
*/

// Отмена запроса
/*
func processRequest(ctx context.Context, wg *sync.WaitGroup,
	count int) {
	total := 0
	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			Printfln("Stopping processing - request cancelled")
			goto end
		default:
			Printfln("Processing request: %v", total)
			total++
			time.Sleep(time.Millisecond * 250)
		}
	}
	Printfln("Request processed...%v", total)
end:
	wg.Done()
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	Printfln("Request dispatched...")
	ctx, cancel := context.WithCancel(context.Background())
	go processRequest(ctx, &waitGroup, 10)
	time.Sleep(time.Second)
	Printfln("Canceling request")
	cancel()
	waitGroup.Wait()
}
*/

// Установка крайнего срока
/*
func processRequest(ctx context.Context, wg *sync.WaitGroup,
	count int) {
	total := 0
	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.Canceled {
				Printfln("Stopping processing - request cancelled")
			} else {
				Printfln("Stopping processing - deadline reached")
			}
			goto end
		default:
			Printfln("Processing request: %v", total)
			total++
			time.Sleep(time.Millisecond * 250)
		}
	}
	Printfln("Request processed...%v", total)
end:
	wg.Done()
}
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	Printfln("Request dispatched...")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	go processRequest(ctx, &waitGroup, 10)
	waitGroup.Wait()
}
*/

// Предоставление данных запроса
const (
	countKey = iota
	sleepPeriodKey
)

func processRequest(ctx context.Context, wg *sync.WaitGroup) {
	total := 0
	count := ctx.Value(countKey).(int)
	sleepPeriod := ctx.Value(sleepPeriodKey).(time.Duration)
	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.Canceled {
				Printfln("Stopping processing - request cancelled")
			} else {
				Printfln("Stopping processing - deadline reached")
			}
			goto end
		default:
			Printfln("Processing request: %v", total)
			total++
			time.Sleep(sleepPeriod)
		}
	}
	Printfln("Request processed...%v", total)
end:
	wg.Done()
}
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	Printfln("Request dispatched...")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	ctx = context.WithValue(ctx, countKey, 4)
	ctx = context.WithValue(ctx, sleepPeriodKey, time.Millisecond*250)
	go processRequest(ctx, &waitGroup)
	waitGroup.Wait()
}
