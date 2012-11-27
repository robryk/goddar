package chunkedio

import "io"

type ChunkedReader func(offset int64) (startOffset int64, contents io.ReaderAt, err error)

func (cr ChunkedReader) ReadAt(p []byte, off int64) (n int, err error) {
	n = 0
	for len(p) > 0 {
		var startOff int64
		var r io.ReaderAt

		startOff, r, err = cr(off)
		if err != nil {
			return
		}

		var m int
		m, err = r.ReadAt(p, off - startOff)

		n += m
		off += int64(m)
		p = p[m:]

		if err == nil {
			if len(p) > 0 {
				panic("ReaderAt returned a non-full read without errors");
			}
			return
		}

		if err == io.EOF {
			err = nil
		}
		if err != nil {
			return
		}
	}
	return
}
