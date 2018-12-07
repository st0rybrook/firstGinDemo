package models

import (
	"fmt"
	"log"

	db "gin/database"
)

type Person struct {
	Id        int    `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

func (p *Person) AddPerson() (id int64, err error) {
	res, err := db.SqlDb.Exec("insert into person(first_name,last_name) values (?,?)", p.FirstName, p.LastName)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	id, err = res.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	return
}
func (p *Person) GetPersons() (persons []Person, err error) {
	persons = make([]Person, 0)
	rows, err := db.SqlDb.Query("select id,first_name,last_name from person")
	defer rows.Close()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err.Error())
		return
	}
	return
}
func (p *Person) GetPerson() (person Person, err error) {
	err = db.SqlDb.QueryRow("select id,first_name,last_name from person where id=?", p.Id).Scan(
		&person.Id, &person.FirstName, &person.LastName,
	)
	fmt.Println(person.LastName)
	return
}
func (p *Person) UpdatePerson() (ra int64, err error) {
	stmt, err := db.SqlDb.Prepare("update person set first_name=?,last_name=? where id=?")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err.Error())
		return 0, err
	}
	res, err := stmt.Exec(p.FirstName, p.LastName, p.Id)
	if err != nil {
		log.Fatal(err.Error())
		return 0, err
	}
	ra, err = res.RowsAffected()
	return
}
func (p *Person) DeletePerson() (ra int64, err error) {
	res, err := db.SqlDb.Exec("delete  from person where id=? ", p.Id)
	if err != nil {
		log.Fatal(err.Error())
		return 0, err
	}
	ra, err = res.RowsAffected()
	return
}
