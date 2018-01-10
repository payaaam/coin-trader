# Coin Trader

# Setup

**Install Go**

```bash
$ brew update
$ brew install golang
```

**Update your .bashrc file**

Add these lines to your `.bashrc` file.

```bash
export GOPATH=$HOME/{{ GO_WORKSPACE_DIR }} # Change this to your go workspace.
export GOROOT=/usr/local/bin/go # Location of your go installation
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
```

`$ source ~/.bash_profile`

**Install Postgres**

1. Download [Postgres](https://postgresapp.com/)
2. Download [Postico](https://eggerapps.at/postico/)
3. Open Postgres App and `Initialize`
4. Open Postico and connect with these settings...

```
Host: localhost
Password: _empty_
Port: 5432
Database: postgres
```

**Setup DB**

`$ make setup-db`

## Bittrex

**Create an API Key**

Navigate to [Bittrex Settings](https://bittrex.com/Manage#sectionApi)

## Build

`$ make cli`


**Run**

`$ BITTREX_API_KEY={{ API_KEY }} BITTREX_API_SECRET={{ API_SECRET }} go run main.go`

