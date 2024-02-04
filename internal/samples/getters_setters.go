package samples

/*
Getters and setters aren't idiomatic in Go; instead, you should use
exported fields and methods.

Some examples are below just for reference anyway.
*/

import (
	"time"
)

type User struct {
	forename  string
	surname   string
	birthDate time.Time
}

func (user *User) SetForename(name string) {
	user.forename = name
}

func (user *User) Forename() string {
	return user.forename
}

func (user *User) SetSurname(name string) {
	user.surname = name
}

func (user *User) Surname() string {
	return user.surname
}

func (user *User) SetBirthDate(date string) {
	dateTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}
	user.birthDate = dateTime
}

func (user *User) BirthDate() time.Time {
	return user.birthDate
}

func (user *User) Email() string {
	return user.forename + "." + user.surname + "@example.com"
}
