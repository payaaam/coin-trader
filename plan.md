// CMD LINE UTIL takes exchange flag

// Get all Markets
// For each BTC Market
  - Get Candles
  - Create Cloud
  - Determine if a cross happened in the last 5 days
  - Determine if a cross will happen in the next few days
  - Print 
    * Kumo Breakouts
    * TK Crosses
    * If price is at Kijun line



Auto trading



// exchange interface
GetCandles
GetMarketPairs
ExecuteLimitTrade
ExecutLimitSell
GetHoldings


// Create Charts from Database
// Get All OPen Balances and enter into trading object


// Actor Model
// Create channel that gets the btc value

// Trading Object
```javascript
 {
   balances: {
     neo: 99.45,
     btc: 0.031234,
     ltc: 0.0,
   },
   openOrders: [
     {
       type: 'buy',
       marketKey: 'btc-eth',
       quantity: 1.5,
       limi: 0.1112345,
       uuid: 'asdf'
     },
     {
       type: 'sell',
       marketKey: 'btc-neo',
       quantity: 2,
       limi: 0.20000,
       uuid: 'asdf'
     },
   ],
   charts: {
    neo: {
      chart: CloudChart,
      interval: '1h',
      active: true,
    },
    btc: {
      chart: CloudChart,
      interval: '1h',
      active: true
    },
    ltc: {
      chart: CloudChart,
      interval: '1h',
      active: false,
    }
   }
 }
 ```

```go
// Check Open positions
if state.balances[neo] != "" {
  data:= readGraph("neo")
  if bearishTKCross() == true {
    // Update balances
    // Update Database

    executeSell("neo")
    // Tak eorder  and add it to open orders array
  }
}

if state.balances[neo] == "" {
  data:= readGraph("neo")
  if bullishTKCross() == true {
    executeBuy("neo")
    // take order and add it to open orders aray
  }
}
```