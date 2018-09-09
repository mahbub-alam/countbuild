# Written by: MAB

UNAME=$(shell uname)
NAME=countbuild
#PREFIX=github.com/mahbub-alam/$(NAME)
#GOPATH=$(shell go env GOPATH)
BUILD_DIR=./build

ifeq "$(UNAME)" "Linux"
	EXE_EXT=""
	COPY_FILES=./build.txt ./build-conf.txt \
	./countbuild-cmd ./countbuild-new
	CREATE_BUILD_TXT=sh ./countbuild-new	
else ifeq "$(UNAME)" "Darwin"
	EXE_EXT=
	COPY_FILES=./build.txt ./build-conf.txt \
	./countbuild-cmd ./countbuild-new	
	CREATE_BUILD_TXT=sh ./countbuild-new
else
	EXE_EXT=".exe"
	COPY_FILES=./build.txt ./build-conf.txt \
	./countbuild-cmd ./countbuild-new \
	./countbuild-cmd.bat ./countbuild-new.bat \
	./CheckBuildConf.exe
	CREATE_BUILD_TXT=cmd /C countbuild-new.bat
endif

.PHONY: clean

clean:
	rm -rf $(BUILD_DIR)
	rm -f $(NAME)$(EXE_EXT) ./build.txt ./build-conf.txt

_compile: main.go
	go build

build: _compile
	mkdir -p $(BUILD_DIR)
	$(CREATE_BUILD_TXT)
	mv ./$(NAME)$(EXE_EXT) $(BUILD_DIR)/
	cp $(COPY_FILES) $(BUILD_DIR)/

