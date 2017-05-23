package user

type User struct {
    ID  int `json:"id"`
    Email string `json:"email"`
    First_Name string `json:"first_name"`
    Last_Name string `json:"last_name"`
}