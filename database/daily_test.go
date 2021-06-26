package database_test

import (
	"github.com/chensienyong/stocky/customerror"
	"github.com/chensienyong/stocky/database"
	"github.com/chensienyong/stocky/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestPostgres_FetchDailySeriesByStock_Success(t *testing.T) {
	exp := []entity.Daily{
		{
			ID: 1, StockID: 1, Date: "2018-04-05", Open: 3.4, High: 5.6, Low: 1.2, Close: 4.5, Volume: 100,
		},
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "stock_id", "date", "open", "high", "low", "close", "volume"})
	rows.AddRow("1", "1", "2018-04-05", "3.4", "5.6", "1.2", "4.5", "100")

	mock.ExpectQuery("^SELECT (.+)").WillReturnRows(rows)

	pg := database.Postgres{}
	pg.Db, err = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	res, err := pg.FetchDailySeriesByStock(1)

	assert.Nil(t, err)
	assert.Equal(t, exp, res)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPostgres_FetchDailySeriesByStock_Fail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("^SELECT (.+)").WillReturnError(customerror.DBError)

	pg := database.Postgres{}
	pg.Db, err = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	_, err = pg.FetchDailySeriesByStock(1)

	assert.NotNil(t, err)
	assert.Equal(t, customerror.DBError, err)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPostgres_InsertDailies_Success(t *testing.T) {
	data := []entity.Daily{
		{
			StockID: 1, Date: "2018-04-05", Open: 3.4, High: 5.6, Low: 1.2, Close: 4.5, Volume: 100,
		},
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "dailies" ("stock_id","date","open","high","low","close","volume") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`)).WithArgs(data[0].StockID, data[0].Date, data[0].Open, data[0].High, data[0].Low, data[0].Close, data[0].Volume).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	pg := database.Postgres{}
	pg.Db, err = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	err = pg.InsertDailies(data)

	assert.Nil(t, err)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPostgres_InsertDailies_Fail(t *testing.T) {
	data := []entity.Daily{
		{
			StockID: 1, Date: "2018-04-05", Open: 3.4, High: 5.6, Low: 1.2, Close: 4.5, Volume: 100,
		},
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "dailies" ("stock_id","date","open","high","low","close","volume") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`)).WithArgs(data[0].StockID, data[0].Date, data[0].Open, data[0].High, data[0].Low, data[0].Close, data[0].Volume).
		WillReturnError(customerror.DBError)

	pg := database.Postgres{}
	pg.Db, err = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	err = pg.InsertDailies(data)

	assert.Equal(t, customerror.DBError, err)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPostgres_DeleteDailies_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("^DELETE FROM \"dailies\"(.+)").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	pg := database.Postgres{}
	pg.Db, err = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	err = pg.DeleteDailies(1)

	assert.Nil(t, err)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPostgres_DeleteDailies_Fail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("^DELETE FROM \"dailies\"(.+)").WillReturnError(customerror.RecordNotFound)
	mock.ExpectRollback()

	pg := database.Postgres{}
	pg.Db, err = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	err = pg.DeleteDailies(1)

	assert.NotNil(t, err)
	assert.Equal(t, customerror.RecordNotFound, err)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
