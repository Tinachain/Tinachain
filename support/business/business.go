// support project business.go
package business

import (
	"strconv"
	"strings"
	"time"
)

type Business struct {
}

func New() *Business {

	downloadPtr := new(Download)
	downloadPtr.loadConfig()
	return downloadPtr
}
