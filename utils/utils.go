package utils

import (
	"bytes"
// 	"compress/lzw"  
	"github.com/pierrec/lz4/v4"  
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"time"
)

//还原数据
func ReCopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}
    //打开源文件  
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()
    
	// 创建一个 LZ4 读取器 
	reader  := lz4.NewReader(source)  
  
    // 将压缩后的数据写入到目标文件  
	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	
    // 使用 io.Copy 将解压缩的数据写入到文件中  
	nBytes, err := io.Copy(destination, reader)
	return nBytes, err
}

//压缩数据  lz4
func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}
    //打开源文件  
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()
    
    
	// 创建一个缓冲区来存储压缩后的数据  
	var compressedData bytes.Buffer  
	writer := lz4.NewWriter(&compressedData)  
  
	// 使用io.Copy来读取源文件并写入到LZ4 writer中  
	_, err = io.Copy(writer, source)  
	if err != nil {  
		return 0, err
	}  
  
	// 关闭writer以完成压缩并刷新缓冲区  
	err = writer.Close()  
	if err != nil {  
		return 0, err
	}

    // 将压缩后的数据写入到目标文件  
	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	
    nBytes, err := compressedData.WriteTo(destination)  
	if err != nil {  
		return 0, err
	}  
	
// 	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func IntToBytes(n int) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func StrToBytes(s string) []byte {
	data := string(s)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func Tmd5() string {
	timeInt := time.Now().Unix()
	return StrToMd5(strconv.Itoa(int(timeInt)))
}

func StrToMd5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func GetWd() string {
	cpath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dataDir := ".xgit"
	dataPath := path.Join(cpath, dataDir)
	_, err = os.Stat(dataPath)
	if err != nil { //文件不存在
		err = os.Mkdir(dataPath, os.ModePerm)
		if err != nil {
			fmt.Printf("permission denied![%v]\n", err)
			panic(err)
		}
	}

	return cpath
}

