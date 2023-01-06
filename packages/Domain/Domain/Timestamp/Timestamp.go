package Timestamp

import (
  "time"
)

// Returns timestamp in ISO8601 format
func GetNowTimeByISO8601Format() string {
  return time.Now().Format(time.RFC3339) // RFC3339 = ISO8601
}
