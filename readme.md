# Coin Trader

# Setup

**Install Go**

```bash
$ brew update
$ brew install golang
```

**Update your .bashrc file**

```bash
export GOPATH=$HOME/go-workspace # Change this to your go workspace.
export GOROOT=/usr/local/bin/go # Location of your go installation
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
```

## Bittrex

*Create an API Key on Bittrex*

Navigate to [Bittrex Settings](https://bittrex.com/Manage#sectionApi)


**Run**

`BITTREX_API_KEY={{ API_KEY }} BITTREX_API_SECRET={{ API_SECRET }} go run main.go`

