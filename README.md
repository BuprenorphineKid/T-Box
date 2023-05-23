# Box

Simple package to draw boxes in the terminal, written entirely in Go.

Import the go package for use in other go apps, or compile down to an
executable binary with 'go build .' and use in your shell scripts or as
processess with system calls in other languages.

I think ive got it running pretty good now, just try not to set
x or y to 0 or something negative, this is undefined.

# Install

git clone [this URL]

cd (new dir)

go build .

# Use

./(new binary) or go run main.go

with the following arguments

height, width, [offsetX, offsetY, style]