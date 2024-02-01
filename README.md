## Zemoga Test

### Running the server

Install the [Go compiler](https://go.dev/dl/)

Run in the terminal
`go run ./server.go`
then open [localhost:3000](http://localhost:3000)
in the browser.

### Running the tests

Install [Node.js runtime](https://nodejs.org/en/download/), then run `npm install` and `npm test`

### Folder structure

- `server.go` => a minimal static file server written in Go
- `index.test.ts` => all tests related to `index.html` page
- `assets/`
   - `img/` => contains all images
   - `data.json` => sample data for populating the page
- `css/main.css` => all styles for `index.html` page
- `index.html` => all html and javascript code for the application



