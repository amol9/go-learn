package main
import ("fmt");
import ("io/ioutil");
import ("strings");

func main() {
	text, _ := ioutil.ReadFile("test.txt");
	list := strings.Split(string(text), "\n");

	var i int;
	for i = 0; i < len(list); i++ {
		fmt.Println(list[i]);
	}
}
