package main

//import (
//	"context"
//	"fmt"
//	"math/rand"
//	"time"
//)

//
//import "fmt"
//
//func main() {
//	ch := make(chan bool, 2)
//	ch <- true
//	ch <- true
//	close(ch)
//
//	//we retrieve the first two values we sent on the channel,
//	//then on our third attempt the channel gives us the values of false and false.
//	//The first false is the zero value for that channel’s type, which is false, as the channel is of type chan bool.
//	//The second indicates the open state of the channel, which is now false
//	for i := 0; i < cap(ch) +1 ; i++ {
//		v, ok := <- ch
//		fmt.Println(v, ok)
//	}
//
//	//for v := range ch {
//	//	fmt.Println(v) // called twice
//	//}
//
//
//	//nil channels blocks forever
//	var ch2 chan bool
//	<- ch2 // blocks forever
//}


//import (
//	"fmt"
//	"time"
//)
//
//var cnt int
//func WaitMany(a, b chan bool) {
//	for a != nil || b != nil {
//		select {
//		case <-a:
//			a = nil
//		case <-b:
//			b = nil
//		}
//		cnt++
//		fmt.Println("From WaitMany ", cnt)
//	}
//}
//
//func main() {
//	a, b := make(chan bool), make(chan bool)
//	t0 := time.Now()
//	cnt = 0
//	go func() {
//		close(a)
//		close(b)
//	}()
//	WaitMany(a, b)
//	fmt.Printf("waited %v for WaitMany\n", time.Since(t0))
//}





//Fanout Pattern example
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//func main() {
//	emps := 20
//	ch := make(chan string, emps)
//
//	for e := 0; e < emps; e++ {
//		go func() {
//			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
//			ch <- "paper"
//		}()
//	}
//
//	for emps > 0 {
//		p := <-ch
//		fmt.Println(p)
//		emps--
//	}
//}
//






//Drop pattern
//import (
//	"fmt"
//)
//
//func main() {
//	const cap = 5
//	ch := make(chan string, cap)
//
//	go func() {
//		for p := range ch {
//			fmt.Println("employee : received :", p)
//		}
//	}()
//
//	const work = 20
//	for w := 0; w < work; w++ {
//		select {
//		case ch <- "paper":
//			fmt.Println("manager : send ack")
//		default:
//			fmt.Println("manager : drop")
//		}
//	}
//
//	close(ch)
//}




//Signal With Data - Delayed Guarantee - Buffered Channel 1
//import (
//	"fmt"
//)
//
//func main() {
//	ch := make(chan string, 1)
//
//	go func() {
//		for p := range ch {
//			fmt.Println("employee : working :", p)
//		}
//	}()
//
//	const work = 10
//	for w := 0; w < work; w++ {
//		fmt.Println("Manager : Sending :")
//		ch <- "paper"
//	}
//	close(ch)
//}



//Signal Without Data - Context
import (
"context"
"fmt"
"math/rand"
"time"
)

func main() {
	duration := 50 * time.Millisecond

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		ch <- "paper"
	}()

	select {
	case p := <-ch:
		fmt.Println("work complete", p)

	case <-ctx.Done():
		fmt.Println("moving on")
	}
}
