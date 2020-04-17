package repository

import "github.com/vamshi/go-crud/model"

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