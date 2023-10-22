package main

import (
	"fmt"
	"sync"
	"time"
)

// array
func TotalArray() (int, int) {
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	var total = 0
	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}

	rata2 := total / len(arr)

	return total, rata2
}

func MinMax() (int, int) {
	var numbers [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	min := numbers[0]
	max := numbers[0]

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		} else if numbers[i] < min {
			min = numbers[i]
		}
	}
	return min, max
}

func main() {
	//total, rata2 := TotalArray()
	//min, max := MinMax()
	//println("Total Array: ", total)
	//println("Rata-rata Array: ", rata2)
	//println("Min: ", min)
	//println("Max: ", max)

	//slice & struct
	//var students []Student
	//addData(&students, "Bima", "bima@gmail.com", "123")
	//addData(&students, "Afrizal", "bima@gmail.com", "123")
	//addData(&students, "Malna", "bima@gmail.com", "123")
	//
	//fmt.Println(getDataByName(&students, "Afrizal"))
	//deletetDataByName(&students, "Malna")
	//fmt.Println(students)

	//interface
	//segitiga := Triangle{tinggi: 12, alas: 5}
	//fmt.Println(segitiga.area())
	//
	//lingkaran := Circle{radius: 21}
	//fmt.Println(lingkaran.area())

	//defer panic
	//fmt.Println(divide(3, 2))
	//fmt.Println(divide(3, 0))

	//channel
	//channel := make(chan string)
	//go ChannelParameter(channel, "Bima")
	//data := getChannel(channel)
	//fmt.Println(data)
	//time.Sleep(5 * time.Second)

	//mutex
	//var mutex sync.Mutex
	//var wg sync.WaitGroup
	//x := 0
	//wg.Add(1000)
	//for i := 0; i < 1000; i++ {
	//	go cobaMutex(&x, &mutex, &wg)
	//}
	//wg.Wait()
	//fmt.Println(x)
	//time.Sleep(5 * time.Second)

	// sync wait group
	//ch := make(chan int)
	//var wg sync.WaitGroup
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	for i := 0; i < 1000; i++ {
	//		ch <- i
	//	}
	//	close(ch)
	//}()
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	for num := range ch {
	//		fmt.Println("Received:", num) // Menerima data dari channel
	//	}
	//}()
	//go AddDataWait(ch, 20, &wg)
	//fmt.Println(ch)
	//fmt.Println(ReceiveDataWait(ch))
	//wg.Wait()

	//cond
	var locker sync.Mutex
	cond := sync.NewCond(&locker)
	var wg sync.WaitGroup
	value := 0
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go WaitCondition(&value, &wg, cond)
	}

	//tambahkan waktu jeda
	time.Sleep(time.Millisecond)
	cond.Broadcast()
	wg.Wait()
	fmt.Println(value)
}
