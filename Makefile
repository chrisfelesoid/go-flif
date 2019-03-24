
.PHONY: all
all:
	$(MAKE) prepare
	$(MAKE) build

.PHONY: prepare
prepare:
	cd build && cmake ..

.PHONY: build
build:
	cd build && $(MAKE) all

.PHONY: install
install:
	cd build && $(MAKE) install

.PHONY: uninstall
uninstall:
	xargs rm < build/install_manifest.txt
