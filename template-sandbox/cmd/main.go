package main

import (
	_ "embed"
	"fmt"
	"os"

	templatex "github.com/qawatake/go-sandbox/template-sandbox"
)

type Member struct {
	Name string
	Age  int
}

func main() {
	members := []struct {
		Name string
		Age  int
	}{
		{
			Name: "qawatake",
			Age:  100,
		},
		{
			Name: "bob",
			Age:  20,
		},
	}

	for _, member := range members {
		if err := templatex.NameAge.Execute(os.Stdout, member); err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("###############################")
	for _, member := range members {
		if err := templatex.AgeName.Execute(os.Stdout, member); err != nil {
			fmt.Println(err)
		}
	}
}
