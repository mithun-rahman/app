package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountById(int) (*Account, error)
}

type PostgresStore struct {
	db *gorm.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	dbHost := "localhost"
	dbPort := 5432
	dbName := "my_app"
	dbUser := "postgres"
	dbPassword := "1234"
	dbSSLMode := "disable"
	dbTimeZone := "UTC"

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode, dbTimeZone)

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}))
	if err != nil {
		return nil, err
	}
	fmt.Println("connected to database")
	return &PostgresStore{db}, nil
}

func (s *PostgresStore) CreateAccount(account *Account) error {
	return s.db.Save(account).Error
}

func (s *PostgresStore) UpdateAccount(account *Account) error {
	return s.db.Save(account).Error
}

func (s *PostgresStore) DeleteAccount(id int) error {
	account := &Account{}
	if err := s.db.Where("id = ?", id).First(&account).Error; err != nil {
		return err
	}
	return s.db.Delete(account).Error
}

func (s *PostgresStore) GetAccountById(id int) (*Account, error) {
	account := &Account{}
	if err := s.db.Where("id = ?", id).First(&account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	var accounts []*Account
	if err := s.db.Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	err := s.db.AutoMigrate(&Account{})
	if err == nil {
		log.Println("account table has been created")
	}
	return err
}

func (s *PostgresStore) dropAccountTable() error {
	err := s.db.Migrator().DropTable(&Account{})
	if err == nil {
		log.Println("account table has been deleted")
	}
	return err
}
