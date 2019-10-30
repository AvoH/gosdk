package main

import (
	"fmt"
	"github.com/AvoH/gosdk/converter"
	"github.com/AvoH/gosdk/translator"
)
func main() {
	fmt.Println(converter.Convert())
	fmt.Println(translator.Translate())
}
