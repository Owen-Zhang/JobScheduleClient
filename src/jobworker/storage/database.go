package storage

import (
	"sync"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

const (
	DATABASENAME = "jobschedule"
	TASK		 = "task"
	LOGS         = "logs"
)

type M bson.M

type DataStorageArgs struct {
	Hosts    string
	DBName   string
	Auth     map[string]string
}

type DataStorage struct {
	mutex  *sync.Mutex
	mdb    *mgo.Database
	ss     *mgo.Session
	cc map[string]*mgo.Collection
}

func NewDataStorage(arg *DataStorageArgs) (*DataStorage, error) {
	session, err := mgo.Dial(arg.Hosts)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Strong, true)
	dbname := strings.TrimSpace(arg.DBName)
	if dbname == "" {
		dbname = DATABASENAME
	}

	db := session.DB(dbname)
	if len(arg.Auth) > 0 {
		if errdb := db.Login(arg.Auth["user"], arg.Auth["password"]); errdb != nil {
			return  nil, errdb
		}
	}

	collection := make(map[string]*mgo.Collection)
	collection[TASK] = db.C(TASK)
	collection[LOGS] = db.C(LOGS)

	datastorage := &DataStorage{
		mutex	: new(sync.Mutex),
		mdb		: db,
		ss     : session,
		cc      : collection,
	}
	return datastorage, nil
}

func (this *DataStorage) Close() {
	if this.ss != nil {
		this.ss.Close()
	}
}
