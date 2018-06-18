package main

func main() {
	SetupProject()

	myRouter.Run()
}

func SetupProject() {
	myConfiguration = &configuration{
		databaseConfiguration: &databaseConfiguration{},
	}
	myDatabase = &database{}
	myRouter = &router{}
	myHelper = &helper{}

	myConfiguration.Load()
	myDatabase.Connect()
	myDatabase.Setup()
	myRouter.Init()
}
