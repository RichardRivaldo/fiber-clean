package configs

import (
	"log"

	"github.com/cloudinary/cloudinary-go"
)

func GenerateStorageClient() *cloudinary.Cloudinary {
	storage, err := cloudinary.NewFromParams(
		GetEnv("CLOUDINARY_CLOUD_NAME"),
		GetEnv("CLOUDINARY_API_KEY"),
		GetEnv("CLOUDINARY_API_SECRET"),
	)

	if err != nil {
		log.Fatal(err)
	}

	return storage
}

var Storage = GenerateStorageClient()
