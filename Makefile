VERSION=v0.0.2
TARGET_DIR=pkg/dist/${VERSION}

build:
	gox -output "${TARGET_DIR}/{{.OS}}_{{.Arch}}/{{.Dir}}"

zip: build
	bash scripts/zip.sh ${TARGET_DIR}

release: zip
	ghr ${VERSION} ${TARGET_DIR}

clean:
	rm -rf ${TARGET_DIR}
