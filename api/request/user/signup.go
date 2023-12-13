package user

type SignUp struct {
	Gender   string `json:"gender"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
