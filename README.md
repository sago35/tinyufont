# TinyUFont

TinyUFont is a font/text package for TinyGo displays.  
This repository is for making PR for tinyfont.  


![](tinyufont.jpg)

## VS. github.com/tinygo-org/tinyfont

* TinyGo font files can be created from BDF fonts (tinyufontgen)
* Unicode font support
  * Tested with jisx0208 (Japanese) font only
* Usage is the same as tinygo-org/tinyfont

## Description

Generate a font file using `tinyufontgen` described later.  
Codepoints <= 255 are always included in the font.  
For fonts with a code point of 256 or more, only those specified by the argument are included in the font.  

## Usage

    // https://mplus-fonts.osdn.jp/mplus-bitmap-fonts/download/index.html
    tinyufontgen --ascii-font mplus_j10b-iso.bdf --font mplus_j10r.bdf --output font.go 日本語

    // http://x-tt.osdn.jp/ayu.html
    tinyufontgen --ascii-font 10x20grkm.bdf --font k20gm.bdf --output font.go 日本語

## Installation

    go get github.com/sago35/tinyufont/cmd/tinyufontgen

### Environment

* go
* kingpin.v2

## Notice

None.

## License

Software License Agreement (BSD License)

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

- Redistributions of source code must retain the above copyright notice,
  this list of conditions and the following disclaimer.
- Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
POSSIBILITY OF SUCH DAMAGE.

## Author

sago35 - <sago35@gmail.com>
