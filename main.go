package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8088"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Home Page")
}
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	fmt.Fprintf(w, fmt.Sprintf("This is about page and 2+32 is %d", sum))
}
func addValues(x, y int) int {
	return x + y
}
func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 10.0, 0.0, f))
}
func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by 0")
		return 0, err
	}
	result := x / y
	return result, nil
}
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
