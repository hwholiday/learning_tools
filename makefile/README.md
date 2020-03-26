## 什么是MakeFile
Makefile文件的作用是告诉make工具需要如何去编译和链接程序，在需要编译工程时只需要一个make命令即可，避免了每次编译都要重新输入完整命令的麻烦，大大提高了效率，也减少了出错率。
## 目录
```base
├── main.go
├── Makefile
└── README.md
```
## Makefile 例子
```base
#环境变量
export VERSION=1.0.0
export ENV=prod
export PROJECT=test
#变量
PWD=$(shell pwd)
OUT_DIR=$(PWD)/cmd
OUT_NAME=$(PROJECT)_$(VERSION)
OUT_FILE=$(OUT_DIR)/$(OUT_NAME)
MAIN_FILE=main.go
PACK_FILE=$(OUT_NAME).tar.gz

build:
	@echo "移除老项目"
	@ rm -rf $(OUT_DIR)/*
	@go build -o $(OUT_FILE) $(MAIN_FILE)

	@echo "构建项目成功 : " $(OUT_NAME)

run:build
	@# 执行pack前先执行一次build
	@echo "运行项目"
	@$(OUT_FILE)

pack:build
	@# 执行pack前先执行一次build
	tar -czvf $(PACK_FILE) -C $(OUT_DIR) .
	@echo "打包结束 :"
```
## 构建命令
```base
# 带上环境变量
make build ENV="prod" VERSION="2.0.0" PROJECT="test_makefile"
```

## 打包命令
```base
make pack
```
## 运行命令
```base
make run
```