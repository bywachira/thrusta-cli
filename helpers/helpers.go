package helpers

import "time"

// SetInterval method to call function on set period
func SetInterval(handler func(), period time.Duration) {
	ticker := time.NewTicker(period * time.Second)

	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <- ticker.C:
				handler()
			case <- quit:
				ticker.Stop()
				return
			}			
		}
	}()
}