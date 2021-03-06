# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], and this project adheres to
[Semantic Versioning].

<!-- references -->
[Keep a Changelog]: https://keepachangelog.com/en/1.0.0/
[Semantic Versioning]: https://semver.org/spec/v2.0.0.html

## [0.4.0] - 2019-02-06

### Added

- Add `Count64()` method to `count.Reader` and `Writer`

### Changed

- **[BC]** Move `iago.MustXXX()` functions to `must.XXX()`

## [0.3.0] - 2019-01-29

### Added

- Add `count.Reader` and `Writer`

## [0.2.0] - 2019-01-29

### Added

- Add `ioutil.TestWriteXXX()` helpers for testing write operations
- Add `MustWriteByte()`

### Fixed

- `Indenter.Write()` no longer includes the indent length in the returned byte count

## [0.1.0] - 2019-01-16

- Initial release

<!-- references -->
[Unreleased]: https://github.com/dogmatiq/iago
[0.1.0]: https://github.com/dogmatiq/iago/releases/tag/v0.1.0
[0.2.0]: https://github.com/dogmatiq/iago/releases/tag/v0.2.0
[0.3.0]: https://github.com/dogmatiq/iago/releases/tag/v0.3.0

<!-- version template
## [0.0.1] - YYYY-MM-DD

### Added
### Changed
### Deprecated
### Removed
### Fixed
### Security
-->
