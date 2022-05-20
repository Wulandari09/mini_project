package main


// func init() {
// 	config.InitialMigration()
// }
//
func main() {
	e := routers.Router()
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}