# buildinfo

Extract build information from Go binary

## Usage

Mostly used to update binaries that were install with `go install`.

```
buildinfo $(which pkgsite) | jq -r '.Path' | xargs -I {} go install "{}@latest"
```

## Tips

- Find the origin of the binary: `buildinfo $(which hugo) | jq -r '.Main.Path'`
- Find the version of Go the binary was compiled with: `buildinfo $(which gopls) | jq -r '.GoVersion'`
- Find the version of the binary: `buildinfo $(which pkgsite) | jq -r '.Main.Version'`
- Filter dependencies: `buildinfo $(which hugo) | jq -r -c '.Deps.[]' | grep yaml`

Noteworthy:

- Check for vulnerable binaries: `govulncheck -mode=binary $(which hugo)`

