package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/sago35/go-jisx0208"
	"github.com/sago35/tinyufont"
	"github.com/zachomedia/go-bdf"
)

type fontgen struct {
	pkgname  string
	fontname string
	isofont  string
	font     string
}

func (f *fontgen) generate(w io.Writer, runes []rune) error {
	if len(f.isofont) == 0 {
		return fmt.Errorf("length error (isofont)")
	}

	isoFont, err := readFont(f.isofont)
	if err != nil {
		return err
	}

	font, err := readFont(f.font)
	if err != nil {
		return err
	}

	ufont := tinyufont.Font{}
	ufont.YAdvance = uint8(font.Size)
	for _, f := range isoFont.Characters {
		img := f.Alpha

		bmp := []byte{}
		bitpos := 0
		for y := img.Rect.Min.Y; y < img.Rect.Max.Y; y++ {
			b := byte(0)
			for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
				r, _, _, _ := img.At(x, y).RGBA()
				if r != 0 {
					b += (byte(1) << (8 - bitpos - 1))
				}

				bitpos++
				if bitpos == 8 {
					bitpos = 0
					bmp = append(bmp, b)
					b = 0
				}
			}

			if bitpos != 0 {
				bitpos = 0
				bmp = append(bmp, b)
			}
		}

		g := tinyufont.Glyph{
			Width:    uint8(img.Rect.Max.X),
			Height:   uint8(img.Rect.Max.Y),
			XAdvance: uint8(f.Advance[0]),
			XOffset:  int8(f.LowerPoint[0]),
			YOffset:  int8(f.LowerPoint[1]),
			Bitmaps:  bmp,
		}
		ufont.Glyphs = append(ufont.Glyphs, g)
		ufont.RuneToIndex = append(ufont.RuneToIndex, tinyufont.RuneToIndex{
			Rune:  f.Encoding,
			Index: uint16(len(ufont.Glyphs) - 1),
		})
	}

	for _, r := range runes {

		c, err := jisx0208.Code(r)
		if err != nil {
			return err
		}

		for _, f := range font.Characters {
			if c == int(f.Encoding) {
				img := f.Alpha

				bmp := []byte{}
				bitpos := 0
				for y := img.Rect.Min.Y; y < img.Rect.Max.Y; y++ {
					b := byte(0)
					for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
						r, _, _, _ := img.At(x, y).RGBA()
						if r != 0 {
							b += (byte(1) << (8 - bitpos - 1))
						}
						bitpos++
						if bitpos == 8 {
							bitpos = 0
							bmp = append(bmp, b)
							b = 0
						}
					}
					if bitpos != 0 {
						bitpos = 0
						bmp = append(bmp, b)
					}
				}

				g := tinyufont.Glyph{
					Width:    uint8(img.Rect.Max.X),
					Height:   uint8(img.Rect.Max.Y),
					XAdvance: uint8(f.Advance[0]),
					XOffset:  int8(f.LowerPoint[0]),
					YOffset:  int8(f.LowerPoint[1]),
					Bitmaps:  bmp,
				}
				ufont.Glyphs = append(ufont.Glyphs, g)
				ufont.RuneToIndex = append(ufont.RuneToIndex, tinyufont.RuneToIndex{
					Rune:  r,
					Index: uint16(len(ufont.Glyphs) - 1),
				})
			}
		}
	}

	fmt.Fprintln(w, `package `+f.pkgname)
	fmt.Fprintln(w, ``)
	fmt.Fprintln(w, `import (`)
	fmt.Fprintln(w, `	"github.com/sago35/tinyufont"`)
	fmt.Fprintln(w, `)`)
	fontname := strings.ToUpper(f.fontname[0:1]) + f.fontname[1:]
	fmt.Fprintf(w, "var %s = %T{\n", fontname, ufont)
	fmt.Fprintf(w, "	Glyphs:%T{\n", ufont.Glyphs)
	for i, g := range ufont.Glyphs {
		fmt.Fprintf(w, "		/* %c */ %#v,\n", ufont.RuneToIndex[i].Rune, g)
	}
	fmt.Fprintf(w, "	},\n")
	fmt.Fprintf(w, "	RuneToIndex:%T{\n", ufont.RuneToIndex)
	for _, rti := range ufont.RuneToIndex {
		fmt.Fprintf(w, "		/* %c */ %#v,\n", rti.Rune, rti)
	}
	fmt.Fprintf(w, "	},\n")
	fmt.Fprintf(w, "	YAdvance:%#v,\n", ufont.YAdvance)
	fmt.Fprintf(w, "}\n")

	return nil
}

func readFont(p string) (*bdf.Font, error) {
	buf, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}

	font, err := bdf.Parse(buf)
	if err != nil {
		return nil, err
	}

	return font, nil

}
