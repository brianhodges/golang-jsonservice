package main
import (
	"fmt"
    "log"
    "os"
    "net/http"
    "database/sql"
    "encoding/json"
    "golang-jsonservice/pkg/user"
    _ "github.com/lib/pq"
)

const br = "<br>"

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
	fmt.Println("Running local server @ http://localhost:8080")
    
    http.HandleFunc("/users.json", func(w http.ResponseWriter, r *http.Request) {
        url := r.FormValue("url")
        if url == "" {    
            db, err := sql.Open("postgres", os.Getenv("BASIC_APP_DATABASE_URL"))
            check(err)
            defer db.Close()

            rows, err := db.Query("SELECT id, email, first_name, last_name FROM users")
            users := []user.User{}
            check(err)
            for rows.Next() {
                var u user.User
                err = rows.Scan(&u.ID, &u.Email, &u.First_Name, &u.Last_Name)
                check(err)
                users = append(users, u)
            }
            
            w.Header().Set("Content-Type", "application/json; charset=utf-8")
            b, err := json.Marshal(users)
            check(err)
            fmt.Fprintf(w, string(b[:]))
            return
        }
    })

    log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))
}
