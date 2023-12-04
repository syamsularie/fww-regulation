package usecase

type PeduliLindungiUsecase struct{}

type PeduliLindungiExecutor interface {
	CheckPeduliLindungi(ktp string) (string, error)
}

func NewPeduliLindungiUsecase(peduliLindungiUsecase *PeduliLindungiUsecase) PeduliLindungiExecutor {
	return peduliLindungiUsecase
}

func (uc *PeduliLindungiUsecase) CheckPeduliLindungi(ktp string) (string, error) {
	switch ktp {
	case "3602041211870001":
		return "vaksin1", nil
	case "3702043111970002":
		return "vaksin2", nil
	case "3803013516970302":
		return "Belum Vaksin", nil
	default:
		return "KTP not found", nil
	}
}
