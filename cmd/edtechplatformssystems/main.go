// cmd/edtechplatformssystems/main.go
package main

import (
"flag"
"log"
"os"

"edtechplatformssystems/internal/edtechplatformssystems"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := edtechplatformssystems.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
