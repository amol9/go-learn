package main
import ("fmt")
import ("time")

func test() {
	var i int;
	for i = 0; i < 10; i++ {
		fmt.Println("test hello", i);
		time.Sleep(2 * time.Second);
	}
}

func main() {
	go test();
	var i int;
	for i = 0; i < 10; i++ {
		fmt.Println("hello", i);
		time.Sleep(1 * time.Second);
	}
}
