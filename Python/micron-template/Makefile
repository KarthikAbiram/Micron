# -*- coding: utf-8 -*-

# Automatically get the name of the top-level project folder
APP_NAME := $(notdir $(CURDIR))

PYTHON = uv run python
BUILD_DIR = build

SERVER_APP = $(APP_NAME)
CLIENT_APP = $(APP_NAME)_client

# Detect OS for cross-platform clean
ifeq ($(OS),Windows_NT)
    RM = powershell Remove-Item -Recurse -Force
else
    RM = rm -rf
endif

# ===== Targets =====

.PHONY: all
all: help

.PHONY: install
install:
	uv sync

# ===== Run Targets =====
.PHONY: run run-server
run-server:
	@echo "🚀 Starting server..."
	$(PYTHON) server.py

.PHONY: run-client
run-client:
	@echo "🚀 Starting client..."
	$(PYTHON) client.py

# ===== Build Targets =====
.PHONY: build
build: build-server build-client
	@echo "✅ Built $(SERVER_APP) and $(CLIENT_APP)"

.PHONY: build-server
build-server:
	$(PYTHON) -m cx_Freeze server.py --target-name $(SERVER_APP) --build-exe $(BUILD_DIR)/$(SERVER_APP)
	@echo "🧊 Server build complete -> $(BUILD_DIR)/$(SERVER_APP)"

.PHONY: build-client
build-client:
	$(PYTHON) -m cx_Freeze client.py --target-name $(CLIENT_APP) --build-exe $(BUILD_DIR)/$(CLIENT_APP)
	@echo "🧊 Client build complete -> $(BUILD_DIR)/$(CLIENT_APP)"

# ===== Clean Targets =====
.PHONY: clean
clean:
	$(RM) $(BUILD_DIR)
	@echo "🧹 Cleaned build directories."

# ===== Lock Dependencies =====
.PHONY: lock
lock:
	uv lock
	@echo "🔒 Dependencies locked."

# ===== Help =====
.PHONY: help
help:
	@echo "Makefile commands for project: $(APP_NAME)"
	@echo "  make install       - Install dependencies"
	@echo "  make run-server    - Run server.py"
	@echo "  make run-client    - Run client.py"
	@echo "  make build         - Build both executables"
	@echo "  make clean         - Remove build artifacts"
	@echo "  make lock          - Lock dependencies"
