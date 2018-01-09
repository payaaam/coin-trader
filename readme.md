# Coin Trader

# Setup

**Install Go**

```bash
$ brew update
$ brew install golang
```

**Update your .bashrc / .bash_profile file**

```bash
export GOPATH=$HOME/go-workspace # Change this to your go workspace.
export GOROOT=/usr/local/opt/go/libexec # Location of your go installation
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
```

**Setup your environment**
```bash
$ mkdir -p $GOPATH $GOPATH/src $GOPATH/pkg $GOPATH/bin
$ cd $GOPATH/src
$ git clone https://github.com/payaaam/coin-trader.git
$ go get
```

## Bittrex

**Create an API Key**

Navigate to [Bittrex Settings](https://bittrex.com/Manage#sectionApi)


**Run**

`$ BITTREX_API_KEY={{ API_KEY }} BITTREX_API_SECRET={{ API_SECRET }} go run main.go`

