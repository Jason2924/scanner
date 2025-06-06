package databases

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Jason2924/scanner/backend/config"
	ntt "github.com/Jason2924/scanner/backend/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	conn *gorm.DB
	once sync.Once
)

type IMysqlDatabase interface {
	Connect() *gorm.DB
	Close() error
	Ping(ctxt context.Context) error
}

type connection struct {
	idle     int
	maximum  int
	lifetime time.Duration
}

type mysqlDatabase struct {
	host         string
	rootPassword string
	name         string
	username     string
	password     string
	migrateTable bool
	connection   *connection
}

func NewMysqlDatabase(options *config.ConfigMysql) IMysqlDatabase {
	pool := &connection{
		idle:     10,
		maximum:  20,
		lifetime: 1 * time.Hour,
	}
	return &mysqlDatabase{
		host:         options.Host,
		rootPassword: options.RootPassword,
		name:         options.Name,
		username:     options.Username,
		password:     options.Password,
		migrateTable: options.MigrateTable,
		connection:   pool,
	}
}

// creating a connection to database
func (dtb *mysqlDatabase) Connect() *gorm.DB {
	once.Do(func() {
		// connnecting to database
		link := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", dtb.rootPassword, dtb.host, dtb.name)
		var erro error
		conn, erro = gorm.Open(mysql.Open(link), &gorm.Config{
			SkipDefaultTransaction: true,
		})
		if erro != nil {
			log.Fatalln("Error occured while connecting database", erro)
		}
		// set conection pool
		if dtb.connection != nil {
			dtbs, erro := conn.DB()
			if erro != nil {
				log.Fatalln("Error occured while setting connection pool", erro)
			}
			dtbs.SetMaxIdleConns(dtb.connection.idle)
			dtbs.SetMaxOpenConns(dtb.connection.maximum)
			dtbs.SetConnMaxLifetime(dtb.connection.lifetime)
		}
		// auto migrate schema
		if dtb.migrateTable {
			// drop all existed tables first
			migr := conn.Migrator()
			migr.DropTable(&ntt.ReportSchema{})
			// then migrate new tables
			conn.Set("gorm:table_options", "ENGINE=InnoDB")
			if erro := conn.AutoMigrate(&ntt.ReportSchema{}); erro != nil {
				log.Fatalln("Error occured while migrating table", erro)
			}
		}
	})
	return conn
}

// closing a connection to database
func (dtb *mysqlDatabase) Close() error {
	if conn == nil {
		return nil
	}
	dtbs, erro := conn.DB()
	if erro != nil {
		return erro
	}
	return dtbs.Close()
}

// pinging a connection to database
func (dtb *mysqlDatabase) Ping(ctxt context.Context) error {
	dtbs, erro := dtb.Connect().DB()
	if erro != nil {
		return erro
	}
	return dtbs.PingContext(ctxt)
}
