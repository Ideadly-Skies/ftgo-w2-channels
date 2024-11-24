package main

import "fmt"

func introduce(student string, c chan string){
	// construct result string 
	result := fmt.Sprintf("Hi, my name is %s", student)
	
	// store result to channel 
	c <- result
}


func main(){
	// create a channel of string
	c := make(chan string)
	
	// create three go-routine function
	go introduce("Obie", c)
	go introduce("Helena", c)
	go introduce("Sydney", c)

	// store first greeting to message1
	msg1 := <- c
	fmt.Println(msg1)	
	
	// store second greeting to message2
	msg2 := <- c	
	fmt.Println(msg2)

	// store third greeting to message3
	msg3 := <- c	
	fmt.Println(msg3)

	// close the channel
	close(c)
}