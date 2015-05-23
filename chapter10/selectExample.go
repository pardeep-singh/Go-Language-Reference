package main
import(
	"fmt"
	"time"
)

func c1Sender(c chan string){
	for{
		c <- "from 1"
		time.Sleep(time.Second*2)
	}
}

func c2Sender(c chan string){
	for{
		c <- "from 2"
		time.Sleep(time.Second*3)
	}
}

func receiver(c1 ,c2 chan string){
	for{
		select{
			case msg1 := <- c1:
				fmt.Println(msg1)
			case msg2 := <- c2:
				fmt.Println(msg2)
			case <- time.After(time.Second):
				fmt.Println("timeout")
		}
	}
}
func main(){
	c1 := make(chan string)
	c2 := make(chan string)
	
	go c1Sender(c1)
	go c2Sender(c2)
	go receiver(c1,c2)
	
	var input string
	fmt.Scanln(&input)
	
}