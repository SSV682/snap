package v1

import (
	"bytes"
	"errors"

	analyzer_v1 "github.com/SSV682/snap/protos/gen/go/analyzer"
	"github.com/go-playground/validator/v10"
)

const (
	ContentTypeApplicationJson = "application/json"

	datetimeStringFormat = "2024-02-06T10:00:00.731Z"
)

type MultiError []error

func (m MultiError) Error() string {
	b := bytes.NewBufferString("")

	for i, err := range m {
		if i > 0 {
			b.WriteString(", ")
		}

		b.WriteString(err.Error())
	}

	return b.String()
}

func validate(object any, validatorIns Validator) error {
	if err := validatorIns.Struct(object); err != nil {
		var validationErrs validator.ValidationErrors
		if errors.As(err, &validationErrs) {
			var resultErr MultiError

			for _, errItem := range validationErrs {
				resultErr = append(resultErr, errItem)
			}

			return resultErr
		}

		return err
	}

	return nil
}

func ListSettingsResponseFromPB(settings *analyzer_v1.ActualSettingsResponse) (response ListSettingsResponse) {
	data := settings.GetData()

	response.Data = make([]settingResponse, len(data))

	for i := range data {
		response.Data[i] = settingResponse{
			ID:             data[i].GetId(),
			StrategyName:   data[i].GetStrategyName(),
			Ticker:         data[i].GetTicker(),
			StartTime:      data[i].GetStart().AsTime(),
			EndTime:        data[i].GetEnd().AsTime(),
			StartInsideDay: data[i].GetStartInsideDay().AsTime(),
			EndInsideDay:   data[i].GetEndInsideDay().AsTime(),
		}
	}

	return response
}
