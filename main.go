package main

import (
	"fmt"
	"time"

	"github.com/intellect-sam/my-go-project/types"
)

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

}
