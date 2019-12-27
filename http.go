package main
import ("fmt");
import ("net/http");
import ("io/ioutil");

func main() {
	res, err := http.Get("https://i.imgur.com/ca5GXkgg.jpg");

	defer res.Body.Close();

	body, err := ioutil.ReadAll(res.Body);

	fmt.Println("error", err);

	ioutil.WriteFile("image.jpg", body, 0644);

	fmt.Println("done");
}

