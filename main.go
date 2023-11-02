package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// Случайное целочисленное число
func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	rand.Seed(time.Now().Unix())
	c1 := make(chan int)
	c2 := make(chan int)
	countOfc1 := 0
	countOfc2 := 0
	var wg sync.WaitGroup
	//Отправитель сообщений
	sender := func() {
		for {
			//С помощью этой переменной будем определять
			//в какой канал отправлять сообщение
			chance := RandInt(1, 100)
			//Если число в chance меньше либо равно 50, отправляем сообщение в канал c1,
			//если больше 50, то в канал c2
			if chance <= 50 {
				c1 <- RandInt(1, 100)
			} else {
				chance = RandInt(1, 100)
				if chance <= 50 {
					c2 <- RandInt(1, 100)
				} else {
					continue
				}
			}
		}
	}

	//Принимающий сообщения
	receiving := func() {
		for {
			select {
			case num := <-c1:
				{
					countOfc1++
					fmt.Println("Канал с1 принял "+strconv.Itoa(countOfc1)+" сообщений, сообщение из канала: ", num)
				}
			case num := <-c2:
				{
					countOfc2++
					fmt.Println("Канал с2 принял "+strconv.Itoa(countOfc1)+" сообщений, сообщение из канала: ", num)
				}
			default:
				fmt.Println("Сообщений не поступило")
			}
		}
	}

	go sender()
	go receiving()
	wg.Add(2)
	wg.Wait()
}
