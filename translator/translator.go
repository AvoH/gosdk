package translator

import "github.com/AvoH/gosdk/runtime"

func translate() string {
	return "Translated " + runtime.Hello()
}