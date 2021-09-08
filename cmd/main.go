package main

import "github.com/EstebanBorai/goras/pkg"

func main() {
	options, createOptionsError := pkg.NewOptions()

	if createOptionsError != nil {
		panic(createOptionsError)
	}

	app := pkg.NewApp(*options)

	if err := app.Start(); err != nil {
		panic(err)
	}
}
