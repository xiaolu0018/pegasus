package organization

import (
	"errors"
	"sort"
	"strings"
)

var (
	ErrOffDayInvalid      error = errors.New("offday too much")
	ErrAvoidNumberInvalid error = errors.New("avoid number invalid")
	ErrCapacityInvalid    error = errors.New("param capacity invalid")
	ErrWarnNumInvalid     error = errors.New("param warnnum invalid")
	ErrSaleCodeNotFound   error = errors.New("param sale code not found")
)

func (c *Config_Basic) Validate() error {
	if c.Capacity > 1000 || c.Capacity < 0 {
		return ErrCapacityInvalid

	}
	if c.WarnNum > c.Capacity {
		return ErrWarnNumInvalid
	}

	if len(c.OffDays) > 50 {
		return ErrOffDayInvalid
	}

	sort.Slice(c.AvoidNumbers, func(i, j int) bool {
		return c.AvoidNumbers[i] < c.AvoidNumbers[j]
	})
	if len(c.AvoidNumbers) > 0 {
		if c.AvoidNumbers[len(c.AvoidNumbers)-1] > 1000 || c.AvoidNumbers[0] < 0 {
			return ErrAvoidNumberInvalid
		}
	}

	for i := range c.AvoidNumbers {
		if i == 0 {
			continue
		}
		if c.AvoidNumbers[i] == c.AvoidNumbers[i-1] {
			return ErrAvoidNumberInvalid
		}
	}

	return nil
}

func (c *Config_Special) Validate() error {
	if len(strings.TrimSpace(c.CheckupCode)) == 0 {
		return ErrSaleCodeNotFound
	}
	if c.Capacity < 0 || c.Capacity > 1000 {
		return ErrCapacityInvalid
	}
	return nil
}
