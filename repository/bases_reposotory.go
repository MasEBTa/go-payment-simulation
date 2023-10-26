package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-payment-simulation/model"
	"os"
)

// Buat variabel untuk menampung hasil dekoding
	type FileConfig struct {
    Files []struct {
        Name string `json:"name"`
        File string `json:"file"`
    } `json:"file"`
	}

type BasesRepository interface {
	FindByName(name string) (*model.FileJson, error)
}

type basesRepository struct {
	data FileConfig
}

// FindById implements BasesRepository.
func (u *basesRepository) FindByName(name string) (*model.FileJson, error) {
	// fmt.Println(name)
	// fmt.Println(u.data.Files)
	var mod model.FileJson
	for _, Bases := range u.data.Files {
		if Bases.Name == name {
			mod.File = Bases.File
			mod.Name = Bases.Name
			return &mod, nil
		}
	}
	return nil, errors.New("file not found")
}

func NewBasesRepository(jsonFilePath string) (BasesRepository, error) {
	file, err := os.Open(jsonFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// var data []model.FileJson
	// decoder := json.NewDecoder(file)
	// err = decoder.Decode(&data)
	// if err != nil {
	// 	return nil, err
	// }

	var fileConfig FileConfig

	// Dekode berkas JSON ke dalam variabel fileConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&fileConfig); err != nil {
			fmt.Println(err)
	}

	// fmt.Println(fileConfig.Files)
	// Sekarang Anda dapat mengakses data dalam variabel fileConfig
	// for _, f := range fileConfig.Files {
	// 		fmt.Printf("Name: %s\n", f.Name)
	// 		fmt.Printf("File: %s\n", f.File)
	// }
	
	return &basesRepository{
		data: fileConfig,
	}, nil
}
