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
