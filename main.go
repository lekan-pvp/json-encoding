package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	id       int    `json:"id"`
	name     string `json:"name,omitempty"`
	email    string `json:"email,omitempty"`
	password string `json:"-"`
}

func main() {
	users := []User{
		{
			id:       1,
			name:     "Gopher",
			email:    "gopher@example.com",
			password: "Im4G0pH3r",
		},
		{
			id:       2,
			name:     "Rustocaen",
			email:    "rustocean@example.com",
			password: "iT$Ru$t0C34n",
		},
	}

	out, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		fmt.Printf("serialization error: %s\n", err.Error())
		return
	}

	fmt.Println(string(out))
}
