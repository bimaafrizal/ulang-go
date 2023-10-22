package main

import "sync"

// channel
func ChannelParameter(channel chan string, name string) {
	channel <- name
}

func getChannel(channel chan string) string {
	return <-channel
}

// wait group
func AddDataWait(channel chan int, rangeData int, group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)
	for i := 0; i <= rangeData; i++ {
		channel <- i
	}
	close(channel)
}

func ReceiveDataWait(channel chan int) int {
	return <-channel
}

// mutex
func cobaMutex(i *int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	*i = *i + 1
	mutex.Unlock()
}

// cond
func WaitCondition(value *int, group *sync.WaitGroup, locker *sync.Cond) {
	defer group.Done()
	locker.L.Lock()
	*value = *value + 1
	locker.Wait()
	locker.L.Unlock()
}
