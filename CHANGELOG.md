# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [0.2.1] - 2026-10-31

### Security
- Raised the required Go version from 1.21.1 to 1.26.10 to build with a patched standard library (`go.mod`)
- Fixed the module path to `github.com/OZON08/qrencode`, so consumers no longer resolve to the upstream repository without the 0.2.0 fixes (`go.mod`)

### Migration
- v0.2.1 is the first version installable under `github.com/OZON08/qrencode`. Projects that used a `replace` directive for `github.com/d3vilh/qrencode` must remove it and update their imports.

### Changed
- The CLI prints encoding errors to stderr and exits with code 1 instead of panicking with a stack trace and exit code 2 (`main.go`)
- CLI error prefixes are available in German and English, selected through `LC_ALL`, `LC_MESSAGES` or `LANG`. Texts from the library stay English (`main.go`)

### Fixed
- Set the empty Dependabot `package-ecosystem` to `gomod` (`.github/dependabot.yml`)

## [0.2.0] - 2026-04-12

### Security
- Fixed race condition in `buildGenerator`: `cachedGenerators` is now protected by a `sync.Mutex`, making concurrent use of the library safe (`qrencode-go/errorcorrection.go`)
- Added `recover()`-based panic handler in `Encode` and `EncodeBytes`: internal panics are now returned as `error` values instead of crashing the caller (`qrencode-go/encode.go`)
- Fixed ignored error return of `png.Encode` in `main.go`: errors are now printed to stderr and cause a non-zero exit

### Changed
- Replaced magic index `1` in `totalCodewords()` with named constant `ECLevelL` and added an explanatory comment (`qrencode-go/version.go`)

### Fixed
- Fixed ignored error return of `png.Encode` in `ExampleEncode` test (`qrencode-go/qr_test.go`)
