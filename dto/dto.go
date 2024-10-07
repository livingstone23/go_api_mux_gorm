package dto

type CategoryDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProductDto struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	CategoryID  uint   `json:"category_id"`
}

type PerfilDto struct {
	Name string `json:"name"`
}

type UserDto struct {
	
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	PerfilID uint   `json:"perfil_id"`
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseDto struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}