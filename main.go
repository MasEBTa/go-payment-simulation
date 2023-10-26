package main

import "go-payment-simulation/delivery"

func main() {
	delivery.NewServer().Run()
}

// import (
// 	"fmt"
// 	"go-payment-simulation/config"
// 	"go-payment-simulation/model"
// 	"go-payment-simulation/repository"
// )

// func main() {
// 		cfg, err := config.NewConfig()
//     if err != nil {
//         // Handle error
//         fmt.Println("Error reading configuration:", err)
//         return
//     }
//     // Membuat instance basesRepository dengan data dari berkas JSON
//     basesRepository, err := repository.NewBasesRepository(cfg.JSFile)
//     if err != nil {
//         panic(err)
//     }

//     // Menggunakan basesRepository untuk mencari nama file
//     name := "user"
//     file, err := basesRepository.FindByName(name)
//     if err != nil {
//         // Handle error jika pengguna tidak ditemukan atau ada kesalahan lain
//         panic(err)
//     }

//     // Menampilkan file
//     // fmt.Printf("Name: %s, File: %s", file.Name, file.File)

//     // Melanjutkan logika aplikasi setelah mendapatkan data file
// 		// Membuat instance basesRepository dengan data dari berkas JSON
//     userRepository, err := repository.NewUserRepository(file.File)
//     if err != nil {
//         panic(err)
//     }
// 		user, err := userRepository.FindById("a04378fb-77ef-4878-8cec-1227b078432a")
// 		if err!=nil {
// 			fmt.Print(err)
// 		}
// 		fmt.Printf("Id: %s, Username: %s, Email: %s", user.Id, user.Username, user.Email)

// 		user = model.UserCredential{
// 			Username: "d",
// 			Email:    "jhdns@example.com",
// 			Password: "12345",
// 			// ... tambahkan field lain yang sesuai
// 		}
// 		userRepository.NewUser(user)
// 	// 	userusecase := usecase.NewUserUseCase(userRepository)

		
// 	// // membuat user baru lewat usecase
// 	// user := model.UserCredential{
// 	// 	Username: "Bomkiadn",
// 	// 	Email:    "jhdns@example.com",
// 	// 	Password: "12345",
// 	// 	// ... tambahkan field lain yang sesuai
// 	// }
	
// 	// newUser, err := userusecase.NewUser(user)
// 	// 	if err != nil {
// 	// 		// Handle error jika pengguna tidak ditemukan atau ada kesalahan lain
// 	// 		panic(err)
// 	// 	}
// 	// 	// Menampilkan file
// 	// 	fmt.Printf("Id: %s, Username: %s, Email: %s", newUser.Id, newUser.Username, newUser.Email)

// 	// 	// find user by id
// 	// 	user, err = userusecase.FindById(newUser.Id)
// 	// 	// Menampilkan file
// 	// 	fmt.Printf("Id: %s, Username: %s, Email: %s", user.Id, user.Username, user.Email)

// }
