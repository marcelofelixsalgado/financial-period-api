package period_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	. "marcelofelixsalgado/financial-period-api/api/controllers/period"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	createUseCaseMock "marcelofelixsalgado/financial-period-api/pkg/usecase/period/create/mocks"
	deleteUseCaseMock "marcelofelixsalgado/financial-period-api/pkg/usecase/period/delete/mocks"
	findUseCaseMock "marcelofelixsalgado/financial-period-api/pkg/usecase/period/find/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/list"
	listUseCaseMock "marcelofelixsalgado/financial-period-api/pkg/usecase/period/list/mocks"
	updateUseCaseMock "marcelofelixsalgado/financial-period-api/pkg/usecase/period/update/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"

	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListHandlerSuccess(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/v1/periods", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)

	createUseCaseMock := &createUseCaseMock.CreateUseCaseMock{}
	createUseCaseMock.On("Execute", mock.Anything).Return(nil)

	deleteUseCaseMock := &deleteUseCaseMock.DeleteUseCaseMock{}
	deleteUseCaseMock.On("Execute", mock.Anything).Return(nil)

	findUseCaseMock := &findUseCaseMock.FindUseCaseMock{}
	findUseCaseMock.On("Execute", mock.Anything).Return(nil)

	output := list.OutputListPeriodDto{
		Periods: []list.Period{
			{
				Id:        "123",
				Code:      "period code",
				Name:      "period name",
				Year:      2023,
				StartDate: "2023-01-01",
				EndDate:   "2023-02-01",
			},
		},
	}

	listUseCaseMock := &listUseCaseMock.ListUseCaseMock{}
	listUseCaseMock.On("Execute", list.InputListPeriodDto{}, []filter.FilterParameter{}).Return(output, status.Success, nil)

	updateUseCaseMock := &updateUseCaseMock.UpdateUseCaseMock{}
	updateUseCaseMock.On("Execute", mock.Anything).Return(nil)

	handler := NewPeriodHandler(createUseCaseMock, deleteUseCaseMock, findUseCaseMock, listUseCaseMock, updateUseCaseMock)

	handler.ListPeriods(c)

	if assert.NoError(t, handler.ListPeriods(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var periods []list.Period
		// outputPeriods := new(list.OutputListPeriodDto)

		bodyBytes, err1 := io.ReadAll(rec.Body)

		err2 := json.Unmarshal(bodyBytes, &periods)

		fmt.Println(err1)
		fmt.Println(err2)
		fmt.Println(periods)

		assert.Len(t, periods, 1)
	}
}
