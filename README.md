# go-args-parser

## Getting started

### download

```bash
go get -v github.com/Sotaneum/go-args-parser
```

### API

| name        | params              | returns             | example                                 | -                                                                                           |
| ----------- | ------------------- | ------------------- | --------------------------------------- | ------------------------------------------------------------------------------------------- |
| Args        | `map[string]string` | `map[string]string` | parser.Args(map[string]string{})        | Argument로 넘어오는 값을 `map[string]string` 포맷으로 반환합니다.                           |
| EnvAll      | `map[string]string` | `map[string]string` | parser.EnvAll(map[string]string{})      | 모든 환경 변수 값을 `map[string]string` 포맷으로 반환합니다.                                |
| Env         | `map[string]string` | `map[string]string` | parser.Env(map[string]string{})         | 환경 변수 값 중 `params`에 해당하는 key만 추출하여 `map[string]string` 포맷으로 반환합니다. |
| ArgsJoinEnv | `map[string]string` | `map[string]string` | parser.ArgsJoinEnv(map[string]string{}) | `params`에 해당 하는 key만 추출하여 env로 부터 값을 가져오고 그 위에 Args를 덮어쓰기합니다. |

## LICENSE

[MIT](./LICENSE)
