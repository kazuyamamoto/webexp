package main

import (
	"log"
	"net/http"

	"github.com/rakyll/statik/fs"

	_ "github.com/kazuyamamoto/webexp/statik"
)

func main() {
	staticFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.HandlerFunc(handleRoot))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(staticFS)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRoot(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte(html))
}

const html = `
<!doctype html>
<html lang="ja">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>webexp</title>
    <link rel="stylesheet" href="public/bulma.css">
    <script src="public/mithril.js"></script>
    <!-- font awesome -->
    <script defer src="public/all.js"></script>
</head>
<body>
    <script src="public/app.js"></script>
</body>
</html>
`
