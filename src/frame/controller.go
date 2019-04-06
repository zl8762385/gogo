// controller鸡肋
package frame

import (
	"net/http"
)

type Controller struct {
	Response http.ResponseWriter
	Request  *http.Request
}
