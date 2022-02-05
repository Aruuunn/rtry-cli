<h1 align="center">Rtry</h1>

<p align="center">(Re-)Trys executing a command till it succeeds.</p>

<p align="center">
 <a href="https://goreportcard.com/report/github.com/arunmurugan78/rtry"><img src="https://goreportcard.com/badge/github.com/arunmurugan78/rtry"/></a>
<a href="https://github.com/ArunMurugan78/rtry/blob/master/LICENSE" target="blank">
<img src="https://img.shields.io/github/license/ArunMurugan78/rtry?style=flat-square" alt="rtry licence" />
</a>
<a href="https://github.com/ArunMurugan78/rtry/fork" target="blank">
<img src="https://img.shields.io/github/forks/ArunMurugan78/rtry?style=flat-square" alt="rtry forks"/>
</a>
<a href="https://github.com/ArunMurugan78/rtry/stargazers" target="blank">
<img src="https://img.shields.io/github/stars/ArunMurugan78/rtry?style=flat-square" alt="rtry stars"/>
</a>
<a href="https://github.com/ArunMurugan78/rtry/issues" target="blank">
<img src="https://img.shields.io/github/issues/ArunMurugan78/rtry?style=flat-square" alt="rtry issues"/>
</a>
<a href="https://github.com/ArunMurugan78/rtry/pulls" target="blank">
<img src="https://img.shields.io/github/issues-pr/ArunMurugan78/rtry?style=flat-square" alt="rtry pull-requests"/>
</a>
<img src="https://github.com/ArunMurugan78/rtry/actions/workflows/releaser.yml/badge.svg" />
</p>

>Note: Currently Only supports **Unix** Based systems.

## Installation

### Binary Download
Checkout the [releases](https://github.com/ArunMurugan78/rtry/releases/latest) to download the binary for your distribution.

### Install from source
```bash
 go install github.com/ArunMurugan78/rtry-cli@latest
```

## Usage
```
rtry-cli [OPTIONS] COMMAND
```

### Example
```bash
rtry-cli -timeout 3000 "ping google.com"
```

### Options
| flag         | Description | Default | Required    |
| :---         |    :----:   | :----:  |  ---: |
| timeout      | Retry timeout in milliseconds  | 1000  | false   |
| code         | Expected exit code | 0 | false |

## License
This project is licensed under the **MIT** License.
