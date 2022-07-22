package mysqlstorage

type BaseStorage struct {
	dbConnection string // represent for a connection model in real use case
}

func NewBaseStorage(dbConnection string) *BaseStorage {
	return &BaseStorage{dbConnection: dbConnection}
}

func (s *BaseStorage) GetDB() string {
	return s.dbConnection
}
