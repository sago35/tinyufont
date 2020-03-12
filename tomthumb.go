/**
** The original 3x5 font is licensed under the 3-clause BSD license:
**
** Copyright 1999 Brian J. Swetland
** Copyright 1999 Vassilii Khachaturov
** Portions (of vt100.c/vt100.h) copyright Dan Marks
**
** All rights reserved.
**
** Redistribution and use in source and binary forms, with or without
** modification, are permitted provided that the following conditions
** are met:
** 1. Redistributions of source code must retain the above copyright
**    notice, this list of conditions, and the following disclaimer.
** 2. Redistributions in binary form must reproduce the above copyright
**    notice, this list of conditions, and the following disclaimer in the
**    documentation and/or other materials provided with the distribution.
** 3. The name of the authors may not be used to endorse or promote products
**    derived from this software without specific prior written permission.
**
** THIS SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR
** IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
** OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
** IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT,
** INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
** NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
** DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
** THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
** (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
** THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
**
** Modifications to Tom Thumb for improved readability are from Robey Pointer,
** see:
** http://robey.lag.net/2010/01/23/tiny-monospace-font.html
**
** The original author does not have any objection to relicensing of Robey
** Pointer's modifications (in this file) in a more permissive license.  See
** the discussion at the above blog, and also here:
** http://opengameart.org/forumtopic/how-to-submit-art-using-the-3-clause-bsd-license
**
** Feb 21, 2016: Conversion from Linux BDF --> Adafruit GFX font,
** with the help of this Python script:
** https://gist.github.com/skelliam/322d421f028545f16f6d
** William Skellenger (williamj@skellenger.net)
** Twitter: @skelliam
**
 */

package tinyufont

var TomThumb = Font{
	Glyphs: []Glyph{
		/*   */ Glyph{Width: 0x8, Height: 0x1, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x0}},
		/* ! */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80, 0x80, 0x0, 0x80}},
		/* " */ Glyph{Width: 0x8, Height: 0x2, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0}},
		/* # */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xe0, 0xa0, 0xe0, 0xa0}},
		/* $ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xc0, 0x60, 0xc0, 0x40}},
		/* % */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x20, 0x40, 0x80, 0x20}},
		/* & */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xc0, 0xe0, 0xa0, 0x60}},
		/* ' */ Glyph{Width: 0x8, Height: 0x2, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80}},
		/* ( */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0x80, 0x80, 0x40}},
		/* ) */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x40, 0x40, 0x40, 0x80}},
		/* * */ Glyph{Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x40, 0xa0}},
		/* + */ Glyph{Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x40, 0xe0, 0x40}},
		/* , */ Glyph{Width: 0x8, Height: 0x2, XAdvance: 0x3, XOffset: 0, YOffset: -2, Bitmaps: []uint8{0x40, 0x80}},
		/* - */ Glyph{Width: 0x8, Height: 0x1, XAdvance: 0x4, XOffset: 0, YOffset: -3, Bitmaps: []uint8{0xe0}},
		/* . */ Glyph{Width: 0x8, Height: 0x1, XAdvance: 0x2, XOffset: 0, YOffset: -1, Bitmaps: []uint8{0x80}},
		/* / */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x20, 0x40, 0x80, 0x80}},
		/* 0 */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xa0, 0xa0, 0xa0, 0xc0}},
		/* 1 */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xc0, 0x40, 0x40, 0x40}},
		/* 2 */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x20, 0x40, 0x80, 0xe0}},
		/* 3 */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x20, 0x40, 0x20, 0xc0}},
		/* 4 */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0xe0, 0x20, 0x20}},
		/* 5 */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x80, 0xc0, 0x20, 0xc0}},
		/* 6 */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x80, 0xe0, 0xa0, 0xe0}},
		/* 7 */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x20, 0x40, 0x80, 0x80}},
		/* 8 */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0xa0, 0xe0, 0xa0, 0xe0}},
		/* 9 */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0xa0, 0xe0, 0x20, 0xc0}},
		/* : */ Glyph{Width: 0x8, Height: 0x3, XAdvance: 0x2, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x80, 0x0, 0x80}},
		/* ; */ Glyph{Width: 0x8, Height: 0x4, XAdvance: 0x3, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x40, 0x0, 0x40, 0x80}},
		/* < */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x40, 0x80, 0x40, 0x20}},
		/* = */ Glyph{Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xe0, 0x0, 0xe0}},
		/* > */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x40, 0x20, 0x40, 0x80}},
		/* ? */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x20, 0x40, 0x0, 0x40}},
		/* @ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xa0, 0xe0, 0x80, 0x60}},
		/* A */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xa0, 0xe0, 0xa0, 0xa0}},
		/* B */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xa0, 0xc0, 0xa0, 0xc0}},
		/* C */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x80, 0x80, 0x80, 0x60}},
		/* D */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xa0, 0xa0, 0xa0, 0xc0}},
		/* E */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x80, 0xe0, 0x80, 0xe0}},
		/* F */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x80, 0xe0, 0x80, 0x80}},
		/* G */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x80, 0xe0, 0xa0, 0x60}},
		/* H */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0xe0, 0xa0, 0xa0}},
		/* I */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x40, 0x40, 0x40, 0xe0}},
		/* J */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x20, 0x20, 0xa0, 0x40}},
		/* K */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0xc0, 0xa0, 0xa0}},
		/* L */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80, 0x80, 0x80, 0xe0}},
		/* M */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xe0, 0xe0, 0xa0, 0xa0}},
		/* N */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xe0, 0xe0, 0xe0, 0xa0}},
		/* O */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xa0, 0xa0, 0xa0, 0x40}},
		/* P */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xa0, 0xc0, 0x80, 0x80}},
		/* Q */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xa0, 0xa0, 0xe0, 0x60}},
		/* R */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xa0, 0xe0, 0xc0, 0xa0}},
		/* S */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x80, 0x40, 0x20, 0xc0}},
		/* T */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x40, 0x40, 0x40, 0x40}},
		/* U */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0xa0, 0xa0, 0x60}},
		/* V */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0xa0, 0x40, 0x40}},
		/* W */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0xe0, 0xe0, 0xa0}},
		/* X */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0x40, 0xa0, 0xa0}},
		/* Y */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0x40, 0x40, 0x40}},
		/* Z */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x20, 0x40, 0x80, 0xe0}},
		/* [ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x80, 0x80, 0x80, 0xe0}},
		/* \ */ Glyph{Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x80, 0x40, 0x20}},
		/* ] */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x20, 0x20, 0x20, 0xe0}},
		/* ^ */ Glyph{Width: 0x8, Height: 0x2, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xa0}},
		/* _ */ Glyph{Width: 0x8, Height: 0x1, XAdvance: 0x4, XOffset: 0, YOffset: -1, Bitmaps: []uint8{0xe0}},
		/* ` */ Glyph{Width: 0x8, Height: 0x2, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x40}},
		/* a */ Glyph{Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xc0, 0x60, 0xa0, 0xe0}},
		/* b */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0xc0, 0xa0, 0xa0, 0xc0}},
		/* c */ Glyph{Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0x80, 0x80, 0x60}},
		/* d */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x60, 0xa0, 0xa0, 0x60}},
		/* e */ Glyph{Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0xa0, 0xc0, 0x60}},
		/* f */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x40, 0xe0, 0x40, 0x40}},
		/* g */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0xa0, 0xe0, 0x20, 0x40}},
		/* h */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0xc0, 0xa0, 0xa0, 0xa0}},
		/* i */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x0, 0x80, 0x80, 0x80}},
		/* j */ Glyph{Width: 0x8, Height: 0x6, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x0, 0x20, 0x20, 0xa0, 0x40}},
		/* k */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0xa0, 0xc0, 0xc0, 0xa0}},
		/* l */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x40, 0x40, 0x40, 0xe0}},
		/* m */ Glyph{Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xe0, 0xe0, 0xe0, 0xa0}},
		/* n */ Glyph{Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xc0, 0xa0, 0xa0, 0xa0}},
		/* o */ Glyph{Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x40, 0xa0, 0xa0, 0x40}},
		/* p */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xc0, 0xa0, 0xa0, 0xc0, 0x80}},
		/* q */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0xa0, 0xa0, 0x60, 0x20}},
		/* r */ Glyph{Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0x80, 0x80, 0x80}},
		/* s */ Glyph{Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0xc0, 0x60, 0xc0}},
		/* t */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xe0, 0x40, 0x40, 0x60}},
		/* u */ Glyph{Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xa0, 0xa0, 0xa0, 0x60}},
		/* v */ Glyph{Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xa0, 0xa0, 0xe0, 0x40}},
		/* w */ Glyph{Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xa0, 0xe0, 0xe0, 0xe0}},
		/* x */ Glyph{Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xa0, 0x40, 0x40, 0xa0}},
		/* y */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xa0, 0xa0, 0x60, 0x20, 0x40}},
		/* z */ Glyph{Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xe0, 0x60, 0xc0, 0xe0}},
		/* { */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x40, 0x80, 0x40, 0x60}},
		/* | */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80, 0x0, 0x80, 0x80}},
		/* } */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x40, 0x20, 0x40, 0xc0}},
		/* ~ */ Glyph{Width: 0x8, Height: 0x2, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xc0}},
		/*  */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x0, 0x80, 0x80, 0x80}},
		/*  */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xe0, 0x80, 0xe0, 0x40}},
		/*  */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x40, 0xe0, 0x40, 0xe0}},
		/*  */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x40, 0xe0, 0x40, 0xa0}},
		/*  */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0x40, 0xe0, 0x40}},
		/*  */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80, 0x0, 0x80, 0x80}},
		/*  */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x40, 0xa0, 0x40, 0xc0}},
		/*  */ Glyph{Width: 0x8, Height: 0x1, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0}},
		/*  */ Glyph{Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x80, 0x60}},
		/*  */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xa0, 0xe0, 0x0, 0xe0}},
		/*  */ Glyph{Width: 0x8, Height: 0x3, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0x40}},
		/*  */ Glyph{Width: 0x8, Height: 0x2, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xe0, 0x20}},
		/*  */ Glyph{Width: 0x8, Height: 0x1, XAdvance: 0x3, XOffset: 0, YOffset: -3, Bitmaps: []uint8{0xc0}},
		/*  */ Glyph{Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xc0, 0xa0}},
		/*  */ Glyph{Width: 0x8, Height: 0x1, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0}},
		/*  */ Glyph{Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xa0, 0x40}},
		/*  */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xe0, 0x40, 0x0, 0xe0}},
		/*  */ Glyph{Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x40, 0x60}},
		/*  */ Glyph{Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x60, 0xe0}},
		/*  */ Glyph{Width: 0x8, Height: 0x2, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80}},
		/*  */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0xa0, 0xc0, 0x80}},
		/*  */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xa0, 0x60, 0x60, 0x60}},
		/*  */ Glyph{Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xe0, 0xe0, 0xe0}},
		/*  */ Glyph{Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -3, Bitmaps: []uint8{0x40, 0x20, 0xc0}},
		/*  */ Glyph{Width: 0x8, Height: 0x3, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80, 0x80}},
		/*  */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xa0, 0x40, 0x0, 0xe0}},
		/*  */ Glyph{Width: 0x8, Height: 0x3, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x40, 0x80}},
		/*  */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80, 0x0, 0x60, 0x20}},
		/*  */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80, 0x0, 0xc0, 0x60}},
		/*  */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xc0, 0x0, 0x60, 0x20}},
		/*  */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x0, 0x40, 0x80, 0xe0}},
		/*  */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x20, 0x40, 0xe0, 0xa0}},
		/*  */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0x40, 0xe0, 0xa0}},
		/*   */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0x40, 0xe0, 0xa0}},
		/* ¡ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xc0, 0x40, 0xe0, 0xa0}},
		/* ¢ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x40, 0xa0, 0xe0, 0xa0}},
		/* £ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xc0, 0xa0, 0xe0, 0xa0}},
		/* ¤ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xc0, 0xe0, 0xc0, 0xe0}},
		/* ¥ */ Glyph{Width: 0x8, Height: 0x6, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x80, 0x80, 0x60, 0x20, 0x40}},
		/* ¦ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x20, 0xe0, 0xc0, 0xe0}},
		/* § */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0xe0, 0xc0, 0xe0}},
		/* ¨ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0xe0, 0xc0, 0xe0}},
		/* © */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0xe0, 0xc0, 0xe0}},
		/* ª */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x20, 0xe0, 0x40, 0xe0}},
		/* « */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0xe0, 0x40, 0xe0}},
		/* ¬ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0xe0, 0x40, 0xe0}},
		/* ­ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0xe0, 0x40, 0xe0}},
		/* ® */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xa0, 0xe0, 0xa0, 0xc0}},
		/* ¯ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x60, 0xa0, 0xe0, 0xa0}},
		/* ° */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x20, 0xe0, 0xa0, 0xe0}},
		/* ± */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0xe0, 0xa0, 0xe0}},
		/* ² */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0xe0, 0xa0, 0xe0}},
		/* ³ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x60, 0xe0, 0xa0, 0xe0}},
		/* ´ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0xe0, 0xa0, 0xe0}},
		/* µ */ Glyph{Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xa0, 0x40, 0xa0}},
		/* ¶ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xa0, 0xe0, 0xa0, 0xc0}},
		/* · */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x40, 0xa0, 0xa0, 0xe0}},
		/* ¸ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x40, 0xa0, 0xa0, 0xe0}},
		/* ¹ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0xa0, 0xa0, 0xe0}},
		/* º */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0xa0, 0xa0, 0xe0}},
		/* » */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x40, 0xa0, 0xe0, 0x40}},
		/* ¼ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0xe0, 0xa0, 0xe0, 0x80}},
		/* ½ */ Glyph{Width: 0x8, Height: 0x6, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xa0, 0xc0, 0xa0, 0xc0, 0x80}},
		/* ¾ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x20, 0x60, 0xa0, 0xe0}},
		/* ¿ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0x60, 0xa0, 0xe0}},
		/* À */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0x60, 0xa0, 0xe0}},
		/* Á */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xc0, 0x60, 0xa0, 0xe0}},
		/* Â */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0x60, 0xa0, 0xe0}},
		/* Ã */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x60, 0x60, 0xa0, 0xe0}},
		/* Ä */ Glyph{Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0xe0, 0xe0, 0xc0}},
		/* Å */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0x80, 0x60, 0x20, 0x40}},
		/* Æ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x20, 0x60, 0xe0, 0x60}},
		/* Ç */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0x60, 0xe0, 0x60}},
		/* È */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0x60, 0xe0, 0x60}},
		/* É */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0x60, 0xe0, 0x60}},
		/* Ê */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x40, 0x80, 0x80, 0x80}},
		/* Ë */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0x40, 0x40, 0x40}},
		/* Ì */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0x40, 0x40, 0x40}},
		/* Í */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0x40, 0x40, 0x40}},
		/* Î */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xc0, 0x60, 0xa0, 0x60}},
		/* Ï */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x60, 0xc0, 0xa0, 0xa0}},
		/* Ð */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x20, 0x40, 0xa0, 0x40}},
		/* Ñ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0x40, 0xa0, 0x40}},
		/* Ò */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0x40, 0xa0, 0x40}},
		/* Ó */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x60, 0x40, 0xa0, 0x40}},
		/* Ô */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0x40, 0xa0, 0x40}},
		/* Õ */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x0, 0xe0, 0x0, 0x40}},
		/* Ö */ Glyph{Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0xe0, 0xa0, 0xc0}},
		/* × */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x40, 0xa0, 0xa0, 0x60}},
		/* Ø */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x40, 0xa0, 0xa0, 0x60}},
		/* Ù */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0xa0, 0xa0, 0x60}},
		/* Ú */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0xa0, 0xa0, 0x60}},
		/* Û */ Glyph{Width: 0x8, Height: 0x6, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x40, 0xa0, 0x60, 0x20, 0x40}},
		/* Ü */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x80, 0xc0, 0xa0, 0xc0, 0x80}},
		/* Ý */ Glyph{Width: 0x8, Height: 0x6, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0xa0, 0x60, 0x20, 0x40}},
		/* Þ */ Glyph{Width: 0x8, Height: 0x1, XAdvance: 0x2, XOffset: 0, YOffset: -1, Bitmaps: []uint8{0x0}},
		/* ß */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xc0, 0xe0, 0xc0, 0x60}},
		/* à */ Glyph{Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0xe0, 0xc0, 0xe0}},
		/* á */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x60, 0xc0, 0x60, 0xc0}},
		/* â */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x60, 0xc0, 0x60, 0xc0}},
		/* ã */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0xa0, 0x40, 0x40}},
		/* ä */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xe0, 0x60, 0xc0, 0xe0}},
		/* å */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xe0, 0x60, 0xc0, 0xe0}},
		/* æ */ Glyph{Width: 0x8, Height: 0x1, XAdvance: 0x2, XOffset: 0, YOffset: -1, Bitmaps: []uint8{0x0}},
		/* ç */ Glyph{Width: 0x8, Height: 0x1, XAdvance: 0x2, XOffset: 0, YOffset: -1, Bitmaps: []uint8{0x0}},
		/* è */ Glyph{Width: 0x8, Height: 0x1, XAdvance: 0x2, XOffset: 0, YOffset: -3, Bitmaps: []uint8{0x80}},
		/* é */ Glyph{Width: 0x8, Height: 0x1, XAdvance: 0x4, XOffset: 0, YOffset: -1, Bitmaps: []uint8{0xa0}},
		/* ê */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xe0, 0xe0, 0xc0, 0x60}},
		/* ë */ Glyph{Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0xa0, 0xa0, 0xa0, 0xe0}},
	},

	RuneToIndex: []RuneToIndex{
		/*   */ RuneToIndex{Rune: 32, Index: 0x0},
		/* ! */ RuneToIndex{Rune: 33, Index: 0x1},
		/* " */ RuneToIndex{Rune: 34, Index: 0x2},
		/* # */ RuneToIndex{Rune: 35, Index: 0x3},
		/* $ */ RuneToIndex{Rune: 36, Index: 0x4},
		/* % */ RuneToIndex{Rune: 37, Index: 0x5},
		/* & */ RuneToIndex{Rune: 38, Index: 0x6},
		/* ' */ RuneToIndex{Rune: 39, Index: 0x7},
		/* ( */ RuneToIndex{Rune: 40, Index: 0x8},
		/* ) */ RuneToIndex{Rune: 41, Index: 0x9},
		/* * */ RuneToIndex{Rune: 42, Index: 0xa},
		/* + */ RuneToIndex{Rune: 43, Index: 0xb},
		/* , */ RuneToIndex{Rune: 44, Index: 0xc},
		/* - */ RuneToIndex{Rune: 45, Index: 0xd},
		/* . */ RuneToIndex{Rune: 46, Index: 0xe},
		/* / */ RuneToIndex{Rune: 47, Index: 0xf},
		/* 0 */ RuneToIndex{Rune: 48, Index: 0x10},
		/* 1 */ RuneToIndex{Rune: 49, Index: 0x11},
		/* 2 */ RuneToIndex{Rune: 50, Index: 0x12},
		/* 3 */ RuneToIndex{Rune: 51, Index: 0x13},
		/* 4 */ RuneToIndex{Rune: 52, Index: 0x14},
		/* 5 */ RuneToIndex{Rune: 53, Index: 0x15},
		/* 6 */ RuneToIndex{Rune: 54, Index: 0x16},
		/* 7 */ RuneToIndex{Rune: 55, Index: 0x17},
		/* 8 */ RuneToIndex{Rune: 56, Index: 0x18},
		/* 9 */ RuneToIndex{Rune: 57, Index: 0x19},
		/* : */ RuneToIndex{Rune: 58, Index: 0x1a},
		/* ; */ RuneToIndex{Rune: 59, Index: 0x1b},
		/* < */ RuneToIndex{Rune: 60, Index: 0x1c},
		/* = */ RuneToIndex{Rune: 61, Index: 0x1d},
		/* > */ RuneToIndex{Rune: 62, Index: 0x1e},
		/* ? */ RuneToIndex{Rune: 63, Index: 0x1f},
		/* @ */ RuneToIndex{Rune: 64, Index: 0x20},
		/* A */ RuneToIndex{Rune: 65, Index: 0x21},
		/* B */ RuneToIndex{Rune: 66, Index: 0x22},
		/* C */ RuneToIndex{Rune: 67, Index: 0x23},
		/* D */ RuneToIndex{Rune: 68, Index: 0x24},
		/* E */ RuneToIndex{Rune: 69, Index: 0x25},
		/* F */ RuneToIndex{Rune: 70, Index: 0x26},
		/* G */ RuneToIndex{Rune: 71, Index: 0x27},
		/* H */ RuneToIndex{Rune: 72, Index: 0x28},
		/* I */ RuneToIndex{Rune: 73, Index: 0x29},
		/* J */ RuneToIndex{Rune: 74, Index: 0x2a},
		/* K */ RuneToIndex{Rune: 75, Index: 0x2b},
		/* L */ RuneToIndex{Rune: 76, Index: 0x2c},
		/* M */ RuneToIndex{Rune: 77, Index: 0x2d},
		/* N */ RuneToIndex{Rune: 78, Index: 0x2e},
		/* O */ RuneToIndex{Rune: 79, Index: 0x2f},
		/* P */ RuneToIndex{Rune: 80, Index: 0x30},
		/* Q */ RuneToIndex{Rune: 81, Index: 0x31},
		/* R */ RuneToIndex{Rune: 82, Index: 0x32},
		/* S */ RuneToIndex{Rune: 83, Index: 0x33},
		/* T */ RuneToIndex{Rune: 84, Index: 0x34},
		/* U */ RuneToIndex{Rune: 85, Index: 0x35},
		/* V */ RuneToIndex{Rune: 86, Index: 0x36},
		/* W */ RuneToIndex{Rune: 87, Index: 0x37},
		/* X */ RuneToIndex{Rune: 88, Index: 0x38},
		/* Y */ RuneToIndex{Rune: 89, Index: 0x39},
		/* Z */ RuneToIndex{Rune: 90, Index: 0x3a},
		/* [ */ RuneToIndex{Rune: 91, Index: 0x3b},
		/* \ */ RuneToIndex{Rune: 92, Index: 0x3c},
		/* ] */ RuneToIndex{Rune: 93, Index: 0x3d},
		/* ^ */ RuneToIndex{Rune: 94, Index: 0x3e},
		/* _ */ RuneToIndex{Rune: 95, Index: 0x3f},
		/* ` */ RuneToIndex{Rune: 96, Index: 0x40},
		/* a */ RuneToIndex{Rune: 97, Index: 0x41},
		/* b */ RuneToIndex{Rune: 98, Index: 0x42},
		/* c */ RuneToIndex{Rune: 99, Index: 0x43},
		/* d */ RuneToIndex{Rune: 100, Index: 0x44},
		/* e */ RuneToIndex{Rune: 101, Index: 0x45},
		/* f */ RuneToIndex{Rune: 102, Index: 0x46},
		/* g */ RuneToIndex{Rune: 103, Index: 0x47},
		/* h */ RuneToIndex{Rune: 104, Index: 0x48},
		/* i */ RuneToIndex{Rune: 105, Index: 0x49},
		/* j */ RuneToIndex{Rune: 106, Index: 0x4a},
		/* k */ RuneToIndex{Rune: 107, Index: 0x4b},
		/* l */ RuneToIndex{Rune: 108, Index: 0x4c},
		/* m */ RuneToIndex{Rune: 109, Index: 0x4d},
		/* n */ RuneToIndex{Rune: 110, Index: 0x4e},
		/* o */ RuneToIndex{Rune: 111, Index: 0x4f},
		/* p */ RuneToIndex{Rune: 112, Index: 0x50},
		/* q */ RuneToIndex{Rune: 113, Index: 0x51},
		/* r */ RuneToIndex{Rune: 114, Index: 0x52},
		/* s */ RuneToIndex{Rune: 115, Index: 0x53},
		/* t */ RuneToIndex{Rune: 116, Index: 0x54},
		/* u */ RuneToIndex{Rune: 117, Index: 0x55},
		/* v */ RuneToIndex{Rune: 118, Index: 0x56},
		/* w */ RuneToIndex{Rune: 119, Index: 0x57},
		/* x */ RuneToIndex{Rune: 120, Index: 0x58},
		/* y */ RuneToIndex{Rune: 121, Index: 0x59},
		/* z */ RuneToIndex{Rune: 122, Index: 0x5a},
		/* { */ RuneToIndex{Rune: 123, Index: 0x5b},
		/* | */ RuneToIndex{Rune: 124, Index: 0x5c},
		/* } */ RuneToIndex{Rune: 125, Index: 0x5d},
		/* ~ */ RuneToIndex{Rune: 126, Index: 0x5e},
		/*  */ RuneToIndex{Rune: 127, Index: 0x5f},
		/*  */ RuneToIndex{Rune: 128, Index: 0x60},
		/*  */ RuneToIndex{Rune: 129, Index: 0x61},
		/*  */ RuneToIndex{Rune: 130, Index: 0x62},
		/*  */ RuneToIndex{Rune: 131, Index: 0x63},
		/*  */ RuneToIndex{Rune: 132, Index: 0x64},
		/*  */ RuneToIndex{Rune: 133, Index: 0x65},
		/*  */ RuneToIndex{Rune: 134, Index: 0x66},
		/*  */ RuneToIndex{Rune: 135, Index: 0x67},
		/*  */ RuneToIndex{Rune: 136, Index: 0x68},
		/*  */ RuneToIndex{Rune: 137, Index: 0x69},
		/*  */ RuneToIndex{Rune: 138, Index: 0x6a},
		/*  */ RuneToIndex{Rune: 139, Index: 0x6b},
		/*  */ RuneToIndex{Rune: 140, Index: 0x6c},
		/*  */ RuneToIndex{Rune: 141, Index: 0x6d},
		/*  */ RuneToIndex{Rune: 142, Index: 0x6e},
		/*  */ RuneToIndex{Rune: 143, Index: 0x6f},
		/*  */ RuneToIndex{Rune: 144, Index: 0x70},
		/*  */ RuneToIndex{Rune: 145, Index: 0x71},
		/*  */ RuneToIndex{Rune: 146, Index: 0x72},
		/*  */ RuneToIndex{Rune: 147, Index: 0x73},
		/*  */ RuneToIndex{Rune: 148, Index: 0x74},
		/*  */ RuneToIndex{Rune: 149, Index: 0x75},
		/*  */ RuneToIndex{Rune: 150, Index: 0x76},
		/*  */ RuneToIndex{Rune: 151, Index: 0x77},
		/*  */ RuneToIndex{Rune: 152, Index: 0x78},
		/*  */ RuneToIndex{Rune: 153, Index: 0x79},
		/*  */ RuneToIndex{Rune: 154, Index: 0x7a},
		/*  */ RuneToIndex{Rune: 155, Index: 0x7b},
		/*  */ RuneToIndex{Rune: 156, Index: 0x7c},
		/*  */ RuneToIndex{Rune: 157, Index: 0x7d},
		/*  */ RuneToIndex{Rune: 158, Index: 0x7e},
		/*  */ RuneToIndex{Rune: 159, Index: 0x7f},
		/*   */ RuneToIndex{Rune: 160, Index: 0x80},
		/* ¡ */ RuneToIndex{Rune: 161, Index: 0x81},
		/* ¢ */ RuneToIndex{Rune: 162, Index: 0x82},
		/* £ */ RuneToIndex{Rune: 163, Index: 0x83},
		/* ¤ */ RuneToIndex{Rune: 164, Index: 0x84},
		/* ¥ */ RuneToIndex{Rune: 165, Index: 0x85},
		/* ¦ */ RuneToIndex{Rune: 166, Index: 0x86},
		/* § */ RuneToIndex{Rune: 167, Index: 0x87},
		/* ¨ */ RuneToIndex{Rune: 168, Index: 0x88},
		/* © */ RuneToIndex{Rune: 169, Index: 0x89},
		/* ª */ RuneToIndex{Rune: 170, Index: 0x8a},
		/* « */ RuneToIndex{Rune: 171, Index: 0x8b},
		/* ¬ */ RuneToIndex{Rune: 172, Index: 0x8c},
		/* ­ */ RuneToIndex{Rune: 173, Index: 0x8d},
		/* ® */ RuneToIndex{Rune: 174, Index: 0x8e},
		/* ¯ */ RuneToIndex{Rune: 175, Index: 0x8f},
		/* ° */ RuneToIndex{Rune: 176, Index: 0x90},
		/* ± */ RuneToIndex{Rune: 177, Index: 0x91},
		/* ² */ RuneToIndex{Rune: 178, Index: 0x92},
		/* ³ */ RuneToIndex{Rune: 179, Index: 0x93},
		/* ´ */ RuneToIndex{Rune: 180, Index: 0x94},
		/* µ */ RuneToIndex{Rune: 181, Index: 0x95},
		/* ¶ */ RuneToIndex{Rune: 182, Index: 0x96},
		/* · */ RuneToIndex{Rune: 183, Index: 0x97},
		/* ¸ */ RuneToIndex{Rune: 184, Index: 0x98},
		/* ¹ */ RuneToIndex{Rune: 185, Index: 0x99},
		/* º */ RuneToIndex{Rune: 186, Index: 0x9a},
		/* » */ RuneToIndex{Rune: 187, Index: 0x9b},
		/* ¼ */ RuneToIndex{Rune: 188, Index: 0x9c},
		/* ½ */ RuneToIndex{Rune: 189, Index: 0x9d},
		/* ¾ */ RuneToIndex{Rune: 190, Index: 0x9e},
		/* ¿ */ RuneToIndex{Rune: 191, Index: 0x9f},
		/* À */ RuneToIndex{Rune: 192, Index: 0xa0},
		/* Á */ RuneToIndex{Rune: 193, Index: 0xa1},
		/* Â */ RuneToIndex{Rune: 194, Index: 0xa2},
		/* Ã */ RuneToIndex{Rune: 195, Index: 0xa3},
		/* Ä */ RuneToIndex{Rune: 196, Index: 0xa4},
		/* Å */ RuneToIndex{Rune: 197, Index: 0xa5},
		/* Æ */ RuneToIndex{Rune: 198, Index: 0xa6},
		/* Ç */ RuneToIndex{Rune: 199, Index: 0xa7},
		/* È */ RuneToIndex{Rune: 200, Index: 0xa8},
		/* É */ RuneToIndex{Rune: 201, Index: 0xa9},
		/* Ê */ RuneToIndex{Rune: 202, Index: 0xaa},
		/* Ë */ RuneToIndex{Rune: 203, Index: 0xab},
		/* Ì */ RuneToIndex{Rune: 204, Index: 0xac},
		/* Í */ RuneToIndex{Rune: 205, Index: 0xad},
		/* Î */ RuneToIndex{Rune: 206, Index: 0xae},
		/* Ï */ RuneToIndex{Rune: 207, Index: 0xaf},
		/* Ð */ RuneToIndex{Rune: 208, Index: 0xb0},
		/* Ñ */ RuneToIndex{Rune: 209, Index: 0xb1},
		/* Ò */ RuneToIndex{Rune: 210, Index: 0xb2},
		/* Ó */ RuneToIndex{Rune: 211, Index: 0xb3},
		/* Ô */ RuneToIndex{Rune: 212, Index: 0xb4},
		/* Õ */ RuneToIndex{Rune: 213, Index: 0xb5},
		/* Ö */ RuneToIndex{Rune: 214, Index: 0xb6},
		/* × */ RuneToIndex{Rune: 215, Index: 0xb7},
		/* Ø */ RuneToIndex{Rune: 216, Index: 0xb8},
		/* Ù */ RuneToIndex{Rune: 217, Index: 0xb9},
		/* Ú */ RuneToIndex{Rune: 218, Index: 0xba},
		/* Û */ RuneToIndex{Rune: 219, Index: 0xbb},
		/* Ü */ RuneToIndex{Rune: 220, Index: 0xbc},
		/* Ý */ RuneToIndex{Rune: 221, Index: 0xbd},
		/* Þ */ RuneToIndex{Rune: 222, Index: 0xbe},
		/* ß */ RuneToIndex{Rune: 223, Index: 0xbf},
		/* à */ RuneToIndex{Rune: 224, Index: 0xc0},
		/* á */ RuneToIndex{Rune: 225, Index: 0xc1},
		/* â */ RuneToIndex{Rune: 226, Index: 0xc2},
		/* ã */ RuneToIndex{Rune: 227, Index: 0xc3},
		/* ä */ RuneToIndex{Rune: 228, Index: 0xc4},
		/* å */ RuneToIndex{Rune: 229, Index: 0xc5},
		/* æ */ RuneToIndex{Rune: 230, Index: 0xc6},
		/* ç */ RuneToIndex{Rune: 231, Index: 0xc7},
		/* è */ RuneToIndex{Rune: 232, Index: 0xc8},
		/* é */ RuneToIndex{Rune: 233, Index: 0xc9},
		/* ê */ RuneToIndex{Rune: 234, Index: 0xca},
		/* ë */ RuneToIndex{Rune: 235, Index: 0xcb},
	},

	YAdvance: 0x6,
}
