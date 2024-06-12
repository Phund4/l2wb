package ntptime

import (
	"fmt"
	"testing"
)

func TestNtpTime(t *testing.T) {
	count := 5

	for i := 0; i < count; i++ {
		time, err := ntpTime()
		if err != nil {
			t.Errorf("Error in call func: %v", err)
		}
		fmt.Println(time)
	}
}
