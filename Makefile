DIST_DIR := _output
LOCAL_DIST_DIR := $(DIST_DIR)/local
PLATFORMS := windows linux darwin
PACKAGE_ROOT := gitlab.com/sentinovo/agentkit
OS = $(word 1, $@)
EXECUTABLE := agent


$(DIST_DIR):
	mkdir -p $(DIST_DIR)

$(LOCAL_DIST_DIR):
	mkdir -p $(LOCAL_DIST_DIR)

$(LOCAL_DIST_DIR)/$(EXECUTABLE): $(LOCAL_DIST_DIR)
	go build -o $(LOCAL_DIST_DIR)/$(EXECUTABLE) ./cmd/agent


.PHONY: build
build: $(LOCAL_DIST_DIR)/$(EXECUTABLE)

.PHONY: run
run: $(LOCAL_DIST_DIR)/$(EXECUTABLE)
	$(LOCAL_DIST_DIR)/$(EXECUTABLE)

.PHONY: $(PLATFORMS)
$(PLATFORMS): $(DIST_DIR)
    GOOS=$(OS) GOARCH=amd64 go build -o $(DIST_DIR)/$(EXECUTABLE)-$(OS)-amd64 ./cmd/agent

.PHONY: release
release: windows linux darwin

.PHONY: clean
clean:
	rm -rf $(DIST_DIR)
