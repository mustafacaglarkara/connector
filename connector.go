package connector

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

type Connnector struct {
}
type database struct {
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
}

func getDbConnectionValues(dbType string) (database, error) {
	var db database

	// .env dosyası yüklenir. hata var ise hata döner.
	err := godotenv.Load()
	if err != nil {
		return db, err
	}

	// veritabanı bağlantı bilgilerinin isimleri
	hostName := "DB_" + strings.ToUpper(dbType) + "_HOST"
	portName := "DB_" + strings.ToUpper(dbType) + "_PORT"
	dbName := "DB_" + strings.ToUpper(dbType) + "_NAME"
	userName := "DB_" + strings.ToUpper(dbType) + "_USER"
	passwordName := "DB_" + strings.ToUpper(dbType) + "_PASSWORD"

	// veritabanı bağlantı bilgilerini oku ve "DbModel" struct'ına atayın
	db = database{
		DbHost:     os.Getenv(hostName),
		DbPort:     os.Getenv(portName),
		DbName:     os.Getenv(dbName),
		DbUser:     os.Getenv(userName),
		DbPassword: os.Getenv(passwordName),
	}

	return db, nil
}

func (connector *Connnector) CreateConnectionStr(dbSrc string, dbType string) (*sql.DB, error) {

	var db *sql.DB
	var err error
	var values, _ = getDbConnectionValues(dbType)

	switch dbSrc {
	case "mysql":
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", values.DbUser, values.DbPassword, values.DbHost, values.DbPort, values.DbName))
	case "mssql":
		db, err = sql.Open("mssql", fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", values.DbHost, values.DbUser, values.DbPassword, values.DbName))
	default:
		err = fmt.Errorf("Unsupported database type: %s", dbSrc)
	}

	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
