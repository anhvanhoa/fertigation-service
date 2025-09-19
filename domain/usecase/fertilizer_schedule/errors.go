package fertilizer_schedule

import "github.com/anhvanhoa/service-core/domain/oops"

var (
	ErrInvalidPlantingCycleID     = oops.New("Id vòng trồng không hợp lệ")
	ErrInvalidFertilizerTypeID    = oops.New("Id loại phân bón không hợp lệ")
	ErrInvalidCreatedBy           = oops.New("Người tạo không hợp lệ")
	ErrInvalidApplicationDate     = oops.New("Ngày áp dụng không hợp lệ")
	ErrInvalidDosage              = oops.New("Liều lượng không hợp lệ")
	ErrInvalidUnit                = oops.New("Đơn vị không hợp lệ")
	ErrInvalidApplicationMethod   = oops.New("Phương pháp áp dụng không hợp lệ")
	ErrInvalidGrowthStage         = oops.New("Độ trưởng thành không hợp lệ")
	ErrInvalidWeatherConditions   = oops.New("Điều kiện thời tiết không hợp lệ")
	ErrInvalidSoilConditions      = oops.New("Điều kiện đất không hợp lệ")
	ErrInvalidIsCompleted         = oops.New("Đã hoàn thành không hợp lệ")
	ErrInvalidCompletedDate       = oops.New("Ngày hoàn thành không hợp lệ")
	ErrInvalidActualDosage        = oops.New("Liều lượng thực tế không hợp lệ")
	ErrInvalidEffectivenessRating = oops.New("Đánh giá hiệu quả không hợp lệ")
	ErrInvalidNotes               = oops.New("Ghi chú không hợp lệ")
	ErrInvalidID                  = oops.New("Id lịch bón phân không hợp lệ")
	ErrFertilizerScheduleNotFound = oops.New("Lịch bón phân không tồn tại")
)
