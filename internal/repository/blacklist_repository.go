package repository

import (
	"database/sql"
	"fww-regulation/config"

	"go.uber.org/zap"
)

type BlacklistRepository struct {
	DB   *sql.DB
	Base config.BaseDep
}

type BlacklistPersister interface {
	CheckBlacklist(ktp string) (bool, error)
}

func NewBlacklistRepository(blacklist BlacklistRepository) BlacklistPersister {
	return &blacklist
}

func (repo *BlacklistRepository) CheckBlacklist(ktp string) (bool, error) {
	var count int
	row := repo.DB.QueryRow("SELECT COUNT(*) FROM blacklist WHERE ktp = ?", ktp)
	if err := row.Scan(&count); err != nil {
		repo.Base.Logger.Error("failed to get blacklist", zap.Error(err))
		return false, err
	}

	if count > 0 {
		return false, nil
	}

	return true, nil
}
