package service

import (
	"fmt"
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
	}

	for queryAllUser.Next() {
		errorScanHandler := queryAllUser.Scan(
			&daoMapperUser.ID,
			&daoMapperUser.Name,
			&daoMapperUser.Address,
			&daoMapperUser.PhoneNumber)

		if errorScanHandler != nil {
			log.Printf("Error scan data %s", errorScanHandler)
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
	}

	for queryDetailUser.Next() {
		errorHandlerScan := queryDetailUser.Scan(
			&daoMapperUser.ID,
			&daoMapperUser.Name,
			&daoMapperUser.Address,
			&daoMapperUser.PhoneNumber)

		if errorHandlerScan != nil {
			log.Printf("Error when scan value %s", errorHandlerScan)
		}

		fmt.Println(daoMapperUser)

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
	}

	defer initTransaction.Rollback()

	stmtPrepareNewUser, errorHandlerPrepareStmt := databaseConfig.Prepare(saveNewUserRepository)

	if errorHandlerPrepareStmt != nil {
		log.Printf("Error when prepared stmt %s", errorHandlerPrepareStmt)
	}

	_, errorHandlerExecQuery := stmtPrepareNewUser.Exec(
		createNewUserPayload.Name,
		createNewUserPayload.Address,
		createNewUserPayload.PhoneNumber,
		createNewUserPayload.CreatedAt,
		createNewUserPayload.UpdatedAt)

	if errorHandlerExecQuery != nil {
		log.Printf("Error when inserting db %s", errorHandlerExecQuery)
	}

	errorHandlerCommitTrans := initTransaction.Commit()

	if errorHandlerCommitTrans != nil {
		log.Printf("Error transaction we'll roolback %s", errorHandlerCommitTrans)
	}

	return nil
}
