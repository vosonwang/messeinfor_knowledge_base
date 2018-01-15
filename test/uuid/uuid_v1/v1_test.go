package uuid_v1

import (
	"testing"
	"github.com/satori/go.uuid"
	"fmt"
)

/*0.011*/
/*0.011*/
/*0.010*/
/*v1和v2的性能几乎一致*/
func TestV1(t *testing.T) {

	var a = 0

	for a < 100 {
		u1 := uuid.NewV1()
		a++
		fmt.Printf("UUIDv1: %s\n", u1)
	}

}
