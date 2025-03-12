package database

import (
	"database/sql"
	"testing"

	"github.com/GuilhermeBeneti1990/wallet-go/internal/entities"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDb *ClientDB
}

func (s *ClientDBTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Require().NoError(err)

	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.clientDb = NewClientDB(db)
}

func (s *ClientDBTestSuite) TearDownTest() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestCreate() {
	client := &entities.Client{
		ID:    "1",
		Name:  "John Doe",
		Email: "j@email.com",
	}

	err := s.clientDb.Save(client)
	s.Require().NoError(err)
}

func (s *ClientDBTestSuite) TestGet() {
	client, _ := entities.NewClient("John Doe", "j@email.com")
	s.clientDb.Save(client)

	clientDB, err := s.clientDb.Get(client.ID)
	s.Require().NoError(err)
	s.Equal(client.ID, clientDB.ID)
	s.Equal(client.Name, clientDB.Name)
	s.Equal(client.Email, clientDB.Email)
}
