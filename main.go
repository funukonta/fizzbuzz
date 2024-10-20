package main

import "fmt"

func main() {
	/* Buat fungsi yang mencetak angka dari 1 hingga n.
	Tetapi
	- untuk setiap angka yang kelipatan 3, cetak "Fizz"
	- untuk kelipatan 5, cetak "Buzz".
	- angka kelipatan 3 dan 5, cetak "FizzBuzz".
	*/

	n := 20

	fizzbuzz(n)

	fmt.Println("program selesai")
}

// ini func fizzbuzz
func fizzbuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		} else if i%5 == 0 {
			fmt.Println("buzz")
			continue
		} else if i%3 == 0 {
			fmt.Println("fizz")
			continue
		}
		fmt.Println(i)
	}
}
