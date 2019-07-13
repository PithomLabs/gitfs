// An example that locally serves files from github.com/golang/go/doc
package main

import (
	"context"
	"log"
	"os"

	"github.com/posener/gitfs"
	"github.com/posener/gitfs/fsutil"
)

// Add debug mode environment variable. When running with `LOCAL_DEBUG=.`, the
// local git repository will be used instead of the remote github.
var localDebug = os.Getenv("LOCAL_DEBUG")

func main() {
	ctx := context.Background()
	fs, err := gitfs.New(ctx, "github.com/posener/gitfs/examples/templates", gitfs.OptLocal(localDebug))
	if err != nil {
		log.Fatalf("Failed initializing git filesystem: %s.", err)
	}

	tmpls, err := fsutil.TmplParseGlob(fs, nil, "*.gotmpl")
	if err != nil {
		log.Fatalf("Failed parsing templates.")
	}
	tmpls.ExecuteTemplate(os.Stdout, "tmpl1.gotmpl", "Foo")
}
