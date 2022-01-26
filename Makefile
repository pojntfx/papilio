# Variables
DESTDIR ?=

WWWROOT ?= /var/www/html
WWWPREFIX ?= /felix.pojtinger

PREFIX ?= /usr/local
OUTPUT_DIR ?= out
DST ?=

# Private variables
obj = usb-hub-configurator-cli
sts = usb-hub-configurator-pwa
all: $(addprefix build-cli/,$(obj)) $(addprefix build-pwa/,$(sts))

# Build
build: $(addprefix build-cli/,$(obj)) $(addprefix build-pwa/,$(sts))

# Build binary
$(addprefix build-cli/,$(obj)):
ifdef DST
	go build -o $(DST) ./cmd/$(subst build-cli/,,$@)
else
	go build -o $(OUTPUT_DIR)/$(subst build-cli/,,$@) ./cmd/$(subst build-cli/,,$@)
endif

# Build frontend
$(addprefix build-pwa/,$(sts)):
	mkdir -p $(OUTPUT_DIR) public/web
	GOARCH=wasm GOOS=js go build -o web/app.wasm ./cmd/$(subst build-pwa/,,$@)
	go run ./cmd/$(subst build-pwa/,,$@) -prefix $(WWWPREFIX)
	cp -rf web/* public/web
	tar -cvzf $(OUTPUT_DIR)/$(subst build-pwa/,,$@).tar.gz -C public .

# Install
install: $(addprefix install-cli/,$(obj)) $(addprefix install-pwa/,$(sts))

# Install binary
$(addprefix install-cli/,$(obj)):
	install -D -m 0755 $(OUTPUT_DIR)/$(subst install-cli/,,$@) $(DESTDIR)$(PREFIX)/bin/$(subst install-cli/,,$@)

# Install frontend
$(addprefix install-pwa/,$(sts)):
	mkdir -p $(DESTDIR)$(WWWROOT)$(WWWPREFIX)
	cp -rf public/* $(DESTDIR)$(WWWROOT)$(WWWPREFIX)

# Uninstall
uninstall: $(addprefix uninstall-cli/,$(obj)) $(addprefix uninstall-pwa/,$(sts))

# Uninstall binary
$(addprefix uninstall-cli/,$(obj)):
	rm -f $(DESTDIR)$(PREFIX)/bin/$(subst uninstall-cli/,,$@)

# Uninstall frontend
$(addprefix uninstall-pwa/,$(sts)):
	rm -rf $(DESTDIR)$(WWWROOT)$(WWWPREFIX)

# Run
run:
	GOARCH=wasm GOOS=js go build -o web/app.wasm ./cmd/$(subst build-pwa/,,$@)
	go run ./cmd/$(subst build-pwa/,,$@) -serve

# Develop
dev:
	while inotifywait -r -e close_write --exclude 'out' .; do $(MAKE) run; done

# Clean
clean:
	rm -rf out public web/app.wasm

# Dependencies
depend:
	echo 0