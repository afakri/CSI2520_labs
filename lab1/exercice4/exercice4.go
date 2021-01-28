package main

import "fmt"

// Define a suitable structure
func main() {
	// Create a dynamic map
	m := make(map[string]Course)
	// Add the courses CSI2120 and CSI2110 to the map
	course1 := Course{186, "Lang", 79.500000}
	course2 := Course{211, "Moura", 81.000000}
	m["CSI2110"] = course1
	m["CSI2120"] = course2
	for k, v := range m {
		fmt.Printf("Course Code: %s\n", k)
		fmt.Printf("Number of students: %d\n", v.NStudents)
		fmt.Printf("Professor: %s\n", v.Professor)
		fmt.Printf("Average: %f\n\n", v.Avg)
	}
}

//Course is...
type Course struct {
	NStudents int
	Professor string
	Avg       float64
}
