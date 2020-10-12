```bash
$ docker run -it -d --name testnet -v $(pwd)/finalgravity:/finalgravity -p 12345:12345 -p 1317:1317 -p 26657:26657 -p 8080:8080 -p 26656:26656 finalgravity
$ docker exec -it testnet bash
$ make install

$ appd init beer_master_node --chain-id beer
$ appcli config chain-id beer
$ appcli config output json
$ appcli config indent true
$ appcli config trust-node true
$ appcli config keyring-backend test
$ appcli keys add snappytux

$ appd add-genesis-account $(appcli keys show snappytux -a) 100000000rune,100000000gold,100000000stake
$ appd gentx --name snappytux --keyring-backend test
$ appd collect-gentxs
$ appd validate-genesis

# Open new tab
$ docker exec -it testnet bash
$ appcli tx brewer create-brewer "Home Brew" \
"44/261 Passorn Onnut Prawet Prawet Bangkok 10250" \
"+66925905444" \
"aofiee666@gmail.com" \
"Punk IPA is the beer that kick-started it. This light, golden classic has been subverted with new world hops to create an explosion of flavour. Bursts of caramel and tropical fruit with an all-out riot of grapefruit, pineapple and lychee, precede a spiky bitter finish. This is the beer that started it all - and it’s not done yet... PUNK - Quintessential Empire with an anarchic twist." \
"https://www.brewdog.com/static/version1600847552/frontend/Born/arcticFox/en_US/images/logo.svg" \
"https://www.brewdog.com/static/version1600847552/frontend/Born/arcticFox/en_US/images/logo.svg" \
"2018" \
"Khomkrid Lerdprasert" \
"https://www.aofiee.dev" --from=snappytux
$ appcli rest-server
$ curl http://localhost:1317/brewer/create
```
