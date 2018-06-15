package main

func main() {
	configuration := configuration{}
	configuration.Load()

	database := database{}
	database.Connect(&configuration)

	router := router{}
	router.Init()
	router.Run(&configuration)
}
