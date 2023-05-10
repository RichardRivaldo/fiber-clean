package repositories

import (
	"context"
	"fiber-clean/src/configs"
	"time"

	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func UploadImage(image interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := configs.Storage.Upload.Upload(ctx, image, uploader.UploadParams{
		Folder: configs.GetEnv("CLOUDINARY_UPLOAD_FOLDER"),
	})
	if err != nil {
		return "", err
	}

	return result.SecureURL, nil
}
