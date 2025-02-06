package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Books struct {
	Name string `json:"name"`
}

var books = []Books{
	{Name: "Shashila Heshan"},
}

func getBooks(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	json.NewEncoder(res).Encode(books)
}

// func main() {
// 	sliceTest()

// 	repo := NewUserRepos("heshan")

// 	repo.GetUser()

// 	//port := ":9090"

// 	// http.HandleFunc("/books", getBooks)
// 	// log.Printf("Server has started on port %s", port)
// 	// err := http.ListenAndServe(port, nil)

// 	// if err != nil {
// 	// 	log.Fatalf("Error connecting to the server %s", err.Error())
// 	// }
// }

func sliceTest() {
	slice1 := make([]int, 3)
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3

	slice := []string{
		"heshan", "shashila", "udes", "kanthi",
	}

	//p := slice1[0:3]

	//t := "test"

	book := Book{
		"hesh",
	}

	bookAddress := &book.Test

	fmt.Println(cap(slice))

	fmt.Println(*bookAddress)
	changeTestPnt(&book)
	fmt.Print(book.Test)

}

type Book struct {
	Test string `json:"Test"`
}

func changeTest(book Book) {
	book.Test = "hi"
}
func changeTestPnt(book *Book) {
	book.Test = "hello"
}

// UserRepo interface with GetUser method
type UserRepo interface {
	GetUser()
}

// UserRepos struct implementing UserRepo
type UserRepos struct {
	Client string
}

// Constructor function returning a UserRepo instance
func NewUserRepos(client string) UserRepo {
	return &UserRepos{Client: client} // Returning pointer to UserRepos
}

// Method to implement GetUser()
func (u *UserRepos) GetUser() {
	fmt.Println("Client:", u.Client)
}

func main() {
	// Correctly calling NewUserRepos and invoking GetUser()
	repo := NewUserRepos("Heshan")
	repo.GetUser() // Output: Client: Heshan

	UsageofLoops()
	printTri()
}

func UsageofLoops() {

	slice := make([]string, 3)

	slice[0] = "test"
	slice[1] = "test2"
	slice[2] = "test3"

	for i, item := range slice {
		fmt.Println("Item is ", item, i)
	}
}
func printTri() {
	for i := 1; i <= 5; i++ { // Number of rows
		for j := 1; j <= i; j++ { // Print '*' i times
			fmt.Print("*") // Print '*' on the same line
		}
		fmt.Println() // Move to the next line
	}
}

// *
// **
// ***
// ****
// *****
