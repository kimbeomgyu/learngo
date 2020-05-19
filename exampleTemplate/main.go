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
	temp1, err := template.New("Tmp1").ParseFiles("templates/tmpl1.tmpl")
	if err != nil {
		panic(err)
	}
	temp1.ExecuteTemplate(os.Stdout, "tmpl1.tmpl", user)
	temp1.ExecuteTemplate(os.Stdout, "tmpl1.tmpl", user2)
}
