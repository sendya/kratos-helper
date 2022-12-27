package confbase

import (
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/config"
	"google.golang.org/grpc/encoding"
)

func Decoder(src *config.KeyValue, target map[string]interface{}) error {
	// src.Format is filename~ typo temp file, ignore.
	if strings.HasSuffix(src.Format, "~") || strings.HasSuffix(src.Key, "~") {
		return nil
	}
	codec := encoding.GetCodec(src.Format)
	if codec == nil {
		return fmt.Errorf("unsupported key: %s format: %s", src.Key, src.Format)
	}
	// src.Key is filename, you can also use regexp to filter it.
	if strings.HasPrefix(src.Key, ".") {
		// do nothing, skiped file named src.Key
		return nil
	}
	return codec.Unmarshal(src.Value, &target)
}
