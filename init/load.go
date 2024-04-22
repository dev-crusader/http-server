package init

import (
	"github.com/magiconair/properties"
)

var (
	appProperties = make(map[string]any)
)

const (
	propertyKey = "propKeys"
)

func GetAll() *properties.Properties {
	p, _ := appProperties[propertyKey].(*properties.Properties)
	return p
}
