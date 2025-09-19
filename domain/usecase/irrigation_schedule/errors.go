package irrigation_schedule

import "github.com/anhvanhoa/service-core/domain/oops"

var (
	ErrInvalidID                  = oops.New("Id lịch tưới không hợp lệ")
	ErrIrrigationScheduleNotFound = oops.New("Lịch tưới không tồn tại")
	ErrScheduleNameRequired       = oops.New("Tên lịch tưới là bắt buộc")
	ErrInvalidGrowingZoneID       = oops.New("Id vùng trồng không hợp lệ")
	ErrInvalidPlantingCycleID     = oops.New("Id vòng trồng không hợp lệ")
	ErrInvalidCreatedBy           = oops.New("Người tạo không hợp lệ")
	ErrScheduleNameAlreadyExists  = oops.New("Tên lịch tưới đã tồn tại")
	ErrInvalidTimeRange           = oops.New("Khoảng thời gian không hợp lệ")
	ErrInvalidTimeFormat          = oops.New("Định dạng thời gian không hợp lệ")
	ErrInvalidFrequency           = oops.New("Tần suất không hợp lệ")
)
