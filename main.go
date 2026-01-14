package main

import (
	"fmt"
	"study_3/feature1"
	"study_3/feature2"
	"study_3/feature_postgres/simple_connection"
)



func main()  {
	fmt.Println("Hello, Git!")
	feature1.Feature1()
	feature2.Feature2()
	simple_connection.CheckConnection()
}