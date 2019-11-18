package main

func main() {
	fmt.Println("conting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Prinln("done")
}
