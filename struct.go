package main
import ("fmt");

type Vehicle struct {
	company	string;
	model	string;
}

func (v Vehicle) display() {
	fmt.Println(v.company, v.model);
}

func main() {
	var v Vehicle;
	v.company = "Hyundai";
	v.model = "Xing";

	v.display();

	c := Vehicle{"Maruti", "Zen"};
	c.display();
}
