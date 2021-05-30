package main
import(
	"fmt"
	"math"
)
/*Define an interface*/
type Shape interface{
	area() float64
}
/* define a circle struct*/
type Circle struct{
	x,y,radius float64
}
/* define a rectangle struct*/
type Rectangle struct{
	width,height float64
}
/* 
	define a method for circle ( implementation of Shape.area())
*/
func (circle Circle) area()float64{
	return math.Pi*circle.radius*circle.radius
}

/*
	Define a method of rectangle ( implementation of Shape.are())
*/
func (rect Rectangle) area() float64{
	return rect.width*rect.height
}

func getArea(shape Shape) float64{
	return shape.area()
}

func main(){
	circle:=Circle{x:0,y:0,radius:5}
	rectangle:=Rectangle{width:10,height:5}
	fmt.Printf("Circle area :-%f\n",getArea(circle))
	fmt.Printf("Rectangle area:-%f\n",getArea(rectangle))
}