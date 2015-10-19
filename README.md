# warhol

Create simple warhol-like artworks from digital images.

![Obi Warhol](obi-warhol2-small.jpg)

## Installation

$ go get github.com/atongen/warhol

## Usage

```
$ warhol [OPTIONS] path/to/image.jpg

Options:
  -c="": Use custom color set, CSV of hex values
  -h=false: Print help and exit
  -l=false: Print color palettes and exit
  -o=".": Output directory
  -p="high": Select color palette
  -s=3: Size of output grid, valid values are 3 (3x3), 2 (2x2), or 0 (do not assemble final image)
  -v=false: Verbose output
  -version=false: Print version and exit
  -w=8: Number of workers for processing
```

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
