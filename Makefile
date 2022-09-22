PROJECT_NAME := Pulumi Native Google Cloud Resource Provider

PACK            := google-native
PACKDIR         := sdk
PROJECT         := github.com/pulumi/pulumi-google-native
PROVIDER        := pulumi-resource-${PACK}
CODEGEN         := pulumi-gen-${PACK}
VERSION         := $(shell pulumictl get version)

PROVIDER_PKGS   := $(shell cd ./provider && go list ./...)
WORKING_DIR     := $(shell pwd)

JAVA_GEN 		 := pulumi-java-gen
JAVA_GEN_VERSION := v0.5.4

VERSION_FLAGS   := -ldflags "-X github.com/pulumi/pulumi-${PACK}/provider/pkg/version.Version=${VERSION}"

ensure::
	@echo "go mod download"; cd provider; go mod download

local_generate:: bin/pulumi-java-gen
	$(WORKING_DIR)/bin/$(CODEGEN) schema,nodejs,dotnet,python,go ${VERSION}
	$(WORKING_DIR)/bin/$(JAVA_GEN) generate --schema $(WORKING_DIR)/provider/cmd/$(PROVIDER)/schema.json --out sdk/java --build gradle-nexus
	echo "Finished generating schema."

generate_schema:: bin/pulumi-java-gen
	echo "Generating Pulumi schema..."
	$(WORKING_DIR)/bin/$(CODEGEN) schema ${VERSION}
	#$(WORKING_DIR)/bin/$(JAVA_GEN) generate --schema $(WORKING_DIR)/provider/cmd/$(PROVIDER)/schema.json --out sdk/java --build gradle-nexus
	echo "Finished generating schema."

codegen::
	(cd provider && go build -o $(WORKING_DIR)/bin/$(CODEGEN) $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(CODEGEN))

provider::
	(cd provider && go build -o $(WORKING_DIR)/bin/$(PROVIDER) $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(PROVIDER))

debug_provider::
	(cd provider && go install -gcflags="all=-N -l" $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(PROVIDER))

test_provider::
	(cd provider && go test -v $(PROVIDER_PKGS))

lint_provider:: provider # lint the provider code
	cd provider && GOGC=20 golangci-lint run -c ../.golangci.yml

discovery::codegen
	$(WORKING_DIR)/bin/$(CODEGEN) discovery ${VERSION}

generate_nodejs::
	$(WORKING_DIR)/bin/$(CODEGEN) nodejs ${VERSION}

build_nodejs:: VERSION := $(shell pulumictl get version --language javascript)
build_nodejs::
	cd ${PACKDIR}/nodejs/ && \
		yarn install && \
		node --max-old-space-size=4096 ./node_modules/.bin/tsc --diagnostics && \
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

generate_java:: bin/pulumi-java-gen
	$(WORKING_DIR)/bin/$(JAVA_GEN) generate --schema $(WORKING_DIR)/provider/cmd/$(PROVIDER)/schema.json --out sdk/java --build gradle-nexus

build_java:: PACKAGE_VERSION := $(shell pulumictl get version --language generic)
build_java::
	cd ${PACKDIR}/java/ && \
		gradle --console=plain build

bin/pulumi-java-gen::
	$(shell pulumictl download-binary -n pulumi-language-java -v $(JAVA_GEN_VERSION) -r pulumi/pulumi-java)

generate_go::
	$(WORKING_DIR)/bin/$(CODEGEN) go ${VERSION}

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
build_sdks: build_nodejs build_dotnet build_python build_go #build_java
install_sdks:: install_dotnet_sdk install_python_sdk install_nodejs_sdk #install_java_sdk

# Required for the codegen action that runs in pulumi/pulumi
only_build:: build

.PHONY: init_submodules update_submodules ensure generate_schema generate build_provider build
