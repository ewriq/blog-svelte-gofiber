package Handler

type UserBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserInfo struct {
	Token string `json:"token"`
}
type BlogForm struct {
	Title string `json:"title"`
	Description string `json:"description"`
}
type BlogInfo struct {
	Title string `json:"title"`
}
type BlogDels struct {
	Token string `json:"token"`
}