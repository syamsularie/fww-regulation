package usecase

import "fww-regulation/internal/repository"

type BlacklistUsecase struct {
	BlacklistRepo repository.BlacklistPersister
}

type BlacklistExecutor interface {
	CheckBlacklist(ktp string) (bool, error)
}

func NewBlacklistUsecase(blacklistUsecase *BlacklistUsecase) BlacklistExecutor {
	return blacklistUsecase
}

func (uc *BlacklistUsecase) CheckBlacklist(ktp string) (bool, error) {
	return uc.BlacklistRepo.CheckBlacklist(ktp)
}
