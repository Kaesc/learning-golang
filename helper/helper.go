package helper

var version = "1.0.0" // cant be accessed outside this package. have to start with Uppercase 
var Application = "Belajar Golang"

func SayHello(name string) string {
	return "Hello " + name
}