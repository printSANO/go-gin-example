// 환경변수를 로드하고 설정하는 패키지입니다.
package config

var (
	DEBUG_FLAG bool   // 디버그 모드 활성화 여부. true: 활성화, false: 비활성화. (기본값: true). 디버그 모드가 활성화되면 로그가 출력됩니다.
	DBURL      string // 데이터베이스 URL
)

// init 함수는 패키지가 로드될 때 가장 먼저 실행되는 함수입니다.
func init() {
	loadEnvFile(".env")
	DEBUG_FLAG = getEnvVarAsBool("DEBUG_FLAG", true)
	DBURL = getEnvVarAsString("DATABASE_URL", "")
}
