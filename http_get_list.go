package main
import ("fmt");
import ("net/http");
import ("io/ioutil");
import ("strings");
import ("flag");
import ("strconv");
import ("time");

func read_file(filename string)(string) {
	text, _ := ioutil.ReadFile(filename);
	return string(text);
}

func split_text(text string)([]string) {
	return strings.Split(text, "\n");
}

func parse_cmdline()(string) {
	filename := flag.String("filename", "urls.txt", "filename");
	flag.Parse();
	return *filename;
}

func url_get(url string)([]byte) {
	res, _ := http.Get(url);
	defer res.Body.Close();
	body, _ := ioutil.ReadAll(res.Body);
	return body;
}

func save_bytes(filename string, data []byte) {
	ioutil.WriteFile(filename, data, 0644);
}

func main() {
	filename := parse_cmdline();
	text := read_file(filename);
	urls := split_text(text);

	//var i int;
	//var url string;
	for i, url := range urls {
		bytes := url_get(url);
		var s string = strconv.Itoa(i);
		fmt.Println(s);
		save_bytes("test/" + s + ".jpg", bytes);
		time.Sleep(2 * time.Second); //trying to fix the panic crash
	}

	fmt.Println("done", filename);

}		
