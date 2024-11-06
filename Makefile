TAILPIPE_INSTALL_DIR ?= ~/.tailpipe
BUILD_TAGS = netgo
install:
	go build -o $(TAILPIPE_INSTALL_DIR)/plugins/hub.tailpipe.io/plugins/turbot/chaos@latest/tailpipe-plugin-chaos.plugin -tags "${BUILD_TAGS}" *.go
