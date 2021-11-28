# GoBot

GoBot 是一个基于 Golang 与 [MiraiGo](https://github.com/Mrs4s/MiraiGo) 开发的插件化 QQ 机器人（框架）

## 注意

- 由于本项目的特殊性，本项目随时可能终止开发，请您务必注意这个问题。
- 以及由于本项目的开源协议，您**必须开源**您所编写的插件。
- 您不可以使用本项目盈利
- 本项目使用了 Golang 不能跨平台的特性`plugin`，因此并不推荐在 Windows 平台使用，实际上他在 Windows 平台上并不会有任何能力。
- 由于 Golang `plugin` 的特殊性，您必须自己保证您加载的插件与构建版本使用的 Golang 版本一致。因此我们不会提供任何预构建文件。
