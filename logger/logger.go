package logger

import (
	"io"
	"log"
	"os"

	"github.com/printSANO/go-gin-example/config"
)

var (
	Clog *log.Logger // 커스텀 로깅 "C"ustom "log"ger
)

// init 함수는 패키지가 로드될 때 가장 먼저 실행되는 함수입니다.
func init() {
	if config.DEBUG_FLAG {
		log.Println("디버그 모드가 활성화되었습니다.")
		Clog = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		log.Println("디버그 모드가 비활성화되었습니다.")
		Clog = log.New(io.Discard, "", 0)
	}
}
