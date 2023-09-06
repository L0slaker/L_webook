.PHONY: mock
mock:
# mockgen 要指定三个参数：
# 1.source：接口所在的文件
# 2.destination：生成代码的目标路径
# 3.package：生成代码的文件的package

# 在指定路径下执行该Makefile：D:\go\GO-LEARNING\src\signup_issue
	@mockgen -source=webook\internal\service\user.go -destination=webook\internal\service\mocks\user.mock.go -package=svcmocks
	@mockgen -source=webook\internal\repository\user.go -destination=webook\internal\repository\mocks\user.mock.go -package=repomocks
	@mockgen -source=webook\internal\repository\dao\user.go -destination=webook\internal\repository\dao\mocks\user.mock.go -package=daomocks
	@go mod tidy

