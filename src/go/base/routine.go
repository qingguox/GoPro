package main

import (
	. "fmt"
	"sync"
	"time"
)

func main() {

	// 并发是指程序同时执行多个任务的能力
	// Go语言支持并发, 通过 GoRoutines 和 Channels 提供了一种简介且高效的方式来实现并发

	// Goroutines: 非阻塞额
	// 	1. go中的并发执行单位, 类似于轻量级的线程.
	//  2. goroutine的调度由Go运行时管理, 用户无需手动分配线程。
	//  3. 使用go关键字启动Goroutine.
	//  4. goroutine是非阻塞的, 可以高效地运行成千上万个 Goroutine

	// Channel: 无缓冲器, 同步.  有缓冲区, 没满, 非阻塞, 满了,阻塞发送
	//  1.Go中用于在Goroutine 之间通信的机制.
	//  2. 支持同步和数据共享, 避免了显式的锁机制
	//  3. 使用chan 关键字创建，通过 <- 操作符发送 和 接受数据

	// Scheduler(调度器):
	// Go的调度器基于 GMP 模型, 调度器会将 Goroutine 分配到系统线程中执行, 把那个通过 M 和 P 的
	// 	配合高效管理并发.
	// G: Goroutine
	// M: 系统线程(Machine)
	// P: 逻辑处理器(Processor)

	// Goroutine
	// 1. 语法: go 函数名(参数列表)
	var cnt = 5
	go sayRoutine(cnt)
	for i := 0; i < cnt; i++ {
		Println("Main")
		time.Sleep(100 * time.Millisecond)
	}
	// 上述结果: main 和 hello 输出没有顺序, 两个goroutine处理

	// 2。 Channel 用于Goroutime之间的数据传输
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	// 因为goroutine两个无顺序, 此时通道内的数据也是无序的
	x, y := <-c, <-c
	Println(x, y, x+y)

	// 3. 通道设置缓冲区:
	// 带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于异步状态，就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。
	// 不过由于缓冲区的大小是有限的，所以还是必须有接收端来接收数据的，否则缓冲区一满，数据发送端就无法再发送数据了。
	// 注意：如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内；如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。接收方在有值可以接收之前会一直阻塞。

	cha := make(chan int, 2)
	cha <- 1
	cha <- 2
	Println("通道出一个: ", <-cha)
	cha <- 3
	Println(<-cha)
	Println(<-cha)

	// 4. 遍历通道与关闭通道
	// v, ok := <-ch
	// 如果通道接受不到数据后 ok 就变为false, 此时通道就可以使用close()函数关闭.
	c3 := make(chan int, 10)
	go fibonacci(cap(c3), c3)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i2 := range c3 {
		Println("i:", i2)
	}

	// 5. select操作, 让一个goroutine等待多个通信操作
	cx := make(chan int)
	quit := make(chan int)

	// 本身goroutine是无阻塞的, 但是通道没有值, 只有有值才会触发Println
	go func() {
		for i := 0; i < 10; i++ {
			Println("c:", <-cx)
		}
		quit <- 0
	}()
	fibonacciV2(cx, quit)

	// 6. 使用WaitGroup
	// sync.WaitGroup用于等待多个 Goroutine完成
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // 增加计数器1
		go worker(i, &wg)
	}
	wg.Wait() // 等待所有 Goroutine 完成
	Println("All workers done")

	// context: 用于控制goroutine 生命周期
	// context.WithCancel, context.WithTimeout

	// Mutex 和 RWMutex:
	// sync.Mutex 提供互斥锁, 用于保护共享资源
	var mu sync.Mutex
	mu.Lock()

	mu.Unlock()

	// 9. go的继承, 是组合形态, 比如父类struct实现，拥有一个接口 stop() , 子类struct 组合了父类, 此时子类{父类对象:父类对象{}}, 后继承了父类的行为

	v := Vehicle{Brand: "Toyata"}
	c2 := Car{
		Vehicle: Vehicle{Brand: "Honda"},
		Model:   "Civic",
	}

	v.Start()          // Toyata Started
	c2.Start()         // Honda Civic car Started
	c2.Vehicle.Start() // Honda Started
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Goroutine 完成时调用 Done()
	//  defer后面的函数只有在当前函数执行完毕后才能执行。
	// 多个defer出现的时候，它会把defer之后的函数压入一个栈中延迟执行，也就是先进后出(LIFO).
	Printf("Worker %d started\n", id)
	Printf("Worker %d finished\n", id)
}

func fibonacciV2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			Println("quit")
			return
		}
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func sayRoutine(cnt int) {
	for i := 0; i < cnt; i++ {
		Println("Hello")
		// 100ms
		time.Sleep(100 * time.Millisecond)
	}
}

// 基类
type Vehicle struct {
	Brand string
}

func (v *Vehicle) Start() {
	Println(v.Brand, "started")
}

// 派生类
type Car struct {
	Vehicle // 嵌入
	Model   string
}

// 重写
func (c *Car) Start() {
	Println(c.Brand, c.Model, "car Started")
}

// defer的参数是声明时传入的
func DeferParams() {
	var age = 10
	//如果想要追踪值类型的变化可以传入值类型指针
	defer func(a *int) {
		Printf("最初如果传入指针,defer内参数为%d\n", *a)
	}(&age)

	defer func(a int) {
		Printf("defer内的参数为%d\n", a)
	}(age)

	age = 25
	Printf("age已经变成了%d\n", age)
}

// age已经变成了25
// defer内的参数为10
// 最初如果传入指针,defer内参数为25
