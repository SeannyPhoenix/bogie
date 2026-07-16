.PHONY: all dist serve

all: dist

ESBUILD := npx esbuild
ENTRYPOINTS := src/main.ts static/index.html
SRC_FILES := $(shell find src -type f)
OUT_DIR := local/www
ESBUILD_FLAGS := --bundle --format=esm --minify --sourcemap --outdir=$(OUT_DIR) --platform=browser --target=esnext --loader:.html=copy --entry-names=[name]

dist: $(SRC_FILES)
	rm -rf $(OUT_DIR)
	$(ESBUILD) $(ENTRYPOINTS) $(ESBUILD_FLAGS)

serve:
	$(ESBUILD) $(ENTRYPOINTS) $(ESBUILD_FLAGS) --servedir=$(OUT_DIR) --watch --serve-fallback=static/index.html
