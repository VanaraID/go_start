package main

import (
	"fmt" // formatted

	"github.com/VanaraID/go_start/fundamentals"
	"github.com/VanaraID/go_start/joglo"
)

func main() {
	word, nomor := fundamentals.HelloWorld()

	alamat := joglo.GetAddress()

	fmt.Println(word)
	fmt.Println("Alamat: ", alamat, nomor)
}
