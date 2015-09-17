package main

import ("fmt"; "math")

type Shape interface {
  perimeter() float64
}

type Circle struct {
  r float64
}

type Rectangle struct {
  x, y float64
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) perimeter() float64 {
  return 2*r.x + 2*r.y
}

func totalPerimeter(shapes ...Shape) float64 {
  var p float64
  for _, s := range shapes {
    p += s.perimeter()
  }
  return p
}

func main() {
	var rad, l float64
	fmt.Print("Enter radius of Circle:")
	fmt.Scanf("%f",&rad)
	c := Circle{rad}
	r := Rectangle{10,20}
	fmt.Print("Total perimeter of Circle and rectangle is ")
	fmt.Println(totalPerimeter(&c, &r))
}