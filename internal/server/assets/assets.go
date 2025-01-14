package assets

import "embed"

//go:embed css
//go:embed js
//go:embed img
//go:embed xml
var FS embed.FS
