package user 

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)
type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email string `bson:"email " json:"email"`
	Password string `bson:"password" json:"password"`
	Role string `bson:"role" json:"role"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
}

type PublicUser struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email string `bson:"email" json:"email"`
	Role string `bson:"role" json:"role"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
}

func ToPublicUser(u *User) *PublicUser {
	return &PublicUser{
		ID: u.ID.Hex(),
		Email: u.Email,
		Role: u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}