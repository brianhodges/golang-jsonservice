package user

type User struct {
    ID  int `json:"id"`
    Email string `json:"email"`
    First_Name string `json:"first_name"`
    Last_Name string `json:"last_name"`
    Password_Salt string `json:"password_salt"`
    Password_Hash string `json:"password_hash"`
    Role_ID int `json:"role_id"`
}