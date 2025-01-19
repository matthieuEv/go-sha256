# go-sha256

This is my implementation of the SHA-256 algorithm in Go. This is a simple implementation that is not optimized for performance. It is meant to be a learning exercise for me to understand how the SHA-256 algorithm works and to learn Go.

I had the idea to implement this after finding this website that explains the SHA-256 algorithm in a simple way: [SHA-256 Algorithm](https://sha256algorithm.com/)

## Usage

An example file is provided here: [example.go](https://github.com/matthieuEv/go-sha256/blob/main/example.go)

### Build it yourself

To build the project yourself, you need to have Go installed on your machine. You can download it from the [official website](https://golang.org/).

After you have Go installed, you can clone this repository and build the project by running the following commands:

```bash
git clone https://github.com/matthieuEv/go-sha256.git

cd go-sha256

go build -o go-sha256
```


### Use the pre-built binary

You can also use the pre-built binary that is available in the [releases section](https://github.com/matthieuEv/go-sha256/releases) of this repository. You can download the binary for your operating system and architecture and run it directly.

## Run

After you have built the project or downloaded the pre-built binary, you can run the program by passing the string you want to hash as an argument. You can pass multiple strings as arguments and the program will hash each string separately.

```bash
./go-sha256 hello # one argument
# hello :  2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824

./go-sha256 "this is a test" # a string with spaces
# this is a test :  2e99758548972a8e8822ad47fa1017ff72f06f3ff6a016851f45c398732bc50c

./go-sha256 hello "world here" # multiple arguments
# hello :  2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
# world here :  ab23b7e7a10390cd75956e465a5237834bab2bdce1853f261e80bccbf6788401
```

## Thanks

- [dmarman](https://github.com/dmarman) For his awesome website that explains the SHA-256 algorithm in a simple way.