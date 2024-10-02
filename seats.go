package main

import "errors"

var SeatsCount int

const (
	DEFAULT_SEATS_COUNT = 4
)

func init() {
	SeatsCount = DEFAULT_SEATS_COUNT
}

func DegreaseSeatsCount() error {
	// 席数が0の場合はエラーを返す
	if SeatsCount == 0 {
		return errors.New("Seats count is already 0")
	}
	SeatsCount--
	return nil
}

func IncreaseSeatsCount() error {
	// 席数が最大の場合はエラーを返す
	if SeatsCount+1 > DEFAULT_SEATS_COUNT {
		return errors.New("Seats count is already at maximum")
	}
	SeatsCount++
	return nil
}
