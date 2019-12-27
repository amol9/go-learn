package main
import ("fmt");
import ("io/ioutil");

func main() {
	text, _ := ioutil.ReadFile("test.txt");
	fmt.Printf("%s", text);
}
