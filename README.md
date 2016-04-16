# TDD로 시작하는 GO

## 서문

제가 GO 언어를 처음 만난것은 얼마되지 않습니다. GO를 접하면서 상당히 많은 부분에서 그동안 언어에 몰랐던 새로운 지식을 습득하게 되었습니다.

Agile Practice에서 개발자가 가져야 할 덕목중 하나가 매년 한 개의 언어에 대해 숙달하는 것입니다. GO 언어를 처음 접하게 된 동기는 이를 실천하려는 것 중의 하나였습니다.

새로운 프로그램 언어를 처음 배우는 과정에서 흔한  Bible 류 같은 서적을 참고하여 배울 수도 있습니다. GO에서도 그런 참고서적을 찾을 수도 있을 겁니다. 또 GO에서 자체 제공하는 PlayGround 역시 훌륭한 지침서가 됩니다.

한가지 문제는 집중하기 꽤나 어렵다는 점이고 실제 동작하는 코드의 완성도에 대한 걱정이 앞섰습니다.

이때 오래전 읽었던 켄트 벡의 "테스트 주도 개발"(이하 TDD)가 생각났습니다. 그간의 경험상 언어를 배우는 가장 효과적인 방법은 테스트 프레임워크를 만들어 보는 것이라고 저는 생각합니다. 이를 접목해서 GO 언어를 배우려고 하는 초심자들을 위한 TDD 가이드가 있었으면 좋겠다는 생각을 했습니다.

물론 인터넷에 훌륭한 교제들이 있었습니다. 그중 하나가 Go TDD 였습니다.

* [Test-driven development with Go](https://leanpub.com/golang-tdd/read#leanpub-auto-test-driven-development)

이 문서 하나로도 충분하지만 한글로 번역되어 있지도 않고 테스트 코드가 한정되어 있어서 TDD 의 그 짧은 단계들을 다 심도있게 추적하지 못한다는 생각을 갖게 되었고, 이것이 제가 이 글을 쓰게된 이유입니다.

이글은 철저하게 TDD를 기반으로 합니다. 1장의 테스트 순서도 모두 이책의 내용을 따라 갑니다. 저자나 번역본의 출판사에서 이의를 제기한다면 전 소송을 당할 수도 있을 겁니다.

전 철저하게 테스트 주도 개발 사상을 따라가고자 하는 의도로 작성했습니다.

## 이 글의 구성

이 글은 Money 패키지를 작성하는 과정을 다룹니다. 그 이후에 코드 품질을 위한 가이드를 추가적으로 제공합니다.

[Step01](./Step01/)

## Test 에 대해서
GO의 테스트 방법에 대해 먼저 알아야 할 것이 있습니다.

```go
go test
```
* 테스트 케이스 별로 테스트 결과를 보여준다.
```go
go test -v
```

## Code Quality

### Coverage Metrics
[블로그 Cover](http://blog.golang.org/cover) 에서 Coverage Metrics에 대한 상세한 설명을 찾아볼 수 있다.

* Coverage Profile 파일을 생성한다.

```bash
go test -coverprofile=coverage.out
```

* 함수 단위로 Coverage를 표시한다.

```bash
go tool cover -func=coverage.out
```

* html로 Coverage 보고서를 생성한다.

```bash
go tool cover -html=coverage.out
```

### 정적 분석

Documentation 기능을 활용한 정적 분석이 가능하다.

내부 메소드를 확인하려면

```bash
godoc -http=:6060 -analysis=type
```

* Call Graph 를 볼 수 있다.

godoc -http=:6060 -analysis=pointer

### Documentation

[블로그 Godoc](http://blog.golang.org/godoc-documenting-go-code)

`godoc`은 `javadoc` 과 같은 개념의 도구다.

```go
godoc -v softfactory.org/AgileGO
```

웹으로 패키지의 Documentation을 살펴볼 수 있다.
```go
godoc -http=:6060
```

https://golang.org/lib/godoc/analysis/help.html
