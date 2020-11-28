package model

import (
	pb "github.com/cg-0508/laracom/user-service/proto/user"
	"github.com/jinzhu/gorm"
)

type PasswordReset struct {
	gorm.Model
	Email string `gorm:"index"`
	Token string `gorm:"not null"`
}

func (model *PasswordReset) ToORM(req *pb.PasswordReset) (*PasswordReset, error) {
	model.Email = req.Email
	model.Token = req.Token
	return model, nil
}

func (model *PasswordReset) ToProtobuf() (*pb.PasswordReset, error) {
	var reset = &pb.PasswordReset{}
	reset.Email = model.Email
	reset.Token = model.Token
	return reset, nil
}