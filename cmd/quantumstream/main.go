// cmd/quantumstream/main.go
package main

import (
"flag"
"log"
"os"

"quantumstream/internal/quantumstream"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := quantumstream.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
