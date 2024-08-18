CLANG ?= clang-14
STRIP ?= llvm-strip-14
OBJCOPY ?= llvm-objcopy-14
TARGET_ARCH ?= arm64
CFLAGS := -O2 -g -Wall -Werror -D__TARGET_ARCH_$(TARGET_ARCH) $(CFLAGS)
REPODIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
UIDGID := $(shell stat -c '%u:%g' ${REPODIR})
CONTAINER_ENGINE ?= docker
CONTAINER_RUN_ARGS ?= $(--user "${UIDGID}")
IMAGE_GENERATE := ebpf-builder
VERSION_GENERATE := v1-1.22.1
GENERATE_DOCKERFILE := ebpf-builder/Dockerfile
TARGETS := \

.PHONY: go_builder_image_build
go_builder_image_build:
	docker build -t ${IMAGE_GENERATE}:${VERSION_GENERATE} -f ${GENERATE_DOCKERFILE} .


.PHONY: all clean go_generate container-shell generate

.DEFAULT_TARGET = go_generate

go_generate:
	+${CONTAINER_ENGINE} run --rm ${CONTAINER_RUN_ARGS} \
		-v "${REPODIR}":/ebpf -w /ebpf --env MAKEFLAGS \
		--env CFLAGS="-fdebug-prefix-map=/ebpf=." \
		--env HOME="/tmp" \
		"${IMAGE_GENERATE}:${VERSION_GENERATE}" \
		make all
container-shell:
	${CONTAINER_ENGINE} run --rm -ti \
		-v "${REPODIR}":/ebpf -w /ebpf \
		"${IMAGE_GENERATE}:${VERSION_GENERATE}"


all: generate
generate: export BPF_CLANG := $(CLANG)
generate: export BPF_CFLAGS := $(CFLAGS)
generate:
	go generate ./...

%-el.elf: %.c
	$(CLANG) $(CFLAGS) -target bpfel -g -c $< -o $@
	$(STRIP) -g $@

%-eb.elf : %.c
	$(CLANG) $(CFLAGS) -target bpfeb -c $< -o $@
	$(STRIP) -g $@

MARZYIFY_IMAGE_NAME := MarzyIfy
MARZYIFY_TAG ?= latest
REGISTRY ?= Taxonim
MARZYIFY_DOCKERFILE ?= Dockerfile.default
BUILDX_BUILDER := buildx-multi-arch

ifeq ($(TARGET_ARCH), arm64)
	DOCKER_PLATFORM := linux/arm64
else
	DOCKER_PLATFORM := linux/amd64
endif

.PHONY: build_push_buildx
build_push_buildx:
	docker buildx inspect $(BUILDX_BUILDER) || \
	docker buildx create --name=$(BUILDX_BUILDER) && \
	docker buildx build --push --platform=$(DOCKER_PLATFORM) --builder=$(BUILDX_BUILDER) --build-arg MARZYIFY_TAG=$(MARZYIFY_TAG) --build-arg VERSION=$(MARZYIFY_TAG) --tag=$(REGISTRY)/$(MARZYIFY_IMAGE_NAME):$(MARZYIFY_TAG)-$(TARGET_ARCH) -f $(MARZYIFY_DOCKERFILE) .
.PHONY: docker_merge_platforms
docker_merge_platforms:
	docker buildx imagetools create --tag $(REGISTRY)/$(MARZYIFY_IMAGE_NAME):$(MARZYIFY_TAG) $(REGISTRY)/$(MARZYIFY_IMAGE_NAME):$(MARZYIFY_TAG)-arm64 $(REGISTRY)/$(MARZYIFY_IMAGE_NAME):$(MARZYIFY_TAG)-x86
.PHONY: build_push
build_push:
	docker build --build-arg VERSION=$(MARZYIFY_TAG) -t $(REGISTRY)/$(MARZYIFY_IMAGE_NAME):$(MARZYIFY_TAG)  -f $(MARZYIFY_DOCKERFILE) .
	docker push $(REGISTRY)/$(MARZYIFY_IMAGE_NAME):$(MARZYIFYTAG)
