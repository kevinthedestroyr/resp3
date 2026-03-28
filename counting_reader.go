package resp3

import (
	"bufio"
	"io"
)

type CountingReader struct {
	reader *bufio.Reader
	count  int64
}

var _ io.Reader = (*CountingReader)(nil)

func (r *CountingReader) Count() int64 {
	return r.count
}

func (r *CountingReader) ResetCount() {
	r.count = 0
}

func NewCountingReader(reader *bufio.Reader) *CountingReader {
	return &CountingReader{
		reader: reader,
	}
}

func (r *CountingReader) Read(p []byte) (n int, err error) {
	n, err = r.reader.Read(p)
	r.count += int64(n)
	return n, err
}

func (r *CountingReader) ReadByte() (byte, error) {
	b, err := r.reader.ReadByte()
	if err != nil {
		return 0, err
	}
	r.count++
	return b, nil
}

func (r *CountingReader) ReadBytes(delim byte) ([]byte, error) {
	line, err := r.reader.ReadBytes(delim)
	if err != nil {
		return nil, err
	}
	r.count += int64(len(line))
	return line, nil
}

func (r *CountingReader) ReadString(delim byte) (string, error) {
	line, err := r.reader.ReadString(delim)
	if err != nil {
		return "", err
	}
	r.count += int64(len(line))
	return line, nil
}

func (r *CountingReader) ReadLine() ([]byte, bool, error) {
	line, isPrefix, err := r.reader.ReadLine()
	if err != nil {
		return nil, false, err
	}
	r.count += int64(len(line))
	return line, isPrefix, nil
}

func (r *CountingReader) ReadSlice(delim byte) ([]byte, error) {
	line, err := r.reader.ReadSlice(delim)
	if err != nil {
		return nil, err
	}
	r.count += int64(len(line))
	return line, nil
}

func (r *CountingReader) ReadRune() (rune, int, error) {
	rune, size, err := r.reader.ReadRune()
	if err != nil {
		return 0, 0, err
	}
	r.count += int64(size)
	return rune, size, nil
}

func (r *CountingReader) UnreadByte() error {
	return r.reader.UnreadByte()
}

func (r *CountingReader) UnreadRune() error {
	return r.reader.UnreadRune()
}

func (r *CountingReader) Peek(n int) ([]byte, error) {
	return r.reader.Peek(n)
}

func (r *CountingReader) Discard(n int) (discarded int, err error) {
	discarded, err = r.reader.Discard(n)
	r.count += int64(discarded)
	return discarded, err
}
