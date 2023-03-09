package main

import "fmt"

func main() {

/*
Assignment 1:
menampilkan nilai i : 21 ---> fmt.Printf("%T \n", i) // contoh : fmt.Printf("%v \n", i)
menampilkan tipe data dari variabel i 
menampilkan tanda %
menampilkan nilai boolean j : true

menampilkan nilai base2 dari i : 10101
menampilkan unicode karakter : Я
menampilkan nilai base 10 dari i : 21
menampilkan nilai base 8 dari i :25
menampilkan nilai base 16 : f --> base 16 dari f itu angka berapa?
menampilkan nilai base 16 : F --> base 16 dari f itu angka berapa?
menampilkan unicode karakter Я (bahasa rusia "ya") : U+042F

var k float64 = 123.456;
menampilkan float dari k : 123.456000
menampilkan float scientific dari k: 1.234560E+02
*/
	i := 21
	j := true
	var k float64 = 123.456

	fmt.Printf("%v \n", i)
	fmt.Printf("%T \n", i)
	fmt.Printf("%% \n")
	fmt.Printf("%t \n", j)
	
	fmt.Printf("\n")

	fmt.Printf("%b \n", i)
	fmt.Printf("%c \n", 'Я')
	fmt.Printf("%d \n", i)
	fmt.Printf("%o \n", i)
	fmt.Printf("%x \n", 15)
	fmt.Printf("%X \n", 15)
	fmt.Printf("%U \n", 'Я')

	fmt.Printf("\n")

	fmt.Printf("%f \n", k)
	fmt.Printf("%e \n", k)
}

/*
Cheat Sheet
https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/#character
*/