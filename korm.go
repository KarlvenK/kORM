package korm

import (
	"database/sql"
	"github.com/KarlvenK/korm/log"
	"github.com/KarlvenK/korm/session"
)

//Engine is the main struct of krom, manages all db session and transactions.
type Engine struct {
	db *sql.DB
}

//NewEngine create a instance of Engine
//connect database and ping it to test whether it's alive
func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	//send a ping to make sure the database connnection is alive
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	e = &Engine{
		db: db,
	}
	log.Info("Connnect database success")
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
	return session.New(e.db)
}
