package struct1

import "fmt"

type book struct {
	title  string
	author string
	year   int
}

func BookImplementation() {
	book1 := book{
		title:  "The Devine Comedy",
		author: "Dante Aligheri",
		year:   1320,
	}

	fmt.Println(book1)

	// two struct values are equal if their corresponding fields are equal.
	randomBook := book{title: "Random Title", author: "John Doe", year: 100}
	bookTest := book{title: "Random Title", author: "John Doe", year: 100}
	fmt.Println(randomBook == bookTest)

	// = creates a copy of a struct
	myBook := randomBook
	myBook.year = 2020              // modifying only myBook
	fmt.Println(myBook, randomBook) // => {Random Title John Doe 2020} {Random Title John Doe 100}
}

// define a new struct type
type Contact struct {
	email, address string
	phone          int
}

// define a struct type that contains another struct as a field
type Employee struct {
	name        string
	salary      int
	contactInfo Contact
}

func EmbeddedStruct() {

	// declaring a value of type Employee
	john := Employee{
		name:   "John Keller",
		salary: 3000,
		contactInfo: Contact{
			email:   "jkeller@company.com",
			address: "Street 20, London",
			phone:   042324234,
		},
	}

	fmt.Printf("%+v\n", john)
	// => {name:John Keller salary:3000 contactInfo:{email:jkeller@company.com address:Street 20, London phone:295619381404}}

	// accessing a field
	fmt.Printf("Employee's salary: %d\n", john.salary)

	// accessing a field from the embedded struct
	fmt.Printf("Employee's email:%s\n", john.contactInfo.email) // => Employee's email:jkeller@company.com

	// updating a field
	john.contactInfo.email = "new_email@thecompany.com"
	fmt.Printf("Employee's new email address:%s\n", john.contactInfo.email)
	// =>  Employee's new email address:new_email@thecompany.com
}
