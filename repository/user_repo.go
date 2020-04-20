package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/vamshi/go-crud/model"
)

// GetUserData gets user data.
func GetUserData() []model.User{



	// create a user slice
	users := make([]model.User,0)

	user1 := model.User{
		Firstname : "Vamshi",
		Lastname: "Muthyapu",
		Address: model.Address{
			City: "some city",
			State: "some state",
		},

	}

	user2 := model.User{
		Firstname : "James",
		Lastname: "Bond",
		Address: model.Address{
			City: "london",
			Country: "UK",
		},

	}
	
	// append users to slice
	users = append(users,user1,user2)
	return users

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