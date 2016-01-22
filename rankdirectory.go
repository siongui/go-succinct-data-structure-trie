package bits

import "math"

/**
  Fixed values for the L1 and L2 table sizes in the Rank Directory
*/
var L1 uint = 32 * 32
var L2 uint = 32

/**
  The rank directory allows you to build an index to quickly compute the
  rank() and select() functions. The index can itself be encoded as a binary
  string.
*/
type RankDirectory struct {
	directory   BitString
	data        BitString // data of succinct trie
	l1Size      uint
	l2Size      uint
	l1Bits      uint
	l2Bits      uint
	sectionBits uint
	numBits     uint
}

/**
  Used to build a rank directory from the given input string.

  @param data A javascript string containing the data, as readable using the
  BitString object.

  @param numBits The number of bits to index.

  @param l1Size The number of bits that each entry in the Level 1 table
  summarizes. This should be a multiple of l2Size.

  @param l2Size The number of bits that each entry in the Level 2 table
  summarizes.
*/
func CreateRankDirectory(data string, numBits, l1Size, l2Size uint) RankDirectory {
	bits := BitString{}
	bits.Init(data)
	var p, i uint = 0, 0
	var count1, count2 uint = 0, 0
	l1bits := uint(math.Ceil(math.Log2(float64(numBits))))
	l2bits := uint(math.Ceil(math.Log2(float64(l1Size))))

	directory := BitWriter{}

	for p+l2Size <= numBits {
		count2 += bits.Count(p, l2Size)
		i += l2Size
		p += l2Size
		if i == l1Size {
			count1 += count2
			directory.Write(count1, l1bits)
			count2 = 0
			i = 0
		} else {
			directory.Write(count2, l2bits)
		}
	}

	rd := RankDirectory{}
	rd.Init(directory.GetData(), data, numBits, l1Size, l2Size)
	return rd
}

func (rd *RankDirectory) Init(directoryData, bitData string, numBits, l1Size, l2Size uint) {
	rd.directory.Init(directoryData)
	rd.data.Init(bitData)
	rd.l1Size = l1Size
	rd.l2Size = l2Size
	rd.l1Bits = uint(math.Ceil(math.Log2(float64(numBits))))
	rd.l2Bits = uint(math.Ceil(math.Log2(float64(l1Size))))
	rd.sectionBits = (l1Size/l2Size-1)*rd.l2Bits + rd.l1Bits
	rd.numBits = numBits
}

/**
  Returns the string representation of the directory.
*/
func (rd *RankDirectory) GetData() string {
	return rd.directory.GetData()
}

/**
  Returns the number of 1 or 0 bits (depending on the "which" parameter) to
  to and including position x.
*/
func (rd *RankDirectory) Rank(which, x uint) uint {

	if which == 0 {
		return x - rd.Rank(1, x) + 1
	}

	var rank uint = 0
	o := x
	var sectionPos uint = 0

	if o >= rd.l1Size {
		sectionPos = (o/rd.l1Size | 0) * rd.sectionBits
		rank = rd.directory.Get(sectionPos-rd.l1Bits, rd.l1Bits)
		o = o % rd.l1Size
	}

	if o >= rd.l2Size {
		sectionPos += (o/rd.l2Size | 0) * rd.l2Bits
		rank += rd.directory.Get(sectionPos-rd.l2Bits, rd.l2Bits)
	}

	rank += rd.data.Count(x-x%rd.l2Size, x%rd.l2Size+1)

	return rank
}

/**
  Returns the position of the y'th 0 or 1 bit, depending on the "which"
  parameter.
*/
func (rd *RankDirectory) Select(which, y uint) uint {
	high := int(rd.numBits)
	low := -1
	val := -1

	for high-low > 1 {
		probe := (high+low)/2 | 0
		r := rd.Rank(which, uint(probe))

		if r == y {
			// We have to continue searching after we have found it,
			// because we want the _first_ occurrence.
			val = probe
			high = probe
		} else if r < y {
			low = probe
		} else {
			high = probe
		}
	}

	return uint(val)
}
