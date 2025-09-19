package irrigation_log

import "github.com/anhvanhoa/service-core/domain/oops"

var (
	ErrInvalidID                   = oops.New("Id lịch tưới không hợp lệ")
	ErrIrrigationLogNotFound       = oops.New("Lịch tưới không tồn tại")
	ErrInvalidIrrigationScheduleID = oops.New("Id lịch tưới không hợp lệ")
	ErrInvalidDeviceID             = oops.New("Id thiết bị không hợp lệ")
	ErrInvalidStatus               = oops.New("Trạng thái không hợp lệ")
	ErrInvalidCreatedBy            = oops.New("Người tạo không hợp lệ")
)
