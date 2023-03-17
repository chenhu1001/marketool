# marketool

A market investment analysis tool

## 使用 GoReleaser 发布你的应用

初始化

```
goreleaser init
```

添加tag
```
git tag v1.0.0
```

执行自动发布流程 Mac

```
brew install goreleaser
goreleaser --snapshot --skip-publish --rm-dist
```

Linux

```
./goreleaser --snapshot --skip-publish --rm-dist
```

Windows

```
goreleaser.exe --snapshot --skip-publish --rm-dist
或
.\goreleaser.exe --snapshot --skip-publish --rm-dist
```

# 使用Go.mod官方包管理工具

```
go mod init marketool
```

作用：

- 运行完之后，会在当前目录下生成一个go.mod文件，这是一个关键文件，之后的包的管理都是通过这个文件管理。
-
官方说明：除了go.mod之外，go命令还维护一个名为go.sum的文件，其中包含特定模块版本内容的预期加密哈希，go命令使用go.sum文件确保这些模块的未来下载检索与第一次下载相同，以确保项目所依赖的模块不会出现意外更改，无论是出于恶意、意外还是其他原因。go.mod和go.sum都应检入版本控制。
go.sum 不需要手工维护，所以可以不用太关注。

```bigquery
go mod tidy
```

作用：

- 引用项目需要的依赖增加到go.mod文件。
- 去掉go.mod文件中项目不需要的依赖。