package main
import ("fmt");

func main() {
	var i int;
	for i = 0; i < 4; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "even");
		} else {
			fmt.Println(i, "odd");
		}
	}
}
