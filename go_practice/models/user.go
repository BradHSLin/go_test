package models

import (
	"go_practice/db"
)

type User struct {
	ID     int64
	Name   string  `binding:"required"`
	Email  string  `binding:"required"`
	Height float64 `binding:"required"`
	Weight float64 `binding:"required"`
	Group  int64   `binding:"required"`
}

func (u *User) Save() error {
	query := `
	INSERT INTO users ('name', 'email', 'height', 'weight', 'group') 
	VALUES(?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.Name, u.Email, u.Height, u.Weight, u.Group)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func GetAllUsers() ([]User, error) {
	query := "SELECT * FROM users"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Height, &user.Weight, &user.Group)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetUserByID(id int64) (*User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Height, &user.Weight, &user.Group)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (user User) Update() error {
	query := `
	UPDATE users
	SET 'name' = ?, 'email' = ?, 'height' = ?, 'weight' = ?, 'group' = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Height, user.Weight, user.Group, user.ID)
	return err

}

func (user User) Delete() error {
	query := "DELETE FROM users WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.ID)
	return err

}

func EmptyUser() error {
	query := "DELETE FROM users"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	defer stmt.Close()

	return nil
}
