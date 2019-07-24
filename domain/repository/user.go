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
