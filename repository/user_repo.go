package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/gits/docker-test/model"
)

// GetUserData gets user data for a given id.
func GetUserData(id int) model.User{
	connString := `postgres://postgres:1165@localhost/boda_orda`
	conn, err := pgx.Connect(context.Background(),connString)

	if err != nil {
		log.Fatalf("error connecting to postgres using pgx. Error: %v",err)
	}

	defer conn.Close(context.Background())

	var user model.User

	err = conn.QueryRow(context.Background(),`select firstname, lastname from public.user where id = $1;`,id).
	Scan(&user.Firstname,&user.Firstname)

	if err != nil {
		log.Fatal("error retreiving data from postgres db")
	}

	return user
	

}

//InsertRecord inserts a user into postgres db
func InsertRecord(user model.User)  {
	log.Printf("value is %+v",user)
	connString := `postgres://steven:password@fullstack-postgres:5432/fullstack_api`
	conn, err := pgx.Connect(context.Background(),connString)

	if err != nil {
		log.Fatalf("error connecting to postgres using pgx. Error: %v",err)
	}

	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(),`insert into public.user(
		firstname,
		lastname) values( $1,$2);`,user.Firstname,user.Lastname)
	
	if err != nil {
		log.Fatalf("error occurred while inserting data to postgres %v",err)
	}
	
	

}