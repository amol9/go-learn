package main
import ("fmt");

func test() {
	fmt.Println("test");
}

func main() {
	defer test();
	fmt.Println("main");
}
