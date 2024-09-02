package main

import "fmt"

func main() {
	var limit int
	fmt.Print("Masukkan batas atas untuk mencari bilangan prima: ")
	fmt.Scan(&limit)
	fmt.Printf("Bilangan prima antara 1 hingga %d adalah: \n", limit)

	for n := 2; n <= limit; n++ {
		isPrime := true
		if n > 1 {
			for i := 2; i*i <= n; i++ {
				if n%i == 0 {
					isPrime = false
					break
				}
			}

			if isPrime {
				fmt.Println(n)
			}
		}
	}
}
