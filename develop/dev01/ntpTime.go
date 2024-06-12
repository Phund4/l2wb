package ntptime

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func ntpTime() (time.Time, error) {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		_, writeStrErr := io.WriteString(os.Stderr, err.Error())
		if writeStrErr != nil {
			fmt.Printf("Error in write string in stderr: %v", writeStrErr)
		}
		return time, err
	}
	return time, nil
}
