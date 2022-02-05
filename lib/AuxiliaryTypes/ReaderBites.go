package AuxiliaryTypes

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)
//type bufSlise []byte
//
//func (b bufSlise)convertByteToStringBite()string{
//	var buf strings.Builder
//	for _, ch := range b{
//		buf.WriteString(string(ch))
//	}
//	return buf.String()
//}

const (
	countReadByteIsOneIter = 100
	Bite0 = Bite(48)
	Bite1 = Bite(49)
)


type Bite uint8

func (b Bite) ToString() string {
	return string(b)
}


type bitesBuffer struct {
	*bytes.Buffer
}

func NewBitesBufferFromByte( b []byte) *bitesBuffer{
	return &bitesBuffer{bytes.NewBuffer(b)}
}

func NewBitesBuffer(bs []Bite)*bitesBuffer{
	res := make([]byte, len(bs))
	for i, b := range bs {
		res[i] = byte(b)
	}
	return &bitesBuffer{bytes.NewBuffer(res)}
}

func (bB *bitesBuffer) WriteBitesFromBytes(bt []byte) (int, error) {
	count, n := 0, 0
	var err  error = nil
	for _ , b := range bt {
		n, err = bB.WriteString(fmt.Sprintf("%08b", b))
		count += n
	}

	return count, err
}

func (bB bitesBuffer) NextBites(n int) []Bite {
	if bB.Len() == 0{
		return nil
	}
	next := bB.Next(n)
	res := make([]Bite, len(next))
	for idx :=0 ; idx < len(next); idx++{
		res[idx] = Bite(next[idx])
	}
	return res
}

func (bB bitesBuffer) NextBitesTOString(n int) string {
	next := bB.Next(n)
	return string(next)
}



type ReaderBites struct {
	buf *bitesBuffer
	readerSlice []byte
	reader *bufio.Reader

}

func NewReaderBites(r *bufio.Reader) *ReaderBites {
	return &ReaderBites{
		NewBitesBufferFromByte(make([]byte, 0)),
		make([]byte, countReadByteIsOneIter),
		r}
}

func (rtb ReaderBites) Read(n int) ([]Bite, error){
	if rtb.buf.Len() == 0 || rtb.buf.Len() < n {
		for {
			_, err := rtb.topUpBuf()
			if err == io.EOF {
				return nil, err
			}
			if err == io.ErrUnexpectedEOF && rtb.buf.Len() < n{
				return rtb.buf.NextBites(rtb.buf.Len()), err
			}
			if rtb.buf.Len() > n{
				break
			}

		}
	}


	return rtb.buf.NextBites(n), nil
}

func (rtb ReaderBites) ReadToString(n int) (string, error) {
	if rtb.buf.Len() == 0 || rtb.buf.Len() < n {
		for {
			_, err := rtb.topUpBuf()
			if err == io.EOF {
				return "", err
			}
			if err == io.ErrUnexpectedEOF && rtb.buf.Len() < n{
				return rtb.buf.NextBitesTOString(rtb.buf.Len()), err
			}
			if rtb.buf.Len() > n{
				break
			}

		}
	}
	return rtb.buf.NextBitesTOString(n), nil
}

func (rtb ReaderBites) topUpBuf() (int,error) {
		n, err := io.ReadFull(rtb.reader, rtb.readerSlice)
		if err == io.ErrUnexpectedEOF {
			rtb.buf.WriteBitesFromBytes(rtb.readerSlice[:n-1])
			return n, err
		}
		if err == io.EOF {
			return  0, err
		}
		return rtb.buf.WriteBitesFromBytes(rtb.readerSlice)
}