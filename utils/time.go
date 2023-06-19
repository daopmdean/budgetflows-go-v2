package utils

import "time"

var TimeZoneVN = time.FixedZone("VN", +7*60*60)

func TimeToMonthlyVersion(t time.Time) string {
	return t.Format("2006_01")
}
