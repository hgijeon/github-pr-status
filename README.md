# GitHub PR Status

[![Go Report Card](https://goreportcard.com/badge/github.com/hgijeon/github-pr-status)](https://goreportcard.com/report/github.com/hgijeon/github-pr-status)
[![Build Status](https://github.com/hgijeon/github-pr-status/actions/workflows/release.yaml/badge.svg)](https://github.com/hgijeon/github-pr-status/actions/workflows/release.yaml)
[![Release](https://img.shields.io/github/release/hgijeon/github-pr-status.svg)](https://github.com/hgijeon/github-pr-status/releases)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

`github-pr-status`는 GitHub Pull Request 상태를 조회하고 요약하는 도구입니다.

## 기능

- GitHub Pull Request 목록 조회 (해당 PR 링크 포함)
- 요약된 PR 상태 출력 (조회 링크 포함)

## 설치

### 바이너리 다운로드

릴리즈 페이지에서 [최신 릴리즈](https://github.com/hgijeon/github-pr-status/releases) 바이너리를 다운로드하고 압축을 푼 후 실행 가능한 파일을 실행합니다.

### 소스에서 설치

```sh
git clone https://github.com/hgijeon/github-pr-status.git
cd github-pr-status
go build -o github-pr-status
```

### 소스에서 실행

```sh
git clone https://github.com/hgijeon/github-pr-status.git
cd github-pr-status
go run . [ARGUMENTS]
```

## 사용법

### 도움말

```sh
./github-pr-status --help
```

### PR 상태 조회

```sh
./github-pr-status --token [YOUR_GITHUB_TOKEN] --users hgijeon,torvalds --verbose
```

결과 예시

```
### hgijeon ###

Created (TOP 10)
UPDATED AT  TITLE  USER

Requested(not draft) (TOP 10)
UPDATED AT  TITLE  USER


### torvalds ###

Created (TOP 10)
UPDATED AT  TITLE  USER

Requested(not draft) (TOP 10)
UPDATED AT           TITLE                                                           USER
2024-07-27 01:53:10  Update libdivecomputer with changes in upstream as of 20240725  mikeller


### PR Counts Summary ###
USER      CREATED  REQUESTED(NOT DRAFT)
hgijeon   0        0
torvalds  0        1
```
