# go-test
チームで一時期行なっていたgoでテスト自動化を弄くり回す

# 使い方（MacOS前提）

## ターミナルでChrome用のWebDriverをインストール
```
brew instsall chromedriver
```
または
```
brew tap homebrew/cask
brew cask install chromedriver
```

## agoutiのインストール
$GOPATHで
```
go get go get github.com/sclevine/agouti
```

## プログラムの実行
```
go run main.go
```

### 参考文献

[Go + WebDriver でブラウザ操作を自動化する](https://qiita.com/utahta/items/17afea7933624371504c)

[Goではじめてみたブラウザの自動操作](https://qiita.com/0829/items/c1e494bb128ade5f0872)

[agoutiというWebDriverクライアントを使って面倒な作業を自動化する](https://qiita.com/tenten0213/items/1f897ff8a64bd8b5270c)