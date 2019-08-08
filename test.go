package main

import "fmt"

func main()  {
	certid := []byte("425566090740792125343006613097721207771083416073")
	newCertId := []byte{}
	newCertId = certid
	fmt.Println(string(newCertId))
}