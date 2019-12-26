package main
import ("fmt");

func main() {
	var i int;
	for i = 0; i < 10; i++ {
		switch i % 4 {
			case 0:
				fmt.Println("perfectly divisible by 4");
			case 1:
				fmt.Println("not divisible by 4, remainder 1");
			case 2:
				fmt.Println("not divisible by 4, remainder 2");
			case 3:
				fmt.Println("not divisible by 4, remainder 3");
			default:
				fmt.Println("how is it possible?");
		}
	}
}
