package libValidate

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gvalid"
	"github.com/tiger1103/gfast/v3/internal/mounter"
	"strconv"
	"strings"
)

func init() {
	mounter.Mount(func(ctx context.Context, s *ghttp.Server) {
		Register()
	})
}

func Register() {
	gvalid.RegisterRule("integer-array", IntegerArray)
	gvalid.RegisterRule("float-array", FloatArray)
	gvalid.RegisterRule("date-array", DateArray)
	gvalid.RegisterRule("datetime-array", DatetimeArray)
}

func IntegerArray(ctx context.Context, RI gvalid.RuleFuncInput) error {
	if RI.Value.IsNil() {
		return nil
	}

	if RI.Value.IsSlice() {
		return gerror.New(RI.Message)
	}
	for _, v := range RI.Value.Array() {
		valueStr := strings.TrimSpace(gconv.String(v))
		if valueStr == "" {
			continue
		}
		if _, err := strconv.Atoi(valueStr); err != nil {
			continue
		}
		return gerror.New(RI.Message)
	}
	return nil
}

func FloatArray(ctx context.Context, RI gvalid.RuleFuncInput) error {
	if RI.Value.IsNil() {
		return nil
	}
	if !RI.Value.IsSlice() {
		return gerror.New(RI.Message)
	}
	for _, v := range RI.Value.Array() {
		valueStr := strings.TrimSpace(gconv.String(v))
		if valueStr == "" {
			continue
		}
		if _, err := strconv.ParseFloat(valueStr, 10); err == nil {
			continue
		}
		return gerror.New(RI.Message)
	}
	return nil
}

func DateArray(ctx context.Context, RI gvalid.RuleFuncInput) error {
	if RI.Value.IsNil() {
		return nil
	}
	if !RI.Value.IsSlice() {
		return gerror.New(RI.Message)
	}
	for _, v := range RI.Value.Array() {
		// support for time value, eg: gtime.Time/*gtime.Time, time.Time/*time.Time.
		if _, ok := v.(gtime.Time); ok {
			continue
		}
		valueStr := strings.TrimSpace(gconv.String(v))
		if valueStr == "" {
			continue
		}
		if gregex.IsMatchString(`\d{4}[\.\-\_/]{0,1}\d{2}[\.\-\_/]{0,1}\d{2}`, valueStr) {
			continue
		}
		return gerror.New(RI.Message)
	}
	return nil
}

func DatetimeArray(ctx context.Context, RI gvalid.RuleFuncInput) error {
	if RI.Value.IsNil() {
		return nil
	}
	if !RI.Value.IsSlice() {
		return gerror.New(RI.Message)
	}
	for _, v := range RI.Value.Array() {
		// support for time value, eg: gtime.Time/*gtime.Time, time.Time/*time.Time.
		if _, ok := v.(gtime.Time); ok {
			continue
		}
		valueStr := strings.TrimSpace(gconv.String(v))
		if valueStr == "" {
			continue
		}
		if _, err := gtime.StrToTimeFormat(valueStr, "Y-m-d H:i:s"); err == nil {
			continue
		}
		return gerror.New(RI.Message)
	}
	return nil
}
