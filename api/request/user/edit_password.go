package user

type EditPassword struct {
	EditPassword    string `json:"edit_password"`
	CurrentPassword string `json:"current_password"`
}
