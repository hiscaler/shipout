package shipout

import (
	"github.com/go-resty/resty/v2"
	"log"
)

type service struct {
	debug      bool          // Is debug mode
	logger     *log.Logger   // Log
	httpClient *resty.Client // HTTP client
}

type wmsServices struct {
}

// OMS API Services
type omsServices struct {
	BaseInfo baseInfoService
}
