package main

import "fmt"

func main() {
	fmt.Println("Netflix Series Watch Time Calculator")
	fmt.Println("===================================")

	var numEpisodes int
	var episodeDuration float64

	// Input the number of episodes watched
	fmt.Print("Enter the number of episodes you watched: ")
	fmt.Scan(&numEpisodes)

	// Input the average duration of each episode in minutes
	fmt.Print("Enter the average duration of each episode (in minutes): ")
	fmt.Scan(&episodeDuration)

	// Calculate total time in minutes and hours
	totalMinutes := float64(numEpisodes) * episodeDuration
	totalHours := totalMinutes / 60

	// Display the result
	fmt.Printf("\nYou watched %d episodes, each lasting %.2f minutes.\n", numEpisodes, episodeDuration)
	fmt.Printf("Total time spent watching: %.2f minutes or %.2f hours.\n", totalMinutes, totalHours)
}
