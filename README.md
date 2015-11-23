# konjac



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

ホーム直下に `.konjac.toml` を配置するとcオプション指定なしでできます.

```bash
$ konjac 寿司食べたい
```


## Install

To install, use `go get`:

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
