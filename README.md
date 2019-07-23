The app exports one endpoint /v1/ordering?limit= where the ordering refers to the Top performers in the last 24 hours by cryptocompare api. The price for the coins are retrieved from the Coinmarketcap API. 

Please export the config parameters CoinMarket_API_Key, CoinMarket_URL, CryptoCompare_API_Key, CryptoCompare_URL 

make install: will setup the docker and run the app at port 8080