package usecase

type DukcapilUsecase struct {
}

type DukcapilExecutor interface {
	CheckDukcapilByKTP(ktp string) (string, error)
}

func NewDukcapilUsecase(dukcapilUsecase *DukcapilUsecase) DukcapilExecutor {
	return dukcapilUsecase
}

func (uc *DukcapilUsecase) CheckDukcapilByKTP(ktp string) (string, error) {
	switch ktp {
	case "3602041211870001":
		return "valid", nil
	case "3702043111970002":
		return "not valid", nil
	case "3803013516970302":
		return "valid", nil
	default:
		return "ktp not found", nil
	}
}
