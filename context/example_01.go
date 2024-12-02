package main

import (
	"context"
	"fmt"
	"time"
)

type User struct {
	Name  string
	Age   int
	Email string
}

var users []*User = []*User{
	{"Daniel", 34, "me@daniel.com"},
	{"Monica", 53, "monica@gmail.com"},
	{"Ramona", 8, "ramonella@gmail.com"},
	{"Andrea", 32, "andrea@asdfasdf.com"},
	{"Roberto", 40, "robocop@ocp.com"},
	{"Sergio", 30, "quinho@processmaker.com"},
}

func listUsers(ctx context.Context) {
	totalUsers := len(users)

	for index, user := range users {
		select {
		case <-ctx.Done():
			fmt.Println("Context is done!")
			return
		default:
			fmt.Printf("%d of %d: %+v\n", index, totalUsers, user)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	listUsers(ctx)
}
