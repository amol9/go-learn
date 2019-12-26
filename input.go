package main
import ("fmt");
import ("bufio");
import ("os");

func input(prompt string)(string) {
	reader := bufio.NewReader(os.Stdin);
	fmt.Print(prompt);
	text, _ := reader.ReadString('\n');
	return text;
	//return "test";
}

func main() {
	name := input("enter name: ");
	fmt.Println("hello,", name);
}

