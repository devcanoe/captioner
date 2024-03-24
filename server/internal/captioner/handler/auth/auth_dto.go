package auth

type Signin struct {
	Email    string `json:"email" bson:"email" validate:"required"`
	Password string `json:"password" bson:"password" validate:"required"`
	IP       string `json:"ip" bson:"ip"`
	Device   string `json:"device" bson:"device"`
}

type Signup struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
type Token string
