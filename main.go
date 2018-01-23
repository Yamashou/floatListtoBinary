package floatListtoBinary

import (
	"encoding/binary"
	"math"
)

var sizeFloat64 = binary.Size(float64(0))

func MarshalBinary(data []float64) []byte {
	p := 0
	dataLen := len(data)
	bufLen := int64(dataLen) * int64(sizeFloat64)
	buf := make([]byte, bufLen)

	for i := 0; i < dataLen; i++ {
		binary.LittleEndian.PutUint64(buf[p:p+sizeFloat64], math.Float64bits(data[i]))
		p += sizeFloat64
	}
	return buf
}

func UnmarshalBinary(buf []byte) []float64 {
	p := 0
	dataLen := len(buf) / sizeFloat64
	data := make([]float64, dataLen)
	for i := range data {
		data[i] = math.Float64frombits(binary.LittleEndian.Uint64(buf[p : p+sizeFloat64]))
		p += sizeFloat64
	}

	return data
}
