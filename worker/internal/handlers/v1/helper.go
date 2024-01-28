package v1

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/valyala/fasthttp"
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

func newBacktestRequest(data []byte, validatorIns Validator) (settings *backtestRequest, err error) {
	if err = json.Unmarshal(data, &settings); err != nil {
		return nil, err
	}

	if err = validate(settings, validatorIns); err != nil {
		return nil, err
	}

	return settings, nil
}

func getTime(args *fasthttp.Args, argName string) (*time.Time, error) {
	startString := string(args.Peek(argName))
	if startString == "" {
		return nil, nil
	}

	atoi, err := strconv.Atoi(startString)
	if err != nil {
		return nil, fmt.Errorf("convert to int: %v", err)
	}

	parsedStartTime := time.Unix(int64(atoi), 0)
	return &parsedStartTime, nil
}
