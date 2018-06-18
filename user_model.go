package main

func (u *user) Create() (*user, error) {
	result, err := myDatabase.Get().Exec(`INSERT INTO Users
													(username, email, password, dateCreated)
													VALUES
													(?, ?, ?, ?)`, u.Username(), u.Email(), u.Password(), u.DateCreated())
	if err != nil {
		return u, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return u, err
	}

	u.id = int(id)

	return u, nil
}

func (u *user) TableExists() bool {
	rows, err := myDatabase.Get().Query(`SHOW TABLES LIKE "Users"`)
	defer rows.Close()
	if err != nil {
		return false
	}

	for rows.Next() {
		return true
	}

	return false
}
