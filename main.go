// Author: Joshua Adeyemi
// Version: 1.0.0
// Date: 2025-12-14
// FileOverview: Calculates the average of student marks and gives performance feedback.

package main

import "fmt"

func main() {
	numMarks := 3
	marks := []float64{75.8, 90.2, 80}

	sum := 0.0
	for i := 0; i < numMarks; i++ {
		sum += marks[i]
	}

	average := sum / float64(numMarks)

	fmt.Printf("You have entered %d marks. The student's average is %.0f%%.\n", numMarks, average)

	if average <= 49 {
		fmt.Println("The student is failing.")
	} else if average <= 69 {
		fmt.Println("The student's performance is just under average.")
	} else if average <= 79 {
		fmt.Println("The student's performance is average.")
	} else {
		fmt.Println("The student is on the honour roll.")
	}

	fmt.Println("\nDone.")
}
