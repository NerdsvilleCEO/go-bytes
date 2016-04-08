package main

type ByteReader struct {
	bytes    []byte
	offset int
	limit  int
}

func (s *StringReader) Len() int {
	// Orig length - parsed
	return len(s.str) - (s.offset)
}

func NewByteReader(str string) *ByteReader {
	return &ByteReader{
		str: []byte(str),
	}
}

// Take in a map of bytes and return n number of bytes parsed with err
func (s *ByteReader) Read(p []byte) (n int, err error) {
	// If len p == 0, error
	if len(p) == 0 {
		err = errors.New("Must initialize byte slice with len > 0")
	}
	// while len(p) > 0,
	// if n < lim and unparsed len byte slice > 0
	// or if s.limit is -1 and len unparsed > 0
	// parse one byte from s.str,
	// increment n and offset
	// else return
	for len(p) > 0 {
		if n < s.limit && s.Len() > 0 || s.limit == -1 && s.Len() > 0 {
			p = append(p, s.str[s.offset])
			n += 1
			s.offset += 1
		}
	}
	return
}
