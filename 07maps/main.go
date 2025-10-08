package main

import "fmt"

func main() {
	langs := make(map[string]string)
	langs["go"] = "Golang"
	langs["js"] = "JavaScript"
	langs["py"] = "Python"
	langs["rb"] = "Ruby"
	langs["rs"] = "Rust"
	langs["kt"] = "Kotlin"

	fmt.Println("Map is ", langs)
	fmt.Println("Language: go ", langs["go"])
	fmt.Println("Language: js ", langs["js"])
	fmt.Println("Language: py ", langs["py"])

	delete(langs, "js")
	fmt.Println("Map is ", langs)

	//loops
	for key, value := range langs {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
