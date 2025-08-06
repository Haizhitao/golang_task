package topic1

import "github.com/jmoiron/sqlx"

type Employee struct {
	ID         int
	Name       string
	department string
	salary     int
}

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func Run(db *sqlx.DB) {
	var emps []Employee
	err := db.Select(&emps, "SELECT id, name, department, salary FROM employee WHERE department = ? ORDER BY id DESC", "技术部")
	if err != nil {
		panic(err)
	}
}

func Run2(db *sqlx.DB) {
	var emp Employee
	err := db.Get(&emp, "SELECT id, name, department, salary FROM employee ORDER BY salary DESC LIMIT 1")
	if err != nil {
		panic(err)
	}
}

func Run3(db *sqlx.DB) {
	var books []Book
	minPrice := 50
	err := db.Select(&books, "SELECT id, title, author, price FROM books WHERE price > ? ORDER BY price DESC", minPrice)
	if err != nil {
		panic(err)
	}
}
