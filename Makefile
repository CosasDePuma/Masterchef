# App
APPNAME		= masterchef
VERSION		= $(shell cat version)
FRONTEND	= frontend
BACKEND		= backend

# Compiler
CC			= go build
OBJS		= version $(BACKEND)/public/app.go $(BACKEND)/public/index.go
LDFLAGS		= -ldflags="-s -w -extldflags \"-static\""

.PHONY: all
all: build

.PHONY: FORCE
FORCE:

.PHONY: build
build: dist/$(APPNAME)-$(VERSION)
	@upx -9 --ultra-brute $<

dist/$(APPNAME)-$(VERSION): fmt $(OBJS)
	$(CC) $(LDFLAGS) -a -o $@ main.go

$(BACKEND)/public/index.go: FORCE $(FRONTEND)/node_modules $(FRONTEND)/package.json
	@yarn --cwd $(FRONTEND) run compile

$(FRONTEND)/node_modules: $(FRONTEND)/package.json
	@yarn --cwd $(FRONTEND) install

.PHONY: fmt
fmt:
	@gofmt -w $(BACKEND)

.PHONY: clean
clean:
	@rm -rf $(FRONTEND)/dist