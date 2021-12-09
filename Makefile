PROJECT_NAME := Pulumi Native Google Cloud Resource Provider

PACK            := google-native
PACKDIR         := sdk
PROJECT         := github.com/pulumi/pulumi-google-native
PROVIDER        := pulumi-resource-${PACK}
CODEGEN         := pulumi-gen-${PACK}
VERSION         := $(shell pulumictl get version)

PROVIDER_PKGS   := $(shell cd ./provider && go list ./...)
WORKING_DIR     := $(shell pwd)

VERSION_FLAGS   := -ldflags "-X github.com/pulumi/pulumi-${PACK}/provider/pkg/version.Version=${VERSION}"

ensure::
	@echo "go mod download"; cd provider; go mod download

local_generate::
	$(WORKING_DIR)/bin/$(CODEGEN) schema,nodejs,dotnet,python,go ${VERSION}
	echo "Finished generating schema."

generate_schema::
	echo "Generating Pulumi schema..."
	$(WORKING_DIR)/bin/$(CODEGEN) schema ${VERSION}
	echo "Finished generating schema."

codegen::
	(cd provider && go build -a -o $(WORKING_DIR)/bin/$(CODEGEN) $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(CODEGEN))

provider::
	(cd provider && go build -a -o $(WORKING_DIR)/bin/$(PROVIDER) $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(PROVIDER))

test_provider::
	(cd provider && go test -v $(PROVIDER_PKGS))

lint_provider:: provider # lint the provider code
	cd provider && GOGC=20 golangci-lint run -c ../.golangci.yml

discovery::codegen
	$(WORKING_DIR)/bin/$(CODEGEN) discovery ${VERSION}

generate_nodejs::
	$(WORKING_DIR)/bin/$(CODEGEN) nodejs ${VERSION}

build_nodejs:: VERSION := $(shell pulumictl get version --language javascript)
build_nodejs:: TSC := $(shell which tsc)
build_nodejs::
	cd ${PACKDIR}/nodejs/ && \
		yarn install && \
		node --max-old-space-size=4096 $(TSC) --diagnostics && \
		cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/ && \
		sed -i.bak -e "s/\$${VERSION}/$(VERSION)/g" ./bin/package.json

generate_python::
	$(WORKING_DIR)/bin/$(CODEGEN) python ${VERSION}

build_python:: PYPI_VERSION := $(shell pulumictl get version --language python)
build_python::
	cd sdk/python/ && \
        cp ../../README.md . && \
        python3 setup.py clean --all 2>/dev/null && \
        rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
        sed -i.bak -e 's/^VERSION = .*/VERSION = "$(PYPI_VERSION)"/g' -e 's/^PLUGIN_VERSION = .*/PLUGIN_VERSION = "$(VERSION)"/g' ./bin/setup.py && \
        rm ./bin/setup.py.bak && \
        cd ./bin && python3 setup.py build sdist

generate_dotnet::
	$(WORKING_DIR)/bin/$(CODEGEN) dotnet ${VERSION}

build_dotnet:: DOTNET_VERSION := $(shell pulumictl get version --language dotnet)
build_dotnet::
	cd ${PACKDIR}/dotnet/ && \
		echo "${PACK}\n${DOTNET_VERSION}" >version.txt && \
		dotnet build /p:Version=${DOTNET_VERSION}

generate_go::
	$(WORKING_DIR)/bin/$(CODEGEN) go ${VERSION}

build_go::
	cd sdk/ && go build github.com/pulumi/pulumi-google-native/sdk/go/google/...

clean::
	rm -rf sdk/nodejs && mkdir sdk/nodejs && touch sdk/nodejs/go.mod
	rm -rf sdk/python && mkdir sdk/python && touch sdk/python/go.mod && cp README.md sdk/python
	rm -rf sdk/dotnet && mkdir sdk/dotnet && touch sdk/dotnet/go.mod
	rm -rf sdk/go/google

install_dotnet_sdk::
	mkdir -p $(WORKING_DIR)/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

install_python_sdk::

install_go_sdk::

install_nodejs_sdk::
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin

test::
	cd examples && go test -v -tags=all -timeout 2h

build:: init_submodules clean codegen local_generate provider build_sdks install_sdks
build_sdks: build_nodejs build_dotnet build_python build_go
install_sdks:: install_dotnet_sdk install_python_sdk install_nodejs_sdk

# Required for the codegen action that runs in pulumi/pulumi
only_build:: build

.PHONY: init_submodules update_submodules ensure generate_schema generate build_provider build
