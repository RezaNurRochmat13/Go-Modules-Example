package service

import (
	"learning-gomod/domain/dao"
	"learning-gomod/domain/repository"
	"learning-gomod/resources"
	"log"
)

func FetchAllUsers() ([]dao.ListUser, error) {
	databaseConfig := resources.DatabaseConnectionPostgres()

	defer databaseConfig.Close()

	var (
		daoMapperUser dao.ListUser
		resultAllUser []dao.ListUser
	)

	findAllUserRepository := repository.FindAllUser()

	queryAllUser, errorHandlerQuery := databaseConfig.Query(findAllUserRepository)

	if errorHandlerQuery != nil {
		log.Printf("Error when executing query %s", errorHandlerQuery)
		return nil, errorHandlerQuery
	}

	for queryAllUser.Next() {
		errorScanHandler := queryAllUser.Scan(
			&daoMapperUser.ID,
			&daoMapperUser.Name,
			&daoMapperUser.Address,
			&daoMapperUser.PhoneNumber)

		if errorScanHandler != nil {
			log.Printf("Error scan data %s", errorScanHandler)
			return nil, errorScanHandler
		}

		resultAllUser = append(resultAllUser, daoMapperUser)
	}

	return resultAllUser, nil
}

func FetchUserByID(userID int) ([]dao.DetailUser, error) {
	databaseConfig := resources.DatabaseConnectionPostgres()

	defer databaseConfig.Close()

	var (
		daoMapperUser    dao.DetailUser
		resultDetailUser []dao.DetailUser
	)

	findUserByIDRepository := repository.FindUserById()

	queryDetailUser, errorHandlerQuery := databaseConfig.Query(findUserByIDRepository, userID)

	if errorHandlerQuery != nil {
		log.Printf("Error when execute query %s", errorHandlerQuery)
		return nil, errorHandlerQuery
	}

	for queryDetailUser.Next() {
		errorHandlerScan := queryDetailUser.Scan(
			&daoMapperUser.ID,
			&daoMapperUser.Name,
			&daoMapperUser.Address,
			&daoMapperUser.PhoneNumber)

		if errorHandlerScan != nil {
			log.Printf("Error when scan value %s", errorHandlerScan)
			return nil, errorHandlerScan
		}

		resultDetailUser = append(resultDetailUser, daoMapperUser)
	}

	return resultDetailUser, nil
}

func SaveNewUser(createNewUserPayload dao.CreateNewUser) error {
	databaseConfig := resources.DatabaseConnectionPostgres()

	defer databaseConfig.Close()

	saveNewUserRepository := repository.SaveUser()

	initTransaction, errorHandlerTransaction := databaseConfig.Begin()

	if errorHandlerTransaction != nil {
		log.Printf("Error when init transaction %s", errorHandlerTransaction)
		return errorHandlerTransaction
	}

	defer initTransaction.Rollback()

	stmtPrepareNewUser, errorHandlerPrepareStmt := databaseConfig.Prepare(saveNewUserRepository)

	if errorHandlerPrepareStmt != nil {
		log.Printf("Error when prepared stmt %s", errorHandlerPrepareStmt)
		initTransaction.Rollback()
		return errorHandlerPrepareStmt
	}

	_, errorHandlerExecQuery := stmtPrepareNewUser.Exec(
		createNewUserPayload.Name,
		createNewUserPayload.Address,
		createNewUserPayload.PhoneNumber,
		createNewUserPayload.CreatedAt,
		createNewUserPayload.UpdatedAt)

	if errorHandlerExecQuery != nil {
		log.Printf("Error when inserting db %s", errorHandlerExecQuery)
		initTransaction.Rollback()
		return errorHandlerExecQuery
	}

	errorHandlerCommitTrans := initTransaction.Commit()

	if errorHandlerCommitTrans != nil {
		log.Printf("Error transaction we'll roolback %s", errorHandlerCommitTrans)
		initTransaction.Rollback()
		return errorHandlerCommitTrans
	}

	return nil
}

func UpdateUser(userID int, updateUserPayload dao.UpdateUser) error {
	databaseConfig := resources.DatabaseConnectionPostgres()

	defer databaseConfig.Close()

	updateUserRepository := repository.UpdateUser()

	initTransaction, errorHandlerTransaction := databaseConfig.Begin()

	if errorHandlerTransaction != nil {
		log.Printf("Error when init transaction %s", errorHandlerTransaction)
		return errorHandlerTransaction
	}

	defer initTransaction.Rollback()

	stmtPrepareUpdateUser, errorHandlerPrepareStmt := databaseConfig.Prepare(updateUserRepository)

	if errorHandlerPrepareStmt != nil {
		log.Printf("Error when prepared stmt %s", errorHandlerPrepareStmt)
		initTransaction.Rollback()
		return errorHandlerPrepareStmt
	}

	_, errorHandlerExecQuery := stmtPrepareUpdateUser.Exec(
		updateUserPayload.Name,
		updateUserPayload.Address,
		updateUserPayload.PhoneNumber,
		updateUserPayload.CreatedAt,
		updateUserPayload.UpdatedAt,
		userID)

	if errorHandlerExecQuery != nil {
		log.Printf("Error when inserting db %s", errorHandlerExecQuery)
		initTransaction.Rollback()
		return errorHandlerExecQuery
	}

	errorHandlerCommitTrans := initTransaction.Commit()

	if errorHandlerCommitTrans != nil {
		log.Printf("Error transaction we'll roolback %s", errorHandlerCommitTrans)
		initTransaction.Rollback()
		return errorHandlerCommitTrans
	}

	return nil
}

func DeleteUser(userID int) error {
	databaseConfig := resources.DatabaseConnectionPostgres()

	defer databaseConfig.Close()

	deleteUserRepository := repository.DeleteUser()

	initTransaction, errorHandlerTransaction := databaseConfig.Begin()

	if errorHandlerTransaction != nil {
		log.Printf("Error when init transaction %s", errorHandlerTransaction)
	}

	defer initTransaction.Rollback()

	stmtPrepareDeleteUser, errorHandlerPrepareStmt := databaseConfig.Prepare(deleteUserRepository)

	if errorHandlerPrepareStmt != nil {
		log.Printf("Error when prepared stmt %s", errorHandlerPrepareStmt)
		initTransaction.Rollback()
		return errorHandlerPrepareStmt
	}

	_, errorHandlerExecQuery := stmtPrepareDeleteUser.Exec(userID)

	if errorHandlerExecQuery != nil {
		log.Printf("Error when inserting db %s", errorHandlerExecQuery)
		initTransaction.Rollback()
		return errorHandlerExecQuery
	}

	errorHandlerCommitTrans := initTransaction.Commit()

	if errorHandlerCommitTrans != nil {
		log.Printf("Error transaction we'll roolback %s", errorHandlerCommitTrans)
		initTransaction.Rollback()
		return errorHandlerCommitTrans
	}

	return nil
}
