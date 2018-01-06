APPNAME=$(shell cat APPNAME)
GITUSER=andygeiss
LDFLAGS="-buildmode=c-shared"
TS=$(shell date -u '+%Y/%m/%d %H:%M:%S') 

all: clean test build

$(APPNAME):
	@echo $(TS) Building binaries ...
	@go build $(LDFLAGS) -o build/$(APPNAME).so platform/$(APPNAME)/main.go
	@echo $(TS) Done.

build: $(APPNAME)

clean:
	@echo $(TS) Cleaning up previous build ...
	@rm -f build/$(APPNAME).so
	@echo $(TS) Done.

init:
	@echo $(TS) Creating initial commit ...
	@rm -rf .git
	@git init
	@git add .
	@git commit -m "Initial commit"
	@git remote add origin git@github.com:$(GITUSER)/$(APPNAME).git
	@git push -u --force origin master
	@echo $(TS) Done.

test:
	@echo $(TS) Testing ...
	@go test -v github.com/$(GITUSER)/$(APPNAME)/...
	@echo $(TS) Done.
