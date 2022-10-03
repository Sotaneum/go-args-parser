# go-args-parser

[![GitHub release](https://img.shields.io/github/release/Sotaneum/go-args-parser.svg)](https://GitHub.com/Sotaneum/go-args-parser/releases/)
[![Go](https://img.shields.io/badge/--00ADD8?logo=go&logoColor=ffffff)](https://golang.org/)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/Sotaneum/go-args-parser.svg)](https://github.com/Sotaneum/go-args-parser)
[![GoReportCard example](https://goreportcard.com/badge/github.com/Sotaneum/go-args-parser)](https://goreportcard.com/report/github.com/Sotaneum/go-args-parser)
[![GitHub license](https://badgen.net/github/license/Sotaneum/go-args-parser)](https://github.com/Sotaneum/go-args-parser/blob/master/LICENSE)

환경변수나 인자로 전달받는 내용을 `map[string]string` 포맷으로 변경합니다.

## Getting started

### download

```bash
go get -v github.com/Sotaneum/go-args-parser
```

### API

| name        | params & returns    | info                                                                                        |
| ----------- | ------------------- | ------------------------------------------------------------------------------------------- |
| Args        | `map[string]string` | Argument로 넘어오는 값을 `map[string]string` 포맷으로 반환합니다.                           |
| EnvAll      | `map[string]string` | 모든 환경 변수 값을 `map[string]string` 포맷으로 반환합니다.                                |
| Env         | `map[string]string` | 환경 변수 값 중 `params`에 해당하는 key만 추출하여 `map[string]string` 포맷으로 반환합니다. |
| ArgsJoinEnv | `map[string]string` | `params`에 해당 하는 key만 추출하여 env로 부터 값을 가져오고 그 위에 Args를 덮어쓰기합니다. |

## LICENSE

[MIT](./LICENSE)
