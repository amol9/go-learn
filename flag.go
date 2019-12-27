package main
import ("fmt");
import ("flag");

func main() {
	ip := flag.Int("test", 100, "test int value");
	flag.Parse();
	fmt.Println("value:", *ip);
}
