package repository

func FindAllUser() string {
	return "SELECT users.id, " +
		"users.name, " +
		"users.address, " +
		"users.phone_number " +
		"FROM users"
}

func FindUserById() string {
	return "SELECT users.id, " +
		"users.name, " +
		"users.address, " +
		"users.phone_number " +
		"FROM users " +
		"WHERE users.id = $1"
}

func SaveUser() string {
	return "INSERT INTO users(name, address, phone_number, created_at, updated_at) " +
		"VALUES ($1, $2, $3, $4, $5)"
}

func UpdateUser() string {
	return "UPDATE users SET name = $1, " +
		"address = $2, " +
		"phone_number = $3, " +
		"created_at = $4, " +
		"updated_at = $5 WHERE users.id = $6"
}

func DeleteUser() string {
	return "DELETE FROM users WHERE users.id = $1"
}
