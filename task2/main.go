package main

import (
	"fmt"
	"golang_task/util"
	"sync"
	"sync/atomic"
	"time"
)

func add(c *int) int {
	return *c + 10
}

func multiplication(a *[]int) []int {
	b := *a
	for i := 0; i < len(b); i++ {
		b[i] *= 2
	}
	return b
}

func gorout() {
	var wg = sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			if i%2 == 1 {
				fmt.Printf("协程1: %d\n", i)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Printf("协程2: %d\n", i)
			}
		}
	}()
	wg.Wait()
}

func scheduler() {
	task1 := func() {
		time.Sleep(2 * time.Second)
		fmt.Println("任务1完成")
	}

	task2 := func() {
		time.Sleep(3 * time.Second)
		fmt.Println("任务2完成")
	}

	task3 := func() {
		time.Sleep(5 * time.Second)
		fmt.Println("任务3完成")
	}

	task4 := func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("任务4完成")
	}

	tasks := []func(){task1, task2, task3, task4}

	wg := sync.WaitGroup{}

	startTime := time.Now()

	for i, task := range tasks {
		wg.Add(1)
		go func(id int, t func()) {
			defer wg.Done()
			taskStart := time.Now()
			fmt.Printf("任务%d开始执行\n", id+1)
			t()
			fmt.Printf("任务%d耗时:%v\n", id+1, time.Since(taskStart))
		}(i, task)
	}
	wg.Wait()
	fmt.Printf("任务总耗时:%v\n", time.Since(startTime))
}

func main() {
	util.Placeholder()
	//=== 指针 =========================================================================
	//1. 编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
	//a := 5
	//fmt.Println(add(&a))

	//2. 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
	//b := []int{2, 4, 6, 8, 10}
	//fmt.Println(multiplication(&b))

	//=== Goroutine =========================================================================
	//1. 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
	//gorout()

	//2. 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
	//scheduler()

	//=== 面向对象 =========================================================================
	//1. 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
	//rectangle := util.Rectangle{
	//	Width:  5,
	//	Height: 8,
	//}
	//circle := util.Circle{Radius: 4}
	//recArea := rectangle.Area()
	//recPerimeter := rectangle.Perimeter()
	//ccArea := circle.Area()
	//ccPerimeter := circle.Perimeter()
	//fmt.Printf("长方形的面积: %d\n", recArea)
	//fmt.Printf("长方形的周长: %d\n", recPerimeter)
	//fmt.Printf("圆形的面积: %.2f\n", ccArea)
	//fmt.Printf("圆形的面积: %.2f\n", ccPerimeter)

	//2.使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
	//person1 := util.Person{Name: "liuhaitao", Age: 20}
	//person2 := util.Person{Name: "huanghuang", Age: 18}
	//person3 := util.Person{Name: "tingting", Age: 23}
	//persons := []util.Person{person1, person2, person3}
	//var employees []util.Employee
	//for i, person := range persons {
	//	employees = append(employees, util.Employee{EmployeeId: i + 1, Person: person})
	//}
	//for _, employee := range employees {
	//	fmt.Println(employee.PrintInfo())
	//}

	//=== Channel =========================================================================
	//1. 编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
	//var ch = make(chan int)
	//var wg = sync.WaitGroup{}
	//wg.Add(2)
	//go func() {
	//	defer wg.Done()
	//	for i := 1; i <= 10; i++ {
	//		fmt.Printf("发送%d到通道\n", i)
	//		ch <- i
	//	}
	//	fmt.Printf("通道已关闭\n")
	//	close(ch)
	//}()
	//go func() {
	//	defer wg.Done()
	//	for {
	//		a, ok := <-ch
	//		if !ok {
	//			fmt.Println("未接收到数据")
	//			break
	//		}
	//		fmt.Printf("接收到%d\n", a)
	//	}
	//}()
	//wg.Wait()
	//fmt.Println("结束！")

	//2. 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
	//var ch = make(chan int, 10)
	//var wg sync.WaitGroup
	//wg.Add(2)
	//go func() {
	//	defer wg.Done()
	//	defer close(ch)
	//	for i := 1; i <= 100; i++ {
	//		ch <- i
	//		fmt.Printf("生产者发送了%d\n", i)
	//	}
	//}()
	//go func() {
	//	defer wg.Done()
	//	for num := range ch {
	//		fmt.Printf("消费者接收到了%d\n", num)
	//	}
	//}()
	//wg.Wait()
	//fmt.Println("结束！")

	//=== 锁机制 =========================================================================
	//1. 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	//wg := sync.WaitGroup{}
	//counter := util.Counter{}
	//start := time.Now()
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		for j := 0; j < 1000; j++ {
	//			counter.Increment()
	//		}
	//	}()
	//}
	//wg.Wait()
	//num := counter.GetCount()
	//fmt.Printf("最终的值: %d, 花费时间: %v\n", num, time.Since(start))

	//2. 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	var counter2 int32
	var wg sync.WaitGroup
	var start = time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&counter2, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Printf("final_count: %d, 花费时间:%v\n", atomic.LoadInt32(&counter2), time.Since(start))

}
