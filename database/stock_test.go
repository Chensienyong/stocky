package database_test

import (
	"github.com/chensienyong/stocky/database"
	"github.com/chensienyong/stocky/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestPostgres_GetOrCreateStock_Success(t *testing.T) {
	stockSymbol := "GGRM"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "stock_symbol"})
	rows.AddRow("1", stockSymbol)

	mock.ExpectQuery("^SELECT (.+)").WillReturnRows(rows)

	pg := database.Postgres{}
	pg.Db, err = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	res, err := pg.GetOrCreateStock(stockSymbol)

	assert.Nil(t, err)
	assert.Equal(t, stockSymbol, res.StockSymbol)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPostgres_GetOrCreateStock_NoRecordAndCreate(t *testing.T) {
	stockSymbol := "GGRM"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "stock_symbol"})
	rows.AddRow("1", stockSymbol)

	mock.ExpectQuery("^SELECT (.+)").WillReturnError(gorm.ErrRecordNotFound)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "stocks" ("stock_symbol") VALUES ($1) RETURNING "id"`)).WithArgs(stockSymbol).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()
	mock.ExpectQuery("^SELECT (.+)").WillReturnRows(rows)

	pg := database.Postgres{}
	pg.Db, err = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	res, err := pg.GetOrCreateStock(stockSymbol)

	assert.Nil(t, err)
	assert.Equal(t, stockSymbol, res.StockSymbol)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPostgres_CreateStock_Success(t *testing.T) {
	stock := entity.Stock{
		StockSymbol: "GGRM",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "stock_symbol"})
	rows.AddRow("1", "GGRM")

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "stocks" ("stock_symbol") VALUES ($1) RETURNING "id"`)).WithArgs(stock.StockSymbol).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()
	mock.ExpectQuery("^SELECT (.+)").WillReturnRows(rows)

	pg := database.Postgres{}
	pg.Db, err = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	res, err := pg.CreateStock(stock)

	assert.Nil(t, err)
	assert.Equal(t, stock.StockSymbol, res.StockSymbol)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
