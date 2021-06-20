package database

import (
	"github.com/chensienyong/stocky/entity"
)

const TableDaily = "dailies"

func (pg *Postgres) FetchDailySeriesByStock(stockID int64) ([]entity.Daily, error) {
	var dailies []entity.Daily

	err := pg.Db.Table(TableDaily).Select("date,open,high,low,close,volume").Where("stock_id = ?", stockID).Order("date DESC").Find(&dailies)
	if err.Error != nil {
		return dailies, err.Error
	}

	return dailies, nil
}

func (pg *Postgres) InsertDailies(dailies []entity.Daily) ([]entity.Daily, error) {
	var newDailies []entity.Daily

	err := pg.Db.Table(TableDaily).Create(&dailies).Scan(&newDailies)
	if err.Error != nil {
		return newDailies, err.Error
	}

	return newDailies, nil
}

func (pg *Postgres) DeleteDailies(stockID int64) error {
	err := pg.Db.Table(TableDaily).Where("stock_id = ?", stockID).Delete(entity.Daily{})

	if err.Error != nil {
		return err.Error
	}

	return nil
}
