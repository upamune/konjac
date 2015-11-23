# konjac
[![Circle CI](https://circleci.com/gh/upamune/konjac/tree/master.svg?style=svg)](https://circleci.com/gh/upamune/konjac/tree/master)
![](https://i.gyazo.com/bdab6bee047af3065ce5f3e71e3587a3.gif)


## Description
日=>英のみの翻訳ツールです. Microsoft翻訳APIを使用しているので, 登録が必要です.
ここにいって, キーを取得しましょう.

[Microsoft Translator](https://datamarket.azure.com/dataset/bing/microsofttranslator)

登録のくわしい説明です.

[Microsoft Translator APIを使ってみる - Qiita](http://qiita.com/kemayako/items/21fe36005e6e729aff77)

## Usage

```bash
$ konjac -c config.toml 寿司食べたい
```

```toml:config.toml
[client]
id = "myid"
secret = "mysecret"
```

## Install

Windows, MacOSX, Linuxとかに対応しています. みんな大好きSolarisにも対応しています!

[Releases](https://github.com/upamune/konjac/releases)

```bash
$ go get -d github.com/upamune/konjac
```

## Contribution

1. Fork ([https://github.com/upamune/konjac/fork](https://github.com/upamune/konjac/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[upamune](https://github.com/upamune)
