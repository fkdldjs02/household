# Go언어로 구현한 가계부 API 서버

* 실행 요구사항
    * golang 1.8 버전 이상(64bit)
    * gcc(64bit)
    * glide 패키지 관리자
* 설치 방법
   `glide install`
* 코드 구성
    * conf/config.go
        * 데이터베이스 생성 정보를 담고 있음
    * api/context.go
        * api의 컨트롤러 부분
    * api/model/**
        * api의 모델 부분, 실질적인 데이터베이스 작업 부분
    * api/test/api_test.go
        * api의 사용법을 구현
* 시연 방법
    * test case를 이용한 시연
        * tests/ 폴더로 이동
        * `go test -v` 명령어로 실행


