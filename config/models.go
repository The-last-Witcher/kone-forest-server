package config

type Chalet struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Address string `json:"address"`
	Price uint32 `json:"price"`
	Rating float32 `json:"rating"`
	GuestsMax uint `json:"guestsMax"`
	Bedrooms uint `json:"bedrooms"`
	Bathrooms uint `json:"bathrooms"`
	IsActive bool `json:"is_active"`
	Amenities []string `json:"amenities"`
}

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	AvatarURL string `json:"avatar_url"`
}

type DTO struct {
	User *User `json:"user"`
	Chalet []Chalet `json:"chalet"`
}

type InquiryReq struct{
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Topic string `json:"topic"`
	Message string `json:"message"`
}