package somemodule

import "github.com/AvoH/gosdk/runtime"

func Convert() string {
	return "Converted " + runtime.Hello()
}
