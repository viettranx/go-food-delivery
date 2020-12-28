package model

import "time"

//type FoodLikes struct {
//	UserId    int32      `json:"user_id" gorm:"column:user_id;"`
//	FoodId    int32      `json:"food_id" gorm:"column:food_id"`
//	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
//}

type FoodLikes struct {
	UserId    int32      `json:"user_id" gorm:"primaryKey;column:user_id;"`
	FoodId    int32      `json:"food_id" gorm:"primaryKey;column:food_id"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
}

func (FoodLikes) TableName() string {
	return "food_likes"
}
