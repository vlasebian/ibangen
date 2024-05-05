# ibangen

[![semantic-release: angular](https://img.shields.io/badge/semantic--release-angular-e10079?logo=semantic-release)](https://github.com/semantic-release/semantic-release)

A CLI tool written in Golang for generating IBANs. For testing purposes only.

# Installation

```bash
go install github.com/vlasebian/ibangen@v1.0.0
```

# Usage

Supported country codes:

- 'nl' (Netherlands)
- 'ie' (Ireland)
- 'at' (Austria)
- 'ch' (Switzerland)
- 'es' (Spain)
- 'it' (Italy)
- 'de' (Germany)
- 'be' (Belgium)
- 'fr' (France)

```bash
ibangen --help

# Country code is optional, a random country code will be used if not specified.
ibangen nl
```
