package utils

import (
	"fmt"

	"github.com/kr/pretty"
)

func Dump(v ...interface{}) {
	fmt.Printf("%# v\n", pretty.Formatter(v))

}

func Sdump(v ...interface{}) string {
	return fmt.Sprintf("%# v\n", pretty.Formatter(v))
}
