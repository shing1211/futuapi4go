#!/usr/bin/env bash
set -euo pipefail

PROJECT="github.com/shing1211/futuapi4go"
PROTO_DIR="api/proto"
OUT_DIR="pkg/pb"

if ! command -v protoc &> /dev/null; then
    echo "Error: protoc is not installed. Install it first:"
    echo "  macOS: brew install protobuf"
    echo "  Linux: sudo apt install protobuf-compiler"
    exit 1
fi

echo "Installing protoc-gen-go and protoc-gen-go-grpc if missing..."
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

echo "Clearing old generated files in $OUT_DIR..."
rm -rf "$OUT_DIR"
mkdir -p "$OUT_DIR"

echo "Generating all protobuf files..."
PROTO_FILES=("$PROTO_DIR"/*.proto)
TOTAL=${#PROTO_FILES[@]}
echo "Found $TOTAL proto files"

# Generate all at once: --go_out=. outputs to ./pkg/pb/<go_package_relative_path>
# The module= option strips the module prefix from go_package to get the relative path
protoc \
    --go_out=. \
    --go_opt=module="$PROJECT" \
    --go-grpc_out=. \
    --go-grpc_opt=module="$PROJECT" \
    --proto_path="$PROTO_DIR" \
    "${PROTO_FILES[@]}"

# Count generated files
GEN_FILES=$(find "$OUT_DIR" -name "*.pb.go" 2>/dev/null | wc -l)
echo ""
echo "Generated $GEN_FILES .pb.go files in $OUT_DIR/"
echo "Proto regeneration complete. Run: go build ./..."
