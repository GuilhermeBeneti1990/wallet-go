package database

import (
	"database/sql"
	"testing"

	"github.com/GuilhermeBeneti1990/wallet-go/internal/entities"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	client        *entities.Client
	client2       *entities.Client
	accountFrom   *entities.Account
	accountTo     *entities.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Require().NoError(err)

	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance float, created_at date)")
	db.Exec("CREATE TABLE transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount float, created_at date)")

	client, err := entities.NewClient("John Doe", "j@email.com")
	s.Require().NoError(err)
	s.client = client

	client2, err := entities.NewClient("Jane Doe", "j2@email.com")
	s.Require().NoError(err)
	s.client2 = client2

	accountFrom := entities.NewAccount(s.client)
	accountFrom.Balance = 1000
	s.accountFrom = accountFrom

	accountTo := entities.NewAccount(s.client2)
	accountTo.Balance = 1000
	s.accountTo = accountTo

	s.transactionDB = NewTransactionDB(db)
}

func (s *TransactionDBTestSuite) TearDownTest() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entities.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)
	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}
