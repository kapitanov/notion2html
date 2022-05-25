go build -o ./notion2html.exe ./cmd/notion2html
if ($LASTEXITCODE -ne 0) {
    exit $LASTEXITCODE
}

./notion2html.exe $args
if ($LASTEXITCODE -ne 0) {
    exit $LASTEXITCODE
}