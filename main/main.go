/*
 Author: Kernel.Huang
 Mail: kernelman79@gmail.com
 File: main
 Date: 8/16/22 2:18 AM
*/
package main

import (
	"github.com/jucci1887/irisapi"
	"github.com/jucci1887/irisapi/router"
)

// Iris api call example.
func main() {
	app := irisapi.App.New()
	router.AppRoute.NewRouter(app.Apply)
	app.Run(app.Apply)
}
