package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/fs"
	"math"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("------the resule------")
	// total_file := 0
	// total_dir := 0
	// total_size := 0

	allArgs := os.Args

	singleArg := allArgs[1]
	//
	// result, err := os.ReadDir(singleArg)
	//
	// if err != nil {
	// 	fmt.Println("error occured while reading", err)
	// 	return
	// }
	//
	// for _, val := range result {
	// 	// if err != nil {
	// 	// 	fmt.Println("err", err)
	// 	// 	return
	// 	// }
	//
	// 	if val.Type().IsRegular() {
	//
	// 		total_file++
	// 		hash, err := fileHash(val.Name())
	// 		if err != nil {
	// 			fmt.Println("error in hash", err)
	// 			return
	// 		}
	//
	// 		size, err := os.Stat(val.Name())
	// 		if err != nil {
	// 			fmt.Println("error in size", err)
	// 			return
	// 		}
	// 		fmt.Println("\nhash:", hash)
	// 		fmt.Println("size:", size.Size())
	// 		total_size += int(size.Size())
	// 	} else {
	// 		total_dir++
	// 	}
	//
	// 	fmt.Println("type:", val.Type())
	// 	fmt.Println("name:", val.Name())
	// }
	//
	// fmt.Println("\ntotal file:", total_file, "\ntotal_dir:", total_dir)

	total_size, err := DirSize(singleArg)

	if err != nil {
		fmt.Println("error occured in dir size", err)
		return
	}

	fmt.Println("the total size is ", total_size)
}

func fileHash(path string) (string, error) {
	file, err := os.Open(path)

	if err != nil {
		return "", err
	}

	defer file.Close()

	hash := sha256.New()

	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func DirSize(path string) (int64, error) {
	var size int64
	max1, max2, max3 := int64(math.Inf(-1)), int64(math.Inf(-1)), int64(math.Inf(-1))

	var maxfile1, maxfile2, maxfile3 string

	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return err
			}

			each_size := info.Size()
			if each_size > max1 {
				max3 = max2
				maxfile3 = maxfile2
				max2 = max1
				maxfile2 = maxfile1
				max1 = each_size
				maxfile1 = info.Name()
			} else if each_size > max2 {
				max3 = max2
				maxfile3 = maxfile2
				max2 = each_size
				maxfile2 = info.Name()
			} else {
				max3 = each_size
				maxfile3 = info.Name()
			}

			size += info.Size()
		}
		return nil
	})

	fmt.Println("the largest files are")
	fmt.Println("Name   Size")
	fmt.Printf("%s\t%d\n", maxfile1, max1)
	fmt.Printf("%s\t%d\n", maxfile2, max2)
	fmt.Printf("%s\t%d\n", maxfile3, max3)

	return size, err
}
