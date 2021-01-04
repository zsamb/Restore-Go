package database

import (
	"database/sql"
	"fmt"
	"github.com/Samb8104/Restore/utils/config"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	username string
	password string
	host     string
	port     int
	db       string
	debug    bool
	database *sql.DB
}

//Initialise Produces new DB
func Initialise(conf config.ConfFile) (*DB, error) {
	database, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		conf.Database.Username,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.DB,
	))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %s", err)
	}

	//Test query to make sure the connection is working
	_, err = database.Exec("SHOW TABLES")
	if err != nil {
		return nil, err
	}

	//Validate tables
	//User table
	_, err = database.Query(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS users (%s,%s,%s,%s,%s,%s,%s,%s,%s);`,
		"user_id int AUTO_INCREMENT PRIMARY KEY",
		"username varchar(255) NOT NULL",
		"first_name varchar(255)",
		"last_name varchar(255)",
		"role varchar(255)",
		"email varchar(255) NOT NULL",
		"createdAt date DEFAULT (NOW())",
		"lastModified date DEFAULT (NOW())",
		"password text NOT NULL"))
	if err != nil {
		return nil, err
	}

	//Group table
	_, err = database.Query(fmt.Sprintf("CREATE TABLE IF NOT EXISTS `groups` (%s,%s,%s,%s,%s,%s)",
		"`group_id` int AUTO_INCREMENT PRIMARY KEY",
		"`group_name` varchar(255) NOT NULL",
		"`createdBy` varchar(255) NOT NULL",
		"`createdAt` date DEFAULT (NOW())",
		"`lastModified` date DEFAULT (NOW())",
		"`permissions` text NOT NULL"))
	if err != nil {
		return nil, err
	}

	//Backup table
	_, err = database.Query(fmt.Sprintf("CREATE TABLE IF NOT EXISTS `backups` (%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)",
		"`backup_id` int AUTO_INCREMENT PRIMARY KEY",
		"`owner_id` int NOT NULL",
		"`notes` text",
		"`backup_name` varchar(255) NOT NULL",
		"`state` bool NOT NULL",
		"`size` int",
		"`sources` text NOT NULL",
		"`targets` text NOT NULL",
		"`createdAt` date DEFAULT (NOW())",
		"`lastModified` date DEFAULT (NOW())"))
	if err != nil {
		return nil, err
	}

	instance := DB{
		username: conf.Database.Username,
		password: conf.Database.Password,
		host:     conf.Database.Host,
		port:     conf.Database.Port,
		db:       conf.Database.DB,
		debug:    conf.Debug,
		database: database,
	}

	return &instance, nil
}

func (db DB) Host() string {
	return fmt.Sprintf("%s:%d", db.host, db.port)
}
