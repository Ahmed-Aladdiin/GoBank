package types

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/ahmed-aladdiin/gobank/utils"
)

type PostgresStorage struct {
	db *sql.DB
}

///=============================================
//Initialize the Connection to the database

func NewPostgresStorage() (*PostgresStorage, error) {
	godotenv.Load(".env")
	conctString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", conctString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{
		db: db,
	}, nil
}

func (pgst *PostgresStorage) Init() error {
	return pgst.CreateAccountTable()
}

///=============================================
// Get the database ready

func (pgst *PostgresStorage) CreateAccountTable() error {
	query := `
  create table if not exists account (
    id serial primary key,
    firstName varchar(50),
    lastName varchar(50),
    number serial,
    balance float,
    encryptedPassword varchar(255),
    createdAt timestamp
  )
  `
	_, err := pgst.db.Exec(query)

	return err
}

///=============================================
// Account functions

func (pgst *PostgresStorage) CreateAccount(acc *Account) (*Account, error) {
	query := `
  insert into account(firstName, lastName, balance, encryptedPassword, createdAt)
  values($1, $2, $3, $4, $5)
  returning id, firstName, lastName, number,balance, createdAt;
  `
	newAcc := new(Account)

	err := pgst.db.QueryRow(query,
		acc.FirstName,
		acc.LastName,
		0,
		utils.StringWithCharset(12),
		time.Now().UTC(),
	).Scan(&newAcc.ID, &newAcc.FirstName, &newAcc.LastName, &newAcc.Number, &newAcc.Balance, &newAcc.CreatedAt)
	if err != nil {
		return nil, err
	}

	return newAcc, nil
}

func (pgst *PostgresStorage) DeleteAccount(id int) error {
	query := `
  delete
  from account
  where id = $1
  `
	result, err := pgst.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("Id was not found")
	}
	return nil
}

func (pgst *PostgresStorage) GetAccounts() ([]*Account, error) {
	query := `
  select * from account;
  `
	rows, err := pgst.db.Query(query)

	if err != nil {
		return nil, err
	} else {
		defer rows.Close()
	}

	accounts := []*Account{}

	for rows.Next() {
		account := new(Account)
		err := rows.Scan(&account.ID, &account.FirstName, &account.LastName,
			&account.Number, &account.Balance, &account.EncryptedPassword, &account.CreatedAt)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (pgst *PostgresStorage) GetAccountByID(id int) (*Account, error) {
	query := `
  select * 
  from account
  where id = $1;
  `
	account := new(Account)
	err := pgst.db.QueryRow(query, id).Scan(&account.ID, &account.FirstName, &account.LastName,
		&account.Number, &account.Balance, &account.EncryptedPassword, &account.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (pgst *PostgresStorage) UpdateAccount(acc *Account) error {
	return nil
}

///=============================================
// Transactions

func (pgst *PostgresStorage) Deposit(accNum int, amount float64) error {
	query := `
  update account
  set balance = balance + $1
  where number = $2
  `
	result, err := pgst.db.Exec(query, amount, accNum)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("Account number was not found")
	}

	return nil
}

func (pgst *PostgresStorage) Withdraw(accNum int, amount float64) error {
	query := `
  update account
  set balance = balance - $1
  where number = $2 and balance >= $1
  `
	result, err := pgst.db.Exec(query, amount, accNum)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("Not sufficient balance")
	}

	return nil
}

func (pgst *PostgresStorage) Transfare(accNumFrom int, accNumTo int, amount float64) error {
	query := `
  update account
  set balance = balance - $1
  where number = $2 and balance >= $1
  `
	result, err := pgst.db.Exec(query, amount, accNumFrom)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("Not sufficient balance")
	}

	query = `
  update account
  set balance = balance + $1
  where number = $2
  `
	_, err = pgst.db.Exec(query, amount, accNumTo)
	if err != nil {
		return err
	}

	return nil
}
