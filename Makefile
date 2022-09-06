# Binary name
BINARY=gscript
GOBUILD=go build -ldflags "-s -w" cmd/gscript.go
GOCLEAN=go clean
RMTARGZ=rm -rf *.gz && rm -rf main
VERSION=v0.0.6

# Build
build:
	$(GOCLEAN)
	$(GOBUILD)


clean:
	$(GOCLEAN)
	$(RMTARGZ)


git-tag:
	git tag $(VERSION); \
	git push origin $(VERSION); \

release:
	# Clean
	$(GOCLEAN)
	$(RMTARGZ)
	# Build for mac
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD)
	tar czvf ${BINARY}-mac64-${VERSION}.tar.gz ./${BINARY}
	# Build for arm
	$(GOCLEAN)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GOBUILD)
	tar czvf ${BINARY}-arm64-${VERSION}.tar.gz ./${BINARY}
	# Build for linux
	$(GOCLEAN)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD)
	tar czvf ${BINARY}-linux64-${VERSION}.tar.gz ./${BINARY}
	# Build for win
	$(GOCLEAN)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD).exe
	tar czvf ${BINARY}-win64-${VERSION}.tar.gz ./${BINARY}.exe
	$(GOCLEAN)

# ANTLR Parser Generator  Version 4.9.2
antlr:
	antlr -Dlanguage=Go -o parser -visitor -listener  GScript.g4

bindata:
	go-bindata -pkg internal -o internal/internal.go internal/