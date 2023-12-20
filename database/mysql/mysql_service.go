package mysqlservice

import (
	"database/sql"
	"sync"
	"time"
)

type MysqlDBService struct {
	Config     *MysqlConfig
	dbinstance *sql.DB
	lock       *sync.Mutex
}

func (dbs *MysqlDBService) createNewDBInstance() (*sql.DB, error) {
	mysql_config_fordb := GetMysqlConfig(dbs.Config)
	dbinst, err := sql.Open("mysql", mysql_config_fordb.FormatDSN())
	if err != nil {
		return nil, err
	}
	dbinst.SetMaxOpenConns(dbs.Config.ConnMaxOpen)
	dbinst.SetConnMaxIdleTime(time.Millisecond * time.Duration(dbs.Config.ConnMaxIdleTime))
	dbinst.SetMaxIdleConns(dbs.Config.ConnMaxIdleConns)

	err = dbinst.Ping()
	if err != nil {
		dbinst.Close()
		return nil, err
	}

	return dbinst, nil
}

func (dbs *MysqlDBService) GetDBInstance() (*sql.DB, error) {
	dbs.lock.Lock()
	defer dbs.lock.Unlock()

	if dbs.dbinstance != nil {
		return dbs.dbinstance, nil
	}

	dbinst, err := dbs.createNewDBInstance()
	if err != nil {
		return nil, err
	}
	dbs.dbinstance = dbinst
	return dbs.dbinstance, nil
}

func (dbs *MysqlDBService) ClearDBInstance() (bool, error) {
	dbs.lock.Lock()
	defer dbs.lock.Unlock()

	if dbs.dbinstance == nil {
		return true, nil
	}

	err := dbs.dbinstance.Close()
	dbs.dbinstance = nil
	return true, err
}

func NewMysqlDBService(mysql_config *MysqlConfig) *MysqlDBService {
	return &MysqlDBService{
		Config:     mysql_config,
		dbinstance: nil,
		lock:       &sync.Mutex{},
	}
}
