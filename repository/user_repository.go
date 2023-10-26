package repository

import (
	"encoding/json"
	"errors"
	"go-payment-simulation/model"
	"os"
	"strings"
)

// Buat variabel untuk menampung hasil dekoding
type fileConfig struct {
	User []model.UserCredential
}

type UserRepository interface {
	FindById(id string) (model.UserCredential, error)
	FindByEmail(email string) (model.UserCredential, error)
	FindByRole(role string) ([]model.UserCredential, error)
	FIndByUsername(username string) (model.UserCredential, error)
	NewUser(payload model.UserCredential) (model.UserCredential, error)
}

type userRepository struct {
	data         []model.UserCredential
	jsonFilePath string
}

// FindByRole implements UserRepository.
func (u *userRepository) FindByRole(role string) ([]model.UserCredential, error) {
	var marchand []model.UserCredential
	for _, mc := range u.data {
		if mc.UserRole == role {
			marchand = append(marchand, mc)
		}
	}
	return marchand, nil
}

// FindByEmail implements UserRepository.
func (u *userRepository) FindByEmail(email string) (model.UserCredential, error) {
	for _, user := range u.data {
		if strings.EqualFold(user.Email, email) {
			return user, nil
		}
	}
	return model.UserCredential{}, errors.New("email not found")
}

// FIndByUsername implements UserRepository.
func (u *userRepository) FIndByUsername(username string) (model.UserCredential, error) {
	for _, user := range u.data {
		if strings.EqualFold(user.Username, username) {
			return user, nil
		}
	}
	return model.UserCredential{}, errors.New("username is not found")
}

// NewUser implements UserRepository.
func (u *userRepository) NewUser(payload model.UserCredential) (model.UserCredential, error) {
	// Tambahkan pengguna baru ke data
	u.data = append(u.data, payload)

	// Simpan data ke berkas JSON
	if err := u.saveDataToJSON(); err != nil {
		return model.UserCredential{}, err
	}

	// Mengembalikan pengguna baru
	return payload, nil
}

func (u *userRepository) FindById(id string) (model.UserCredential, error) {
	// fmt.Println(id)
	for _, user := range u.data {
		if user.Id == id {
			return user, nil
		}
	}
	return model.UserCredential{}, errors.New("user not found")
}

func (u *userRepository) saveDataToJSON() error {
	file, err := os.Create(u.jsonFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	userData := fileConfig{User: u.data}
	// Menggunakan encoder JSON untuk menulis data ke berkas JSON
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(userData); err != nil {
		return err
	}

	return nil
}

func NewUserRepository(jsonFilePath string) (UserRepository, error) {
	file, err := os.Open(jsonFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var fileConfig fileConfig

	// Dekode berkas JSON ke dalam variabel fileConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&fileConfig); err != nil {
		return nil, err
	}

	return &userRepository{
		data:         fileConfig.User,
		jsonFilePath: jsonFilePath,
	}, nil
}
