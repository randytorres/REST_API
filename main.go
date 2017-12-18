package main

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "Doskids23"
	dbname   = "test"
)

func main() {
	a := App{}
	a.Initialize(user, password, dbname)

	a.Run(":8080")
}
