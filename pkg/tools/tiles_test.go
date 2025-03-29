package tools_test

import (
	"fmt"
	"testing"

	"github.com/jrecuero/ebiplay/pkg/tools"
)

func TestMovableTiles(t *testing.T) {
	pos := tools.NewPoint(5, 5)
	result := tools.GetMovableTiles(pos, 2)
	fmt.Println(result)
}
