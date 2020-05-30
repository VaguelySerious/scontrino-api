package models

type Expense struct {
	ID    		uint    `json:"id" gorm:"primary_key"`
	Name  		string  `json:"name" validate:"min=3,max=40"`
	Category	string  `json:"category" validate:"min=1,max=20,regexp=^[a-z]*$"`
	Cost			float32 `json:"cost" validate:"min=0.01,max=1000000"`
	Sharing		float32 `json:"sharing" validate:"min=0.00,max=1.00"`
	Date			string  `json:"date" validate:"regexp=^(\\d{4}-([0]\\d|1[0-2])-([0-2]\\d|3[01]))?$"`
	Notes			string  `json:"notes" validate:"min=0,max=10000"`
	UserID		uint		`json:"user_id"`
	GroupID		uint		`json:"group_id" validate:"min=0"`
}

type CreateExpenseInput struct {
	Name  		string  `json:"name" binding:"required"`
	Category	string  `json:"category" binding:"required"`
	Cost			float32 `json:"cost" binding:"required"`
	Sharing		float32 `json:"sharing" binding:"required"`
	Date			string  `json:"date"`
	Notes			string  `json:"notes"`
	GroupID		uint		`json:"group_id"`
}

type UpdateExpenseInput struct {
	Name  		string  `json:"name"`
	Category	string  `json:"category"`
	Cost			float32 `json:"cost"`
	Sharing		float32 `json:"sharing"`
	Date			string  `json:"date"`
	Notes			string  `json:"notes"`
}