package main

import (
	"fmt"

	"github.com/broucz/tmp-go-arch-test/domain/model/user"
	"github.com/broucz/tmp-go-arch-test/impl/memory"
)

func main() {

	// Setup repositories
	var (
		userRepo user.Repository
	)

	userRepo = memory.NewUserRepository()

	u1 := user.New("ID_1", "user_1", "u1@mail.com")
	u2 := user.New("ID_2", "user_2", "u2@mail.com")
	u3 := user.New("ID_3", "user_3", "u3@mail.com")

	userRepo.Store(u1)
	userRepo.Store(u2)
	userRepo.Store(u3)

	findUser2, _ := userRepo.Find("ID_2")

	fmt.Printf("%#v\n", findUser2)
	fmt.Printf("%#v\n", findUser2.GetInfo())
	fmt.Printf("%#v\n", userRepo.FindAll())
}
