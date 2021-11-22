package main

import "fmt"

type zhangyun struct {
	name   string
	age    int
	gender string
}

type pig interface {
	eat()
	sleep()
}

func (zy zhangyun) eat() {
	fmt.Println("我会吃")
}

func (zy zhangyun) sleep() {
	fmt.Println("我会睡觉")
}

func main() {
	zhangyun := zhangyun{
		name:   "zhangyun",
		age:    12,
		gender: "female",
	}
	zhangyun.eat()
	zhangyun.sleep()
}
