# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [0.2.0] - 2026-04-12

### Security
- Fixed race condition in `buildGenerator`: `cachedGenerators` is now protected by a `sync.Mutex`, making concurrent use of the library safe (`qrencode-go/errorcorrection.go`)
- Added `recover()`-based panic handler in `Encode` and `EncodeBytes`: internal panics are now returned as `error` values instead of crashing the caller (`qrencode-go/encode.go`)
- Fixed ignored error return of `png.Encode` in `main.go`: errors are now printed to stderr and cause a non-zero exit

### Changed
- Replaced magic index `1` in `totalCodewords()` with named constant `ECLevelL` and added an explanatory comment (`qrencode-go/version.go`)

### Fixed
- Fixed ignored error return of `png.Encode` in `ExampleEncode` test (`qrencode-go/qr_test.go`)
