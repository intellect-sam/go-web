package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/intellect-sam/my-go-project/types"
	"golang.org/x/exp/rand"
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
	wg.Add(2)

	fmt.Println("Start Goroutines")
	go PrintCount("A")
	go PrintCount("B")

	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")

}

func PrintCount(label string) {
	defer wg.Done()

	for count := 1; count <= 10; count++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Count: %d from %s\n", count, label)
	}
}
