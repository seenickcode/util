package util

import (
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/Sirupsen/logrus"
)

func LogIfErr(err error) error {
	if err != nil {
		return logAndReportErrWithMsg(err, "", nil)
	}
	return nil
}

func LogIfErrWithMsg(err error, fmtdStr string, fmtdStrArgs ...interface{}) error {
	if err != nil {
		return logAndReportErrWithMsg(err, fmtdStr, fmtdStrArgs...)
	}
	return nil
}

func LogErr(err error) error {
	return logAndReportErrWithMsg(err, "", nil)
}

func LogErrMsg(fmtdStr string, fmtdStrArgs ...interface{}) error {
	return logAndReportErrWithMsg(nil, fmtdStr, fmtdStrArgs...)
}

// log a descriptive err object and custom fmtdStr, then report to bugsnag
func logAndReportErrWithMsg(err error, fmtdStr string, fmtdStrArgs ...interface{}) error {
	// construct
	parts := []string{}
	if len(fmtdStr) > 0 {
		parts = append(parts, fmt.Sprintf(fmtdStr, fmtdStrArgs...))
	}
	if err != nil {
		parts = append(parts, err.Error())
	}
	wrappedErr := strings.Join(parts, ": ")
	wrappedErr = wrappedErr + fmt.Sprintf("\n%s", debug.Stack())
	fullErr := fmt.Errorf(wrappedErr)

	// log
	logrus.Error(fullErr)

	return fullErr
}
