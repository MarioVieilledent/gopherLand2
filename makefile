# Define variables
APP_NAME=GopherLand2.exe
DIST_DIR=dist

# Define targets and dependencies
all: build

build:
	go build -o $(DIST_DIR)/$(APP_NAME) -ldflags "-w -s"
	xcopy /E /I /Y data $(DIST_DIR)\data
	copy README.md $(DIST_DIR)