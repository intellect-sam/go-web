package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/intellect-sam/my-go-project/types"
)

var wg sync.WaitGroup

func main() {

	// x, y := 10, 20

	// fmt.Println("x + y =", calc.Add(x, y))
	// fmt.Println("x - y =", calc.Subtract(x, y))

	// students := []string{"Sam", "John", "Peter", "Alice", "Bob"}
	// newStudents := append(students, "Tom", "Jerry")

	// for _, v := range newStudents {
	// 	fmt.Println(v)
	// }

	nameOfStudents := make([]string, 3)
	nameOfStudents[0] = "Sam"
	nameOfStudents[1] = "John"
	nameOfStudents[2] = "Peter"

	person := types.Person{
		FirstName: "Sam",
		LastName:  "Smith",
		Dob:       time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		Email:     "test@gmail.com",
		Location:  "New York",
	}

	// fmt.Println(person.FullName())

	sam := types.Staff{
		Person: person, Roles: []string{"Admin", "CTO"},
	}

	for _, role := range sam.Roles {
		fmt.Println(role)
		if role == "Admin" || role == "Editor" {
			fmt.Println("You have access to the system")
		}
	}

	fmt.Println(sam.Roles)
	fmt.Println(sam.Person)

	// This is where the gorutines are used
	count := make(chan int)
	wg.Add(2)

	fmt.Println("Start Gorutines")
	go PrintCount("A", count)
	go PrintCount("B", count)

	fmt.Println("Channel begins")
	count <- 1

	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Println("\nTerminating the program")

}

func PrintCount(label string, count chan int) {
	defer wg.Done()

	for {
		val, ok := <-count
		if !ok {
			fmt.Println("Channel was closed")
			return
		}
		fmt.Printf("Count: %d received from %s \n", val, label)
		if val == 10 {
			fmt.Printf("Channel Closed from %s \n", label)
			// Close the channel
			close(count)
			return
		}
		val++
		// Send count back to the other goroutine.
		count <- val
	}

}

func channel() {
	count := make(chan string, 1)
	count <- "Hello Samuel"

	message := <-count
	fmt.Println(message)
}
