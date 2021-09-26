# CC-Gen

Simple golang cli to generate a changelog from your repository using
conventional commits.

This cli was mainly written for internal use, but is okay to use and/or contribute to as much as wanted!

## Usage

### Binary usage

Download the proper version from https://gitlab.com/jitesoft/open-source/go/cc-gen/-/packages
Extract the binary and add to your path.

Usage is quite simple, all you need to do is run the executable:

```
cc-gen [optional path to repo or . ] [optional version name]
```

Make sure to check out `cc-gen --help` for information about flags.

### Docker usage

The application is also distributed as docker images one based on Alpine while the other
is based on scratch.

A /data directory is available to easily mount your repository and just run the image:

```sh 
docker run --rm registry.gitlab.com/jitesoft/open-source/go/cc-gen:latest -v $(pwd):/data generate [args] 
```

## Building

This application is written in go, building it is as easy as:

```sh
go mod download
go build
```

Or if you wish to build for all common platforms:

```sh 
./build.sh [version]
```

No extra stuff required.

## Libraries used

The cli uses the following libraries as direct dependencies:

* `github.com/go-git/go-git/v5`
* `github.com/spf13/cobra`

Check out the `go.sum` file for a full list of dependencies.

## License

```text
MIT License

Copyright (c) 2020 Jitesoft

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
