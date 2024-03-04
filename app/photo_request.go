package app

import "github.com/asaskevich/govalidator"

type CreatePhotoRequest struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photoUrl"`
}

type UpdatePhotoRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photoUrl"`
}

func ValidateCreatePhotoRequest(req CreatePhotoRequest) error {
	_, err := govalidator.ValidateStruct(req)
	return err
}

func ValidateUpdatePhotoRequest(req UpdatePhotoRequest) error {
	_, err := govalidator.ValidateStruct(req)
	return err
}
