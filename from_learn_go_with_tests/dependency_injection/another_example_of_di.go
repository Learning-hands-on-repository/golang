package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, option IOption) {
	// this is 1000line of codes
	fmt.Fprintf(writer, "Hello from 70 years old code, %s", option.GetName()) // writer also the function from HTTP's author writer
	// and this is another 500 line of codes
}

type IOption interface {
	GetName() string
}

// This is how we define class & method in Golang
type LegacyOption struct {
	specialEndingWord string
}

func (legacyOption LegacyOption) GetName() string {
	return fmt.Sprintf("Nara, the old Jedis %s", legacyOption.specialEndingWord)
}

func (legacyOption LegacyOption) GetLastName() string {
	return fmt.Sprintf("Yo, this is the lastname %s", legacyOption.specialEndingWord)
}

// Now, I'm gonna do modern function
type ModernOption struct{}

func (modernOption ModernOption) GetName() string {
	return "Chawit, the new Jedis"
}

func LegacyGreeter(w http.ResponseWriter, r *http.Request) {
	// NOW! I'm gonna avoid using legacy code
	//Greet(w, LegacyOption{specialEndingWord: "HAPPY ENDING"}) // w is the function from HTTP's author writer

	// It's time for modern one

	Greet(w, ModernOption{}) // w is the function from HTTP's author writer
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(LegacyGreeter)))
}
