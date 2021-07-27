package main

import (
	"errors"
	"fmt"

	"github.com/leumas3003/db-client/sqlclient"
)

const (
	queryGetuser = "SELECT id, email FROM public.users WHERE id=%d;"
	DbConn       = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
)

var (
	dbClient sqlclient.SqlClient
)

type User struct {
	Id    int64
	Email string
}

func init() {
	sqlclient.StartMockServer()

	var err error
	dbClient, err = sqlclient.Open("postgres", fmt.Sprintf(DbConn, "localhost", "5432", "USERNAME", "", "DATABASENAME"))
	if err != nil {
		panic(err)
	}
}
func main() {
	user, err := GetUser(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(user.Id)
	fmt.Println(user.Email)
}

func GetUser(id int64) (*User, error) {
	sqlclient.AddMock(sqlclient.Mock{
		Query:   fmt.Sprintf("SELECT id, email FROM public.users WHERE id=%v;", id),
		Args:    []interface{}{1},
		Error:   errors.New("error creating query"),
		Columns: []string{"id", "email"},
		Rows: [][]interface{}{
			{1, "email1"},
			{2, "email2"},
		},
	})

	rows, err := dbClient.Query(fmt.Sprintf(queryGetuser, id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user User
	for rows.HasNext() {
		if err := rows.Scan(&user.Id, &user.Email); err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, errors.New("user not found")
}
