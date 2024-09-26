package entity

import (
	"testing"
	"time"
)

func TestGetVersionVN(t *testing.T) {
	testCases := []struct {
		partitionType PartitionType
		tm            time.Time
		want          string
	}{
		{
			PartitionTypeValue.Daily,
			time.Date(2020, 01, 01, 22, 0, 0, 0, time.UTC),
			"2020_01_02",
		},
		{
			PartitionTypeValue.Monthly,
			time.Date(2020, 01, 31, 21, 0, 0, 0, time.UTC),
			"2020_02",
		},
		{
			PartitionTypeValue.Yearly,
			time.Date(2020, 12, 31, 20, 0, 0, 0, time.UTC),
			"2021",
		},
	}

	for _, tc := range testCases {

		v, err := getVersionVN(tc.partitionType, tc.tm)
		if err != nil {
			t.Errorf("getVersion() error = %v", err)
		}

		if v != tc.want {
			t.Errorf("getVersion() = %v, want %v", v, tc.want)
		}
	}
}

func TestGetVersion(t *testing.T) {
	testCases := []struct {
		partitionType PartitionType
		tm            time.Time
		want          string
	}{
		{
			PartitionTypeValue.Daily,
			time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
			"2020_01_01",
		},
		{
			PartitionTypeValue.Monthly,
			time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
			"2020_01",
		},
		{
			PartitionTypeValue.Yearly,
			time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
			"2020",
		},
	}

	for _, tc := range testCases {

		v, err := getVersion(tc.partitionType, tc.tm)
		if err != nil {
			t.Errorf("getVersion() error = %v", err)
		}

		if v != tc.want {
			t.Errorf("getVersion() = %v, want %v", v, tc.want)
		}
	}
}
