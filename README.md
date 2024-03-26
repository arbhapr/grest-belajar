# grest-belajar

## Getting Started
1. Make sure you have [Go](https://go.dev) installed.
2. Clone the repo
```bash
git clone https://grest-belajar.git
```
3. Go to the directory and run go mod tidy to add missing requirements and to drop unused requirements
```bash
cd grest-belajar && go mod tidy
```
3. Setup your .env file
```bash
cp .env-example .env && vi .env
```
4. Start
```bash
go run main.go
```

## Code Documentation
1. Install godoc
```bash
go install golang.org/x/tools/cmd/godoc@latest
```
2. Run godoc in grest-belajar directory
```bash
godoc -http=:6060
```
3. Open http://localhost:6060/pkg/grest-belajar in browser

## Open API Documentation
1. Update your open api documentation
```bash
go run main.go update
```
2. Start
```bash
go run main.go
```
3. Open http://localhost:4001/api/docs in browser

## Test
1. Make sure you have db with name `main_test.db` with credentials same as DB_XXX
2. Test all with verbose output that lists all of the tests and their results.
```bash
ENV_FILE=$(pwd)/.env go test ./... -v
```
3. Test all with benchmark.
```bash
ENV_FILE=$(pwd)/.env go test ./... -bench=.
```

## Build for production
1. Compile packages and dependencies
```bash
go build -o grest-belajar main.go
```
2. Setup .env file for production
```bash
cp .env-example .env && vi .env
```
3. Run executable file with systemd, supervisor, pm2 or other process manager
```bash
./grest-belajar
```