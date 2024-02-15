package main

import "fmt"

func main() {

	// membuat nama tipe data
	type NoKTP string

	var noKtpAl NoKTP = "12121212121212100"
	fmt.Println(noKtpAl)
}
