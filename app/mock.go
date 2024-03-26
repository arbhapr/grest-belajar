package app

import (
	"gorm.io/gorm"
	"grest.dev/grest"
)

// add your mock service here, for example :
//
//	// original func (add IS_USE_MOCK_SERVICE_ABC and IS_USE_MOCK_SUCCESS to config)
//	func (*UseCaseHandler) CallServiceABC(ctx *Ctx, param RequestServiceABC) (ResponseServiceABC, error) {
//		if IS_USE_MOCK_SERVICE_ABC {
//			return Mock().CallServiceABC(ctx, param)
//		}
//		var err error
//		res := ResponseServiceABC{}
//		// do something here
//		return res, err
//	}
//
//	// mock func
//	func (*mockHandler) CallServiceABC(ctx *Ctx, param RequestServiceABC) (ResponseServiceABC, error) {
//		if IS_USE_MOCK_SUCCESS  {
//			// return success dummy
//		}
//		return ResponseServiceABC{}, NewError(http.StatusInternalServerError, "Something wrong")
//	}
func Mock() *mockUtil {
	if mock == nil {
		mock = &mockUtil{}
	}
	return mock
}

var mock *mockUtil

type mockUtil struct{}

func (*mockUtil) DB() (*gorm.DB, error) {
	mockDB, _, mockErr := grest.NewMockDB()
	return mockDB, mockErr
}
