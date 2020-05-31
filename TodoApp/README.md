# Heroku를 통한 배포하는 방법

## 배포 순서

PORT 번호를 heroku에서 지정해줘야한다.
`port := os.Getenv("PORT")` main함수에 적용
Procfile을 만들어준다.
내용: `web: bin/todos`

1. heroku 설치
   `brew install heroku/brew/heroku`
2. 파일 모듈화
   `go mode init PROJECT_NAME`
3. 파일 확인
   `go build -o ./bin/todos -v .`
   - 빌드를 시작한다.
   - output은 현재폴더안에 bin/todos로 지정한다.
   - 빌드되는내용을 보여준다.
   - 현재폴더를 build 하겠다.
4. heroku 로그인
   `heroku login`
5. git 연결
   `heroku create`
6. git add/commit/push
   `git add .`
   `git commit -m "heroku deploy"`
   `git push heroku master`
7. 웹 확인
   `heroku ps`
   `heroku open`
8. 로그 확인
   `heroku logs --tail`

## 환경설정을 지정해주기 위한 부분

`heroku config:set [key]:[value]`

## 배포과정에서 겪은 오류

- 빌드하기 위해 module 화를 함 `$go mod todos`
- 계속 발생하는 오류메세지가 경로를 찾을 수 없다고 나옴.
- `$GOROOT` 및 `$GOPATH` 를 잘못 지정한 문제라고 생각하고 env를 계속 변경하는 시도를 했었음
- 해결되지 않아 파일의 위치를 `$GOROOT` 및 `$GOPATH` 안에서 이동하고 모듈위치를 변경하고 시도를 했었음
- 해결되지 않아 `go.mod` 파일 안에 강제로 입력하고 배포를 시도함
- 배포는 성공했으나 배포된 파일에서 경로를 인식하지 못하는 문제가 발생함
- 다시 돌아와 heroku에서 문제를 찾아보도록 검색을 함
- 검색을 통해 `godep`이라는 tool을 이용해 local package를 연결 한다고 나옴
- 설치 후 `godep.json`을 만들고 시도를 함
- 해결되지 않아 go mod 를 다시 돌아와서 찾아 보니 설정한 이름으로 패키지위치를 변경해야 한다고 나옴
- 해결함
- mod를 통해 packaging하는 과정에서 설정한 이름으로 내부패키지를 변경해줘야하는 부분을 알지 못해 삽질을 많이함
