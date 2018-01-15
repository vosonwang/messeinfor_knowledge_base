package uuid_v4

import (
	"testing"
	"github.com/satori/go.uuid"
	"fmt"
)

/*0.014*/
/*0.013*/
/*0.011*/
/*0.011*/
func TestV4(t *testing.T) {

	var a = 0

	for a < 100 {
		u4 := uuid.NewV4()
		a++
		fmt.Printf("UUIDv4: %s\n", u4)
	}

}
