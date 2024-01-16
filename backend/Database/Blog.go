package Database

//blog
//title, name, description,

import (
	"basic-blog/Utils"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func BlogView(title string)  (Blog, error) {
	query := "SELECT * FROM blog WHERE title = ?"
	row := db.QueryRow(query, title)

	var blog Blog
	err := row.Scan(&blog.Title, &blog.Description, &blog.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			return Blog{}, fmt.Errorf("kullanıcı bulunamadı")
		}
		return Blog{}, err
	}
	return blog, nil
}


func BlogAdd(title, description string) (bool, error) {
    token := Utils.Token(5)
    query := "INSERT INTO blog (token, title, description) VALUES (?, ?, ?)"
    
    _, err := db.Exec(query, token, title, description)
    if err != nil {
        fmt.Println(err)
        return false, err
    }

    return true, nil
}

func BlogDel(token string) (error, bool) {
    query := "DELETE FROM blog WHERE token=?"
    _, err := db.Exec(query, token)
    if err != nil {
        fmt.Println(err)
        return err, false
    }

    return nil, true
}

func BlogList() ([]Blog, error) {
	query := "SELECT * FROM blog"
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	blogs := make([]Blog, 0)
	for results.Next() {
		var blog Blog
		if err := results.Scan(&blog.Title, &blog.Description, &blog.Token); err != nil {
			fmt.Println(err)
			return nil, err
		}

		blogs = append(blogs, blog)
	}

	return blogs, nil
}
