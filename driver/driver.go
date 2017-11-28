package main

import (
	"fmt"

	"github.com/vishwanathj/golangexercises/rectangle"

)

func main() {
	r1:= rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{3,3,}}
	r2:= rectangle.Rectangle{rectangle.Point{0,0}, rectangle.Point{1,1,}}
	fmt.Println("The intersecting rectangle is: %v", rectangle.GetTheIntersectingRectangle(r1, r2))
}
