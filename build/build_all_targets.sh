SCRIPT_DIR=$(dirname "$0")

# Build macos arm64 and macos amd64 executables and merge them into a universal macos executable
GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 go build -o ${SCRIPT_DIR}/../bin/ffxivautocraft-macos-arm64/ffxivautocraft ${SCRIPT_DIR}/../ffxivautocraft.go
GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o ${SCRIPT_DIR}/../bin/ffxivautocraft-macos-amd64/ffxivautocraft ${SCRIPT_DIR}/../ffxivautocraft.go
lipo -create -output ${SCRIPT_DIR}/../bin/ffxiv-autocraft-macos-universal/ffxivautocraft ${SCRIPT_DIR}/../bin/ffxivautocraft-macos-amd64/ffxivautocraft ${SCRIPT_DIR}/../bin/ffxivautocraft-macos-arm64/ffxivautocraft

# Build linux amd64 executable
GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o ${SCRIPT_DIR}/../bin/ffxivautocraft-linux-amd64/ffxivautocraft ${SCRIPT_DIR}/../ffxivautocraft.go

# Build windows amd64 executable
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -o ${SCRIPT_DIR}/../bin/ffxivautocraft-windows-amd64/ffxivautocraft.exe ${SCRIPT_DIR}/../ffxivautocraft.go

# compress the executables into archives
zip -r deploy/ffxivautocraft-macos-amd64.zip bin/ffxivautocraft-macos-amd64/ffxivautocraft
zip -r deploy/ffxivautocraft-macos-arm64.zip bin/ffxivautocraft-macos-arm64/ffxivautocraft
zip -r deploy/ffxivautocraft-macos-universal.zip bin/ffxivautocraft-macos-universal/ffxivautocraft
tar -cvzf deploy/ffxivautocraft-linux-amd64.tar.gz bin/ffxivautocraft-linux-amd64/ffxivautocraft
zip -r deploy/ffxivautocraft-windows-amd64.zip bin/ffxivautocraft-windows-amd64/ffxivautocraft.exe

