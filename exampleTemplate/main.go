package main

import (
	"html/template"
	"os"
)

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
