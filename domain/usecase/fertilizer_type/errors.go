package fertilizer_type

import "github.com/anhvanhoa/service-core/domain/oops"

var (
	ErrInvalidID                       = oops.New("Id không hợp lệ")
	ErrFertilizerTypeNotFound          = oops.New("Loại phân bón không tồn tại")
	ErrFertilizerTypeNameRequired      = oops.New("Tên loại phân bón là bắt buộc")
	ErrInvalidFertilizerType           = oops.New("Loại phân bón không hợp lệ")
	ErrInvalidApplicationMethod        = oops.New("Phương pháp áp dụng không hợp lệ")
	ErrInvalidGrowthStage              = oops.New("Độ trưởng thành không hợp lệ")
	ErrInvalidEffectivenessRating      = oops.New("Đánh giá hiệu quả không hợp lệ")
	ErrInvalidNotes                    = oops.New("Ghi chú không hợp lệ")
	ErrFertilizerTypeNameAlreadyExists = oops.New("Tên loại phân bón đã tồn tại")
	ErrBatchNumberAlreadyExists        = oops.New("Số lô phân bón đã tồn tại")
	ErrInvalidCreatedBy                = oops.New("Người tạo không hợp lệ")
)
