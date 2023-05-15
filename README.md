# signature
signature is a client implemented in Go that can query signatures from [4byte.directory](https://www.4byte.directory/), [ethereum-lists](https://github.com/ethereum-lists/4bytes), [openchainxyz](https://openchain.xyz/signatures). <br>
`e.g. "0xa9059cbb" -> transfer(address,uint256)` <br>


Actually, 'ethereum-lists' is tied to '4byte.directory', and the difference is that '4byte.directory' is more real-time('ethereum-lists' has fewer signatures than '4byte.directory'). However, it's better for API users to have more endpoint options than one (4byte.directory, github.com, openchain.xyz).

### They has (2023-05-15)
ethereum-lists: `915,173` <br>
4byte.directory: `1,210,015` <br>
openchain: method `2,361,806`, event `372,441` <br>

# Usage

## Install
```bash
$ go get -u github.com/dbadoy/signature
# OR `go mod tidy`
$ go get -u github.com/dbadoy/signature/file 
$ go get -u github.com/dbadoy/signature/fourbytes
$ go get -u github.com/dbadoy/signature/openchain
```

### file client
Get the signature from the [ethereum-lists](https://github.com/ethereum-lists/4bytes) repository.

```go
package main

import (
	"fmt"

	"github.com/dbadoy/signature/file"
)

func main() {
	client, err := file.New(0)
	if err != nil {
		panic(err)
	}

	// [transfer(address,uint256)] <nil>
	fmt.Println(client.Signature("0xa9059cbb"))
	fmt.Println(client.Signature("a9059cbb"))
}

```

### fourbytes client
Get the signature from the [4byte.directory](https://www.4byte.directory/) API.

```go
package main

import (
	"fmt"

	"github.com/dbadoy/signature/fourbytes"
)

func main() {
	client, err := fourbytes.New("", 0)
	if err != nil {
		panic(err)
	}

	// [join_tg_invmru_haha_fd06787(address,bool) func_2093253501(bytes) transfer(bytes4[9],bytes5[6],int48[11]) many_msg_babbage(bytes1) transfer(address,uint256)] <nil>
	fmt.Println(client.Signature("0xa9059cbb"))
	fmt.Println(client.Signature("a9059cbb"))
}
```

### openchain client
Get the signature from the [openchainxyz](https://openchain.xyz/signatures) API.


```go
package main

import (
	"fmt"

	"github.com/dbadoy/signature/openchain"
)

func main() {
	client, err := openchain.New("", 0)
	if err != nil {
		panic(err)
	}
    
	// [transfer(address,uint256)] <nil>
	fmt.Println(client.Signature("0xa9059cbb"))
	fmt.Println(client.Signature("a9059cbb"))
}
```
