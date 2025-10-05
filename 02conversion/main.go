package main

func main(){
	// Conversion example
	var num int = 42
	var floatNum float64 = float64(num) // Convert int to float64
	var str string = "123"
	var strToInt int = int(str[0] - '0') // Convert first character of string to int

	// Print the converted values
	println("Converted int to float64:", floatNum)
	println("Converted string to int:", strToInt)
}