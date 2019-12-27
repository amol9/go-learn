package main
import ("fmt")
import ("time")

func test(ch chan int) {
	var i int;
	for i = 0; i < 10; i++ {
		fmt.Println("test hello", i);
		time.Sleep(1 * time.Second);
	}
	ch <- 10;
}

func test2(ch chan int) {
	var i int;
	for i = 0; i < 10; i++ {
		fmt.Println("test2 hello", i);
		time.Sleep(2 * time.Second);
	}
	ch <- 20;
}

func main() {
	ch := make(chan int);
	ch2 := make(chan int);
	go test(ch);
	go test2(ch2);

	select {
	case x := <- ch:
		fmt.Println("test done", x);
	case y := <- ch2:
		fmt.Println("test2 done", y);
	//default:
	//	fmt.Println("default");
	}
}
