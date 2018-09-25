package usage

import "github.com/octomarat/meta"

func _() {
	t := UsageType{}
	meta.NotNil(t.Field)
}
