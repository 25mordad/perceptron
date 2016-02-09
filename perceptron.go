package main
import (
	"fmt"
)
type Point struct {
	X float64
	Y float64
}
type MyData struct {
	p Point
	r bool
}

var trainingSet []MyData
var weight Point

func main() {
	
	fmt.Println("Start....")
	
	threshold := 0.1
	learningRate := 0.1
	weight = Point{0,0}
	trainingSet = append(trainingSet, MyData{Point{0,0},false})
	trainingSet = append(trainingSet, MyData{Point{1,1},true})
	trainingSet = append(trainingSet, MyData{Point{1,0},true})
	trainingSet = append(trainingSet, MyData{Point{0,1},true})
	
	fmt.Println("There are our information:" )
	fmt.Println("Threshold: " ,threshold)
	fmt.Println("learning Rate: " ,learningRate)
	fmt.Println("Primary Weight: " ,weight)
	fmt.Println("Training Set: " ,trainingSet)
	
	for {
		fmt.Println("*********************************************")
		Error := 0
			for _, t := range trainingSet {
				fmt.Println("Weight: " ,weight)
				fmt.Println("Point: " ,t)
				sum:= sum(t)
				fmt.Println("Product is: " ,sum)
				var result bool
				if sum > threshold {
					fmt.Println("Sum is bigger than threshold ... so I put the result as True")
					result = true
				}else{
					fmt.Println("Sum is smaller than threshold ... so I put the result as False")
					result = false
				}
				if result != t.r {
					fmt.Println("Oh No .. We get an Error")
					Error++
				}else{
					fmt.Println("Hey Buddy ... Be Happy .. Without Error")
				}
				
				weight.X = weight.X + learningRate * float64(Error) * t.p.X
				weight.Y = weight.Y + learningRate * float64(Error) * t.p.Y
				
				fmt.Println("----------")
			}
		if Error == 0 {
			break
		}
	}
}

func sum(v MyData) float64{
	return v.p.X * weight.X + v.p.Y * weight.Y
}
