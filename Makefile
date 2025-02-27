PROJECT_NAME := Pulumi Native Google Cloud Resource Provider

PACK            := google-native
PACKDIR         := sdk
PROJECT         := github.com/pulumi/pulumi-google-native
PROVIDER        := pulumi-resource-${PACK}
CODEGEN         := pulumi-gen-${PACK}

PROVIDER_PKGS   := $(shell cd ./provider && go list ./...)
WORKING_DIR     := $(shell pwd)

JAVA_GEN		 := pulumi-java-gen
JAVA_GEN_VERSION := v0.5.4

# Override during CI using `make [TARGET] PROVIDER_VERSION=""` or by setting a PROVIDER_VERSION environment variable
# Local & branch builds will just used this fixed default version unless specified
PROVIDER_VERSION ?= 0.0.1-alpha.0+dev
# Use this normalised version everywhere rather than the raw input to ensure consistency.
VERSION_GENERIC := $(shell pulumictl convert-version --language generic --version "$(PROVIDER_VERSION)")


VERSION_FLAGS   := -ldflags "-X github.com/pulumi/pulumi-${PACK}/provider/pkg/version.Version=$(VERSION_GENERIC)"

ensure::
	@echo "go mod download"; cd provider; go mod download

local_generate:: bin/pulumi-java-gen
	$(WORKING_DIR)/bin/$(CODEGEN) schema,nodejs,dotnet,python,go $(VERSION_GENERIC)
	$(WORKING_DIR)/bin/$(JAVA_GEN) generate --schema $(WORKING_DIR)/provider/cmd/$(PROVIDER)/schema.json --out sdk/java --build gradle-nexus
	echo "Finished generating schema."

generate_schema:: bin/pulumi-java-gen
	echo "Generating Pulumi schema..."
	$(WORKING_DIR)/bin/$(CODEGEN) schema $(VERSION_GENERIC)
	$(WORKING_DIR)/bin/$(JAVA_GEN) generate --schema $(WORKING_DIR)/provider/cmd/$(PROVIDER)/schema.json --out sdk/java --build gradle-nexus
	echo "Finished generating schema."

codegen::
	(cd provider && go build -o $(WORKING_DIR)/bin/$(CODEGEN) $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(CODEGEN))

provider::
	(cd provider && go build -o $(WORKING_DIR)/bin/$(PROVIDER) $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(PROVIDER))

debug_provider::
	(cd provider && go install -gcflags="all=-N -l" $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(PROVIDER))

test_provider::
	(cd provider && go test -v -coverprofile="coverage.txt" -coverpkg=./... $(PROVIDER_PKGS))

lint_provider:: provider # lint the provider code
	cd provider && GOGC=20 golangci-lint run -c ../.golangci.yml

discovery::codegen
	$(WORKING_DIR)/bin/$(CODEGEN) discovery $(VERSION_GENERIC)

generate_nodejs::
	$(WORKING_DIR)/bin/$(CODEGEN) nodejs $(VERSION_GENERIC)

build_nodejs::
	cd ${PACKDIR}/nodejs/ && \
		yarn install && \
		node --max-old-space-size=4096 ./node_modules/.bin/tsc --diagnostics && \
		cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/

generate_python::
	# Delete files not tracked in Git
	cd sdk/python/ && git clean -fxd

	$(WORKING_DIR)/bin/$(CODEGEN) python $(VERSION_GENERIC)

build_python::
	# Delete files not tracked in Git
	cd sdk/python/ && git clean -fxd
	cd sdk/python/ && \
        cp ../../README.md . && \
        rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
	python3 -m venv venv && \
	./venv/bin/python -m pip install build && \
        cd ./bin && \
        ../venv/bin/python -m build .

generate_dotnet::
	$(WORKING_DIR)/bin/$(CODEGEN) dotnet $(VERSION_GENERIC)

build_dotnet::
	cd ${PACKDIR}/dotnet/ && \
		dotnet build

generate_java:: bin/pulumi-java-gen
	$(WORKING_DIR)/bin/$(JAVA_GEN) generate --schema $(WORKING_DIR)/provider/cmd/$(PROVIDER)/schema.json --out sdk/java --build gradle-nexus

java_sdk:: PACKAGE_VERSION := $(shell pulumictl convert-version --language generic -v "$(VERSION_GENERIC)")
build_java::
	cd ${PACKDIR}/java/ && \
		gradle --console=plain build

bin/pulumi-java-gen::
	$(shell pulumictl download-binary -n pulumi-language-java -v $(JAVA_GEN_VERSION) -r pulumi/pulumi-java)

generate_go::
	$(WORKING_DIR)/bin/$(CODEGEN) go $(VERSION_GENERIC)

build_go::
	cd sdk/ && go build github.com/pulumi/pulumi-google-native/sdk/go/google/...

clean::
	rm -rf sdk/nodejs && mkdir sdk/nodejs && echo "module fake_nodejs_module // Exclude this directory from Go tools\n\ngo 1.17" > 'sdk/nodejs/go.mod'
	rm -rf sdk/python && mkdir sdk/python && echo "module fake_python_module // Exclude this directory from Go tools\n\ngo 1.17" > 'sdk/python/go.mod' && cp README.md sdk/python
	rm -rf sdk/dotnet && mkdir sdk/dotnet && echo "module fake_dotnet_module // Exclude this directory from Go tools\n\ngo 1.17" > 'sdk/dotnet/go.mod'
	rm -rf sdk/java && mkdir sdk/java && echo "module fake_java_module // Exclude this directory from Go tools\n\ngo 1.17" > 'sdk/java/go.mod'
	rm -rf sdk/go/google

install:: install_nodejs_sdk install_dotnet_sdk
	cp $(WORKING_DIR)/bin/${PROVIDER} ${GOPATH}/bin

install_dotnet_sdk::
	mkdir -p $(WORKING_DIR)/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

install_python_sdk::

install_go_sdk::

install_java_sdk::

install_nodejs_sdk::
	-yarn unlink --cwd $(WORKING_DIR)/sdk/nodejs/bin
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin

test::
	cd examples && go test -v -tags=all -timeout 2h

build:: init_submodules clean codegen local_generate provider build_sdks install_sdks
build_sdks: build_nodejs build_dotnet build_python build_go build_java
install_sdks:: install_dotnet_sdk install_python_sdk install_nodejs_sdk install_java_sdk

# Required for the codegen action that runs in pulumi/pulumi
only_build:: build

.PHONY: init_submodules update_submodules ensure generate_schema generate build_provider build

# Set these variables to enable signing of the windows binary
AZURE_SIGNING_CLIENT_ID ?=
AZURE_SIGNING_CLIENT_SECRET ?=
AZURE_SIGNING_TENANT_ID ?=
AZURE_SIGNING_KEY_VAULT_URI ?=
SKIP_SIGNING ?=

bin/jsign-6.0.jar:
	wget https://github.com/ebourg/jsign/releases/download/6.0/jsign-6.0.jar --output-document=bin/jsign-6.0.jar

sign-goreleaser-exe-amd64: GORELEASER_ARCH := amd64_v1
sign-goreleaser-exe-arm64: GORELEASER_ARCH := arm64

# Set the shell to bash to allow for the use of bash syntax.
sign-goreleaser-exe-%: SHELL:=/bin/bash
sign-goreleaser-exe-%: bin/jsign-6.0.jar
	@# Only sign windows binary if fully configured.
	@# Test variables set by joining with | between and looking for || showing at least one variable is empty.
	@# Move the binary to a temporary location and sign it there to avoid the target being up-to-date if signing fails.
	@set -e; \
	if [[ "${SKIP_SIGNING}" != "true" ]]; then \
		if [[ "|${AZURE_SIGNING_CLIENT_ID}|${AZURE_SIGNING_CLIENT_SECRET}|${AZURE_SIGNING_TENANT_ID}|${AZURE_SIGNING_KEY_VAULT_URI}|" == *"||"* ]]; then \
			echo "Can't sign windows binaries as required configuration not set: AZURE_SIGNING_CLIENT_ID, AZURE_SIGNING_CLIENT_SECRET, AZURE_SIGNING_TENANT_ID, AZURE_SIGNING_KEY_VAULT_URI"; \
			echo "To rebuild with signing delete the unsigned windows exe file and rebuild with the fixed configuration"; \
			if [[ "${CI}" == "true" ]]; then exit 1; fi; \
		else \
			file=dist/build-provider-sign-windows_windows_${GORELEASER_ARCH}/pulumi-resource-google-native.exe; \
			mv $${file} $${file}.unsigned; \
			az login --service-principal \
				--username "${AZURE_SIGNING_CLIENT_ID}" \
				--password "${AZURE_SIGNING_CLIENT_SECRET}" \
				--tenant "${AZURE_SIGNING_TENANT_ID}" \
				--output none; \
			ACCESS_TOKEN=$$(az account get-access-token --resource "https://vault.azure.net" | jq -r .accessToken); \
			java -jar bin/jsign-6.0.jar \
				--storetype AZUREKEYVAULT \
				--keystore "PulumiCodeSigning" \
				--url "${AZURE_SIGNING_KEY_VAULT_URI}" \
				--storepass "$${ACCESS_TOKEN}" \
				$${file}.unsigned; \
			mv $${file}.unsigned $${file}; \
			az logout; \
		fi; \
	fi
