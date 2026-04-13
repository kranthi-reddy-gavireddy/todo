package main

import "github.com/kranthi-reddy-gavireddy/internal/api/app"

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
