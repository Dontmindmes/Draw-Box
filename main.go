package main

import "fmt"

const floorD string = "─"
const sidesD string = "|"

func main() {

	var (
		floor int
		sides int
	)

	fmt.Println("Enter the dimensions of the box")

	fmt.Printf("Width: ")
	fmt.Scan(&floor)

	fmt.Printf("Height: ")
	fmt.Scan(&sides)

	Draw(floor, sides)
}

func Draw(W, H int) {

	Minus := H - 2

	for top := 0; top < W; top++ {
		fmt.Print(floorD)
	}
	fmt.Println()

	for Side := 0; Side < H; Side++ {
		fmt.Print(sidesD)
		for Space := 0; Space < Minus; Space++ {
			fmt.Print(" ")
		}
		fmt.Println(sidesD)
	}
	fmt.Print()

	for Bottem := 0; Bottem < W; Bottem++ {
		fmt.Print(floorD)
	}

	fmt.Println()

}
