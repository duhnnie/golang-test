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

func listUsers(ctx context.Context, completed chan bool) {
	totalUsers := len(users)

	for index, user := range users {
		select {
		case <-ctx.Done():
			completed <- false
			return
		case <-time.After(time.Second):
			fmt.Printf("%d of %d: %+v\n", index+1, totalUsers, user)
		}
	}

	completed <- true
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	completed := make(chan bool)
	defer cancel()

	go listUsers(ctx, completed)

	select {
	case isCompleted := <-completed:
		if isCompleted {
			fmt.Println("User listing is complete")
		} else {
			fmt.Println("User listing was interrupted due timeout")
		}
	}
}
