package main
import "fmt"
type square struct {
	side float64
}
type circle struct {
	radius float64
}
type shape interface {
	area() float64
}
func (s square) area() float64 {
	return s.side * s.side
}
func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}
func returnArea(s shape) float64 {
	return s.area()
}
func main() {
	c := circle{7.28}
	s := square{2.5}
	fmt.Println(returnArea(c))
	fmt.Println(returnArea(s))
}