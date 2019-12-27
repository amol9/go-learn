package main
import ("fmt");

func main() {
	var a int = 100;
	var p *int = &a;

	fmt.Println("value:", a);
	fmt.Println("address:", p);

	fmt.Println("dereference:", *p);

	*p = *p + 1;
	fmt.Println("deref and inc, value:", a);
	fmt.Println("deref and inc, deref:", *p);

}

