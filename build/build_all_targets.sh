SCRIPT_DIR=$(dirname "$0")

# Build macos arm64 and macos amd64 executables and merge them into a universal macos executable
GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 go build -o ${SCRIPT_DIR}/../bin/macos_arm64/ffxivautocraft ${SCRIPT_DIR}/../ffxivautocraft.go
GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o ${SCRIPT_DIR}/../bin/macos_amd64/ffxivautocraft ${SCRIPT_DIR}/../ffxivautocraft.go
lipo -create -output ${SCRIPT_DIR}/../bin/macos_universal/ffxivautocraft ${SCRIPT_DIR}/../bin/macos_amd64/ffxivautocraft ${SCRIPT_DIR}/../bin/macos_arm64/ffxivautocraft

# Build linux amd64 executable
GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o ${SCRIPT_DIR}/../bin/linux_amd64/ffxivautocraft ${SCRIPT_DIR}/../ffxivautocraft.go

# Build windows amd64 executable
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -o ${SCRIPT_DIR}/../bin/windows_amd64/ffxivautocraft ${SCRIPT_DIR}/../ffxivautocraft.go

