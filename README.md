# Simple Blockchain algorithm with GOLang

## GO version
`go version go1.9 darwin/amd64`

|Command|Description|
|----|----|
| `createblockchain -address ADDRESS`|**Create a blockchain and send genesis block reward to ADDRESS**|
| `createwallet`|**Generates a new key-pair and saves it into the wallet file**|
| `getbalance -address ADDRESS`|**Get balance of ADDRESS**|
| `listaddresses`|**Lists all addresses from the wallet file**|
| `printchain`|**Print all the blocks of the blockchain**|
| `send -from FROM -to TO -amount AMOUNT`|**Send AMOUNT of coins from FROM address to TO**|
