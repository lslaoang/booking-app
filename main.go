package main

import "fmt"

func main() {
	const conferenceTickets = 50

	res := author{
        name:      "Sona",
        branch:    "CSE",
        particles: 203,
        salary:    34000,
    }
 
    // Calling the method
    res.show()
	printThisLine()

}

func printThisLine() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to %s booking application! \n", conferenceName)
	fmt.Printf("We  have total of %d tickets and %d are still available. \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

type author struct {
	name      string
	branch    string
	particles int
	salary    int
}

func (a author) show() {
 
    fmt.Println("Author's Name: ", a.name)
    fmt.Println("Branch Name: ", a.branch)
    fmt.Println("Published articles: ", a.particles)
    fmt.Println("Salary: ", a.salary)
}
