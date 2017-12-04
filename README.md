# Simple Blockchain algorithm with GOLang

## GO version
`go version go1.9 darwin/amd64`

## **STEPS:**

1. **Download the extra packages:**
    1. `go get golang.org/x/crypto/ripemd160`
    2. `go get github.com/boltdb/bolt`

2. **Build the application:**
    1. `go build *.go`, this generate a file called **base58**.

3. **Execute base58 to start using the cli**

4. **To using the cli use the next command structure `./bae58 --commands`**

## **COMANDS**
|Command|Description|
|----|----|
| `createblockchain -address ADDRESS`|**Create a blockchain and send genesis block reward to ADDRESS**|
| `createwallet`|**Generates a new key-pair and saves it into the wallet file**|
| `getbalance -address ADDRESS`|**Get balance of ADDRESS**|
| `listaddresses`|**Lists all addresses from the wallet file**|
| `printchain`|**Print all the blocks of the blockchain**|
| `send -from FROM -to TO -amount AMOUNT`|**Send AMOUNT of coins from FROM address to TO**|
