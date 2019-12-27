package main
import ("fmt");

func test() {
	fmt.Println("test");
}

func test2() {
	fmt.Println("test 2");
}

func test3() {
	fmt.Println("test 3");
}

func main() {
	defer test();
	defer test2();
	defer test3();
	fmt.Println("main");
}
