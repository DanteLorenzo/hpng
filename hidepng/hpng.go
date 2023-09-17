package hpng

/////////////////////////////////////////////
//                  PNG                  ///
///////////////////////////////////////////
// 	IHDR  //  IDAT  //  dxXm  //  IEND  //
/////////////////////////////////////////

// dxXm - custom chunk, where stored secret file

//////////////////////////////////////////////
//                  dxXm                   //
////////////////////////////////////////////
// 	Length  //  Type  //  Data  //  CRC  //
//////////////////////////////////////////

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
	"log"
	"os"
)

const (
	IEND_CHUNK_LENGTH_OFFSET = -12
	dxXm_CHUNK_CRC			 = -4
)

var (
	CHUNK_TYPE = []byte{0x64, 0x78, 0x58, 0x6D} // "dxXm"
)

//CreateHPNG - create new PNG file with hidden file
func CreateHPNG( mainFileName string, filetoHideName string) {
	
	mainFile := readFile(mainFileName) // mainFile - PNG file, where we hide filetoHide
	filetoHide := readFile(filetoHideName) // filetoHide - file, which we hide in mainFile

	fullChunk := MakeHPNG(mainFile, filetoHide) // fullChunks - all chunks of new PNG file

	// Create new PNG file
	newPng, err := os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}
	// Close new PNG file
	defer newPng.Close()

	// Write new PNG file
	_, err = newPng.Write(fullChunk)
	if err != nil {
		log.Fatal(err)
	}
	
}

//ReverseHPNG - retrieve hide file
func ReverseHPNG( sourceFile string, outputFile string ) {
	
	srcFile := readFile(sourceFile) // srcFile - PNG file, where stored hidden file
	
	hideChunk := TakeFromHPNG(srcFile) // Take hidden chunk

	// Create new file
	newFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Close new PNG file
	defer newFile.Close()

	//Write new PNG file
	_, err = newFile.Write(hideChunk)
	if err != nil {
		log.Fatal(err)
	}
}

//MakeHPNG - compose chunks to one chunk
func MakeHPNG( mainFile []byte, filetoHide []byte) []byte {
	return (append(append(iDAT(mainFile), dxXm(filetoHide)...), iEND(mainFile)...))
}

//TakeHPNG - return hidden chunk
func TakeFromHPNG( srcFile []byte) []byte {

	// Get second part of source file
	_, after, found := bytes.Cut(srcFile, CHUNK_TYPE)
	if found{
		return after[:len(after) + IEND_CHUNK_LENGTH_OFFSET + dxXm_CHUNK_CRC] 
	} else {
		log.Fatal("No hidden file in this PNG")
		return []byte{}
	}
	
}

// readFile - read file and return slice of bytes
func readFile( fileName string) []byte {
	src, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return src
}

// iDAT - return all chunks before IEND chunk(IHDR, IDAT, e.t.c.)
func iDAT( img []byte) []byte {
	iendChunkStart := len(img) + IEND_CHUNK_LENGTH_OFFSET
	return img[:iendChunkStart]
}

// iEND - return IEND chunk
func iEND( mainFile []byte) []byte {
	iendChunkStart := len(mainFile) + IEND_CHUNK_LENGTH_OFFSET
	return mainFile[iendChunkStart:]
}

// dxXm - return custom chunk
func dxXm( filetoHide []byte) []byte {
	threeChunks := append(dxXmLength(filetoHide), append(dxXmType(filetoHide), filetoHide...)...)
	crcChunk := dxXmCRC(threeChunks)
	return append(threeChunks, crcChunk...)
}

// dxXmLength - return length chunk of filetoHide
func dxXmLength( filetoHide []byte) []byte {
	lChunk := make([]byte, 4)
	binary.BigEndian.PutUint32(lChunk, uint32(len(filetoHide)))
	return lChunk
}

// dxXmType - return type chunk of filetoHide
func dxXmType( filetoHide []byte) []byte {
	return CHUNK_TYPE
}

// dxXmCRC - return CRC chunks type, length and data
func dxXmCRC( chunk []byte) []byte {
	crc := crc32.ChecksumIEEE(chunk)
	crcChunk := make([]byte, 4)
	binary.BigEndian.PutUint32(crcChunk, crc)
	return crcChunk
}
