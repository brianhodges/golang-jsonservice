package main
import (
	"fmt"
    "log"
    "os"
    "net/http"
    "strings"
    "regexp"
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

func validIDJsonUrl(url string) bool {
    Re := regexp.MustCompile(`^/users/\d.json$`)
    return Re.MatchString(url)
}

//all users index view handler
func index(w http.ResponseWriter, r *http.Request) {
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
}

//individual user show view handler
func show(w http.ResponseWriter, r *http.Request) {
    f := strings.Trim(r.URL.String(), "/users/")
    id := strings.Trim(f, ".json")
    if id != "" && validIDJsonUrl(r.URL.String()) {
        db, err := sql.Open("postgres", os.Getenv("BASIC_APP_DATABASE_URL"))
        check(err)
        defer db.Close()
    
        rows, err := db.Query("SELECT id, email, first_name, last_name FROM users where id = " + id)
        var user user.User
        check(err)
        for rows.Next() {
            err = rows.Scan(&user.ID, &user.Email, &user.First_Name, &user.Last_Name)
            check(err)
        }
        
        if user.ID != 0 {
            w.Header().Set("Content-Type", "application/json; charset=utf-8")
            b, err := json.Marshal(user)
            check(err)
            fmt.Fprintf(w, string(b[:]))
        } else {
            w.WriteHeader(http.StatusNotFound)
            fmt.Fprint(w, "Error 204 - Content (User) Not Found")
        }
        return
    } else {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprint(w, "Error 404 - Page not found")
    }
}

func main() {
	fmt.Println("Running local server @ http://localhost:8080")
    http.HandleFunc("/users.json", index)
    http.HandleFunc("/users/", show)
    log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))
}
