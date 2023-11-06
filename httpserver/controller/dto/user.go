package dto

type UserResponseDTO struct {
	Id      int
	Name    string
	Age     int
	Email   string
	Mobile  string
	IP      string
	Address string
}

// Set validation function in "valid" tag
// Use ";" as the separator of multiple functions. Spaces accept after ";"
// Wrap parameters with "()" and separate parameter with ",". Spaces accept after ","
// Wrap regex match with "//"
type CreateUserRequestDTO struct {
	Name    string `valid:"Required;Match(/^Bee.*/)"` // Name cannot be empty and must have prefix Bee
	Age     int    `valid:"Range(1, 140)"`            // 1 <= Age <= 140
	Email   string `valid:"Email; MaxSize(100)"`      // Email must be valid email address and the length must be <= 100
	Mobile  string `valid:"Mobile"`                   // Mobile must be valid mobile
	IP      string `valid:"IP"`                       // IP must be a valid IPV4 address
	Address string `valid:"ChinaAddress"`
}

type CreateUserResponseDTO struct {
	Id      int
	Name    string
	Age     int
	Email   string
	Mobile  string
	Address string
}
