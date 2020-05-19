package main

import (
	"html/template" // "text/template"
	"os"
)

// IsOld is age > 20
func (u User) IsOld() bool {
	return u.Age > 20
}

// User is user
type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	user := User{Name: "tucker", Email: "tucker@naver.com", Age: 23}
	user2 := User{Name: "jason", Email: "jason@naver.com", Age: 18}
	users := []User{user, user2}
	tmpl, err := template.New("Tmpl1").ParseFiles("templates/tmpl1.tmpl", "templates/tmpl2.tmpl")
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", users)
}
