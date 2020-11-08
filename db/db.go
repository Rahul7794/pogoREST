package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jmoiron/sqlx"
)

// Connection stores various fields required for connecting to DB
type Connection struct {
	DbType   string // DbType represents different tradition db type
	User     string `json:"user"`     // User is the username of db
	Password string `json:"password"` // Password for connecting to db
	Host     string `json:"host"`     // Host of the db
	Port     string `json:"port"`     // Port to connect
	Db       string `json:"db"`       // Db is the database name
}

// GetConnection returns a Connection reference
func GetConnection(dbType string) (*Connection, error) {
	file, err := ioutil.ReadFile(fmt.Sprintf("connection/%s.json", dbType))
	if err != nil {
		return nil, fmt.Errorf("could not open the file %v", err)
	}
	connection := Connection{}

	err = json.Unmarshal([]byte(file), &connection)
	if err != nil {
		return nil, fmt.Errorf("could not parse connection json %v", err)
	}
	connection.DbType = dbType
	return &connection, nil
}

// ConnectDB connects to DB and returns database pointers or error
func ConnectDB(con *Connection) (*sqlx.DB, error) {
	db, err := sqlx.Connect(con.DbType, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s ",
		con.Host, con.Port, con.User, con.Password, con.Db, "disable"))
	if err != nil {
		return nil, err
	}
	return db, nil
}
