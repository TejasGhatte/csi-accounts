package helpers

import (
	"fmt"

	"csi-accounts/internal/initializers"
)

func LogDatabaseError(customString string, err error, path string) {
	if err == nil {
		err = fmt.Errorf("no error description provided")
	}

	if appError, ok := err.(*AppError); ok {
		initializers.Loggers.Warn.
			Str("Message", appError.Message).
			Str("Path", path).
			Err(appError.Err).
			Msg(customString)
	} else {
		initializers.Loggers.Warn.
			Str("Message", err.Error()).
			Str("Path", path).
			Err(err).
			Msg(customString)
	}
}

func LogServerError(customString string, err error, path string) {
	if err == nil {
		err = fmt.Errorf("no error description provided")
	}

	if appError, ok := err.(*AppError); ok {
		initializers.Loggers.Error.
			Str("Message", appError.Message).
			Str("Path", path).
			Err(appError.Err).
			Msg(customString)
	} else {
		initializers.Loggers.Error.
			Str("Message", err.Error()).
			Str("Path", path).
			Err(err).
			Msg(customString)
	}
}
