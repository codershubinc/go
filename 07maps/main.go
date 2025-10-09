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
	fmt.Println("Length of map is ", len(langs))

	//comma ok idiom
	lang, ok := langs["js"]
	if ok {
		fmt.Println("Language: js ", lang)
	} else {
		fmt.Println("Language js is not present")
	}

	
}
