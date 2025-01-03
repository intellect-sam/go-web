package types

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName, LastName string
	Dob                 time.Time
	Email, Location     string
}

type Staff struct {
	Person
	Roles []string
}

func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

func (p *Person) ChangeLocation(newLocation string) {
	p.Location = newLocation
}
