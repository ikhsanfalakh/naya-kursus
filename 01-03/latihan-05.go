package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
### Latihan-2
Buatlah program yang membalikan semua kata dalam sebuah string. Setiap kata di pisahkan dengan tepat satu spasi dan tidak ada spasi di awal ataupun di akhir string.

```
"The greatest victory is that which requires no battle" --> "battle no requires which that is victory greatest The"
```
*/

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Printf("Inputkan Kalimat = ")
	inputKalimat, _ := in.ReadString('\n')
	inputKalimat = strings.ReplaceAll(inputKalimat, "\n", "")

	kata := strings.Split(inputKalimat, " ")
	var kalimatBalik []string
	pjgKata := len(kata) - 1
	for i := pjgKata; i >= 0; i-- {
		kalimatBalik = append(kalimatBalik, kata[i])
	}
	fmt.Println(strings.Join(kalimatBalik, " "))
}
