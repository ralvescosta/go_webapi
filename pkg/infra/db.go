package infra

type Db interface {
	Create() error
	GetOne() error
	GetAll() error
	Update() error
	Delete() error
}

type PostgreDb struct{}

type dbConn struct {
	db *PostgreDb
}

func (conn dbConn) Create() error {
	return nil
}
func (conn dbConn) GetOne() error {
	return nil
}
func (conn dbConn) GetAll() error {
	return nil
}
func (conn dbConn) Update() error {
	return nil
}
func (conn dbConn) Delete() error {
	return nil
}

func NewDbContext(db *PostgreDb) Db {
	return &dbConn{
		db: db,
	}
}
