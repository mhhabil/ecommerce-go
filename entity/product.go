package entity

type Product struct {
	Name          string   `json:"name" binding:"min=5,max=60,required"`
	Price         int64    `json:"price" binding:"min=0,required"`
	ImageUrl      string   `json:"imageUrl" validate:"required,url"`
	Stock         int64    `json:"stock" binding:"min=0,required"`
	Condition     string   `json:"condition" validate:"oneof=new second"`
	Tags          []string `json:"tags" validate:"min=0"`
	IsPurchasable bool     `json:"isPurchasable" binding:"required"`
}
