package pages

import "embed"

// template files are bundled with binary
// for worry free deployment that needs to copy a single file

//go:embed templates
var templatesFS embed.FS
