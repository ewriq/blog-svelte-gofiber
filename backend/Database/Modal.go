package Database

type User struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Perm 	 string `json:"perm"`
}

type Blog struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Token string `json:"token"`
}