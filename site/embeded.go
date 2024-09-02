package site

import "embed"

// static files are bundled into separate FS
// because full content of that embed.FS is available
// under http://127.0.0.1:8090/static/static/public/

//go:embed static
var staticFilesFS embed.FS
