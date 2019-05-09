package lagou

import (
	"fmt"
	"testing"
)

func TestLagou(t *testing.T) {
	res, _ := Get("北京", "go", 1)
	fmt.Println(res)
}
