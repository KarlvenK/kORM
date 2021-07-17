package korm

import (
	"database/sql"
	"github.com/KarlvenK/korm/dialect"
	"github.com/KarlvenK/korm/log"
	"github.com/KarlvenK/korm/session"
)

//Engine is the main struct of korm, manages all db session and transactions.
type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

//NewEngine create a instance of Engine
//connect database and ping it to test whether it's alive
func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	//send a ping to make sure the database connection is alive
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.ErrorF("dialect %s Not Found", driver)
		return
	}
	e = &Engine{
		db:      db,
		dialect: dial,
	}
	log.Info("Connect database success")
	return
}

//Close closes database connection
func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

//NewSession creates a new session for next operation
func (e *Engine) NewSession() *session.Session {
	return session.New(e.db, e.dialect)
}
