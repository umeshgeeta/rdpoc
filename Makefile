# Define the target executable file name
APP_NAME=my-app

# Define the Go files to build and compile
GO_FILES=$(wildcard src/server/*.go)

# Define the target build directory
BUILD_DIR=build

# Define the target build file
BUILD_FILE=$(BUILD_DIR)/$(APP_NAME)

# Define the target run command
RUN_CMD=./$(BUILD_FILE)

# Define the Go compiler and linker flags
GO_COMPILER=go
GO_COMPILER_FLAGS=-ldflags="-s -w"

# Define the Node.js and webpack commands
NODE=node
WEBPACK=./node_modules/.bin/webpack

# Define the public directory
PUBLIC_DIR=public

# Define the clean target
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR) $(PUBLIC_DIR)/bundle.js

# Define the build target
.PHONY: build
build: clean
	mkdir -p $(BUILD_DIR) $(PUBLIC_DIR)
	$(NODE) $(WEBPACK) --mode=production
	$(GO_COMPILER) build $(GO_COMPILER_FLAGS) -o $(BUILD_FILE) $(GO_FILES)

# Define the run target
.PHONY: run
run: build
	$(RUN_CMD)
