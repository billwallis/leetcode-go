package samples_test

import (
	"github.com/Bilbottom/leetcode-go/internal/samples"
	"testing"
	"time"
)

func TestUser_Forename(t *testing.T) {
	user := samples.User{}
	want := "John"
	user.SetForename(want)
	got := user.Forename()

	if want != got {
		t.Fatalf("User.Forename() = %s; want %s", got, want)
	}
}

func TestUser_Surname(t *testing.T) {
	user := samples.User{}
	want := "Smith"
	user.SetSurname(want)
	got := user.Surname()

	if want != got {
		t.Fatalf("User.Surname() = %s; want %s", got, want)
	}
}

func TestUser_BirthDate(t *testing.T) {
	user := samples.User{}
	date := "2000-01-01"
	want, _ := time.Parse("2006-01-02", date)
	user.SetBirthDate(date)
	got := user.BirthDate()

	if want != got {
		t.Fatalf("User.BirthDate() = %s; want %s", got, want)
	}
}

func TestUser_Email(t *testing.T) {
	user := samples.User{}
	user.SetForename("John")
	user.SetSurname("Smith")
	want := "John.Smith@example.com"
	got := user.Email()

	if want != got {
		t.Fatalf("User.Email() = %s; want %s", got, want)
	}
}
