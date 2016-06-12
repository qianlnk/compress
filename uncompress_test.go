package compress

import (
	"testing"
)

func TestUncompress(t *testing.T) {
	Uncompress("helloXzj.tar.gz", ".")
}
