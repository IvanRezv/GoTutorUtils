package main

func main() {
	var a int8 = 127                  // 1 byte
	var b int16 = 21321               // 2 byte
	var c int32 = 303123123           // 4 byte
	var d int64 = 3012301231232131231 // 8 byte

	var e uint8  // 0 - 255 1 byte
	var f uint16 // 0 - 65545 2 byte
	var g uint32 // 0 - 4bil 4 bytes
	var h uint64 // 0 - 18pent

	var i byte // ==uint8
	var j rune // == int32
	var k int  // 4 or 8 bytesexwithpenisok
	var m uint // bez znakovoe

	var a1 float32 = 1.8         // 1.4*10^45 - 3.4*10^38
	var a2 float64 = 11241.12412 // 4.8*10^302 - 1.8*10^305

	var c1 complex64 = 1 + 2i
	var c2 complex128 = 1 + 5i

	var name string = "Test"
}
