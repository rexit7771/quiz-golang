package structs

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Created_at  string `json:"created_at"`
	Created_by  string `json:"created_by"`
	Modified_at string `json:"modified_at"`
	Modified_by string `json:"modified_by"`
}
