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

$(LOCAL_DIST_DIR)/agent: $(LOCAL_DIST_DIR)
	go build -o $(LOCAL_DIST_DIR)/agent ./cmd/agent

$(LOCAL_DIST_DIR)/agentctl: $(LOCAL_DIST_DIR)
	go build -o $(LOCAL_DIST_DIR)/agentctl ./cmd/agentctl

$(LOCAL_DIST_DIR)/agentcoordinator: $(LOCAL_DIST_DIR)
	go build -o $(LOCAL_DIST_DIR)/agentcoordinator ./cmd/agentcoordinator

.PHONY: build
build: $(LOCAL_DIST_DIR)/agent $(LOCAL_DIST_DIR)/agentctl $(LOCAL_DIST_DIR)/agentcoordinator

.PHONY: run
run: $(LOCAL_DIST_DIR)/$(EXECUTABLE)
	$(LOCAL_DIST_DIR)/$(EXECUTABLE)

.PHONY: $(PLATFORMS)
$(PLATFORMS): $(DIST_DIR)
    GOOS=$(OS) GOARCH=amd64 go build -o $(DIST_DIR)/agent-$(OS)-amd64 ./cmd/agent
	GOOS=$(OS) GOARCH=amd64 go build -o $(DIST_DIR)/agentctl-$(OS)-amd64 ./cmd/agentctl
	GOOS=$(OS) GOARCH=amd64 go build -o $(DIST_DIR)/agentcoordinator-$(OS)-amd64 ./cmd/agentctl

.PHONY: release
release: windows linux darwin

.PHONY: clean
clean:
	rm -rf $(DIST_DIR)
