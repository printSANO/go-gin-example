# Go 언어로 만드는 Rest API Boilerplate

> 이 레포는 테커 GoodNight Hackathon 3차를 위해 만들어진 Boilerplate 레포 입니다

## 사용된 프레임워크와 기술들
- Go 1.23
- Go Gin
- Gorm
- PostgreSQL
- Docker

## 프로젝트 파일 디렉토리
```
📦go-gin-example //프로젝트
 ┣ 📂config //설정 관련 디렉토리(환경변수, 등등..)
 ┃ ┗ 📜env.go //환경변수 함수 파일
 ┣ 📂database //데이터베이스 관련 디렉토리
 ┃ ┗ 📜db.go //디비 관련 함수
 ┣ 📂handlers //핸들러 디렉토리
 ┃ ┣ 📜handler.go //핸들러 총괄 파일
 ┃ ┣ 📜post_handler.go //포스트 관련 핸들러
 ┃ ┗ 📜user_handler.go //유저 관련 핸들러
 ┣ 📂models //모델 디렉토리
 ┃ ┣ 📜post.go //포스트 관련 모델
 ┃ ┗ 📜user.go //유저 관련 모델
 ┣ 📂repositories //레포지토리 디렉토리
 ┃ ┣ 📜post_repository.go //포스트 관련 레포지토리
 ┃ ┣ 📜respository.go //레포지토리 총괄 파일
 ┃ ┗ 📜user_repository.go //유저 관련 레포지토리
 ┣ 📂server //서버 및 라우터 디렉토리
 ┃ ┣ 📜router.go //라우터 설정
 ┃ ┗ 📜server.go //서버 시작 관련 파일
 ┣ 📂services //서비스 레이어 디렉토리
 ┃ ┣ 📜post_service.go //포스트 관련 서비스
 ┃ ┣ 📜service.go //서비스 총괄 파일
 ┃ ┗ 📜user_service.go //유저 관련 서비스
 ┣ 📜.env //환경변수
 ┣ 📜.gitignore
 ┣ 📜Dockerfile //도커파일
 ┣ 📜LICENSE
 ┣ 📜README.md
 ┣ 📜docker-compose.yml //도커 컴포즈 파일. 디비 및 디비 뷰어 포함
 ┣ 📜go.mod //고 모듈 관리
 ┣ 📜go.sum //모듈 관리
 ┗ 📜main.go //실행 시킬 메인 함수 파일
 ``` 

## 고 (Go Lang) 설치 및 설정법
1. `brew install go` 로 설치 아니면 `https://go.dev/doc/install` 에서 설치
2. `go env` 로 설치 확인

## 프로젝트 설정법
1. 터미널에 `go mod init github.com/<깃허브 이름 아니면 organization 이름>/<레포 이름>` 으로 하면`go.mod`와 `go.sum`이 생깁니다.
2. 패키지 설치를 원하면 go get -u <패키지이름>
3. 그 외 사용한 패키지를 삭제, 업데이트, 또는 사용된 패키지를 자동으로 설치 하게 하고 싶으면 `go mod tidy`해주면 됩니다
- `go mod tidy`는 자동으로 소스 코드에 사용되는 라이브러리와 `go.mod` 와 `go.sum`을 싱크 해줍니다.
