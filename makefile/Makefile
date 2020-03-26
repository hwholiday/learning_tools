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