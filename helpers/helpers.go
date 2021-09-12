package helpers

import (
	"strings"
	"time"
)

type Helpers struct {
}

func (h *Helpers) ConcatString(parent string, child string) string {
	var sb strings.Builder

	sb.WriteString(parent)

	sb.WriteString(child)

	return sb.String()
}

func (h *Helpers) SplitString(value string, char string) []string {
	s := strings.Split(value, char)

	return s
}

func (h *Helpers) ContainsString(value string, char string) bool {
	return strings.Contains(value, char)
}

func (h *Helpers) SetInterval(doStuff func(), seconds float32) interface{} {
	ticker := time.NewTicker(time.Duration(seconds) * time.Second)

	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				doStuff()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return quit
}
