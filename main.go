package main

func main() {
	configuration := configuration{}
	configuration.Load()

	router := router{}
	router.Init()
	router.Run(&configuration)
}
