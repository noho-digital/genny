package unexported

import (
	"fmt"

	"github.com/noho-digital/genny/generic"
)

type secret generic.Type

func secretInspect(s secret) string {
	return fmt.Sprintf("%#v", s)
}
