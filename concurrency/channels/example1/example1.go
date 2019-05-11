// This sample programs demonstrate the basic channel machanis for goroutine
// signaling.

package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// waitForResult()
	// waitForTask()
	// fanOut()
	// pooling()
	// fanOutSem()
	// fanOutBounded()
	// drop()
	cancellation()
}

// waitForResult: Yor are a manager and you hire a new employee. Your new employee
// knows immediately what they are expecting to do and start their work.
// You sit waiting for the result of the employee's work. The amount of time you wait on the employee is unknown because you need a
// guarantee that the result sent by the employee is received by you.
func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "paper"
		fmt.Println("employee : sent signal")
	}()

	p := <-ch
	fmt.Println("Manager : received signal: ", p)
	time.Sleep(time.Second) // Wait for goroutine to complete.
}

// waitForTask: You are a manager and you hire a new employee. Your new employee
// doesn't know immediately what they are expected to do and waits for you to tell
// them what to do. You prepare the work and sent it to them. The amount of time they wait is
// unknown because you need a guarantee that the work your sending is received
// by the employee.
func waitForTask() {
	ch := make(chan string)
	go func() {
		p := <-ch
		fmt.Println("employee: received signal: ", p)
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	ch <- "paper"
	fmt.Println("manager: sent signal")
	time.Sleep(time.Second)
}

// fanOut: You are a manager and you hire one new employee for the exact amount of
// work you have to get done. Each new employee knows immediately what they are
// expected to do and start their work. You sit waiting for all the results of the
// employee's work. The amount of time you wait on the employees is unknown because you need a
// guarantee that all the results sent by employees are received by you. No given
// employees needs an immediate guarantee that you received their result.

func fanOut() {
	emps := 2000
	ch := make(chan string, emps)

	for i := 0; i < emps; i++ {
		go func(emp int) {
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			ch <- "paper"
			fmt.Println("employee: send signal: ", emp)
		}(i)
	}

	for emps > 0 {
		p := <-ch
		emps--
		fmt.Println(p)
		fmt.Println("manager: remaining employees: ", emps)
	}
}

// pooling: You are a manager and you hire a team of employees. None of the new
// employees know what they are expecting to do and wait for you to provide work.
// When work is provided to the group, any given employee can take it and you
// don't care who it is. The amount of time you wait for any given employee to take your work is unknown
// because you need a guarantee that the work your sending is received by an
// employee.
func pooling() {
	ch := make(chan string)

	g := runtime.NumCPU()
	for i := 0; i < g; i++ {
		go func(emp int) {
			for p := range ch {
				fmt.Printf("employee %d: received signal: %s\n", emp, p)
			}
			fmt.Printf("employee %d: received shutdown signal\n", emp)
		}(i)
	}

	const work = 100
	for w := 0; w < work; w++ {
		ch <- "paper"
		fmt.Println("manager: sent signal: ", w)
	}
	close(ch)
	fmt.Println("manager: sent shutdown signal")
	time.Sleep(time.Second)
}

// fanOutSem: You are a manager and you hire one new employee for the extact amount
// of work you have to get done. Each new employee knows immediately what they
// are expected to do and starts their work.
// However, you don't want all the employees working at once. You want to limit
// how many of them are working at any given time. You sit waiting for all the
// result of the employees work. The amount of time you wait on the employees is
// unknown because you need a guarantee that all the results sent by employees are
// received by you. No given employee needs an immediate guarantee that you
// received their result
func fanOutSem() {
	emps := 2000
	ch := make(chan string, emps)

	g := runtime.NumCPU()
	sem := make(chan bool, g)

	for i := 0; i < emps; i++ {
		go func(emp int) {
			sem <- true
			{
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
				ch <- "paper"
				fmt.Println("employee: sent signal: ", emp)
			}
			<-sem
		}(i)
	}

	for emps > 0 {
		p := <-ch
		emps--
		fmt.Println(p)
		fmt.Println("manager: remaining employees: ", emps)
	}
}

// fanOutBounded: You are a manager and you hire a team of employees. None of the
// new employees know what they are expected to do and wait for you to provide work.
// The amount of work that needs to get done is fixed and staged ahead of time.
// Any given employee can take work and you don't care who it is or what they take.
// The amount of time you wait on the employees to finish all the work is unknown
// because you need a guarantee that all the work is finished.
func fanOutBounded() {
	works := []string{"paper", "paper", "paper", "paper", "paper", 2000: "paper"}

	g := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan string, g)

	for e := 0; e < g; e++ {
		go func(id int) {
			defer wg.Done()
			for p := range ch {
				fmt.Printf("employee %d: received signal: %s\n", id, p)
			}
			fmt.Printf("employee %d: received shutdown\n", id)
		}(e)
	}

	for _, w := range works {
		ch <- w
	}
	close(ch)
	wg.Wait()
}

// drop: You are a manager and you hire a new employee. Your new employee doesn't
// know immediately what they are expected to do and waits for you to tell them
// what to do. You prepare the work and send it to them. The amount of time they
// wait is unknown because you need a guarantee that the work sending is received
// by the employee. You won't wait for the employee to take the work if they are
// not ready to receive it. In that case you drop the work on the floor and try
// try again with the next piece of work.
func drop() {
	const cap = 100
	const work = 200
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("employee: received signal: ", p)
		}
	}()

	for w := 0; w < 200; w++ {
		select {
		case ch <- "paper":
			fmt.Println("manager: sent signal: ", w)
		default:
			fmt.Println("manager: dropped data: ", w)
		}
	}
	close(ch)
	time.Sleep(time.Second)
}

// cancellation: Yor are a manager and you hire a new employee. Your new employee
// knows immediately what they are expected to do and start their work. You sit
// waiting for the result of the employee's work. The amount of time you wait on
// the employee is unknown because you need a guarantee that the result sent by
// the employee is received by you. Expect you are not willing to wait forever
// for the employee to finish their work. They have a specified amount of time
// and if they are not done, you don't wait and walk away.
func cancellation() {
	duration := 150 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		ch <- "paper"
	}()

	select {
	case d := <-ch:
		fmt.Println("work complete", d)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
}
