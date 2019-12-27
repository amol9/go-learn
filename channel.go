package main
import ("fmt")
import ("time")

func test(ch chan int) {
	var i int;
	for i = 0; i < 10; i++ {
		fmt.Println("test hello", i);
		time.Sleep(2 * time.Second);
	}
	ch <- 10;
}

func main() {
	ch := make(chan int);
	go test(ch);
	var i int;
	for i = 0; i < 10; i++ {
		fmt.Println("hello", i);
		time.Sleep(1 * time.Second);
	}
	c := <-ch;
	fmt.Println("channel", c);
}
