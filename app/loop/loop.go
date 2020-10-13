package loop

import (
	"fmt"
	"math"
)

func Loop() string {
	x := 0.0001
	for i := 0; i <= 1000000; i++ {
		x += math.Sqrt(x)
	}
	return fmt.Sprintf("Code.education Rocks!")
}
