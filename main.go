package main

import (
    "fmt"
    "log"
	"io"
    "os"
)

func main() {
    srcName := "source.txt"
    dstName := "destination.txt"

    written, err := CopyFile(dstName, srcName)
    if err != nil {
        log.Fatalf("CopyFile failed: %v", err)
    }

    fmt.Printf("Copied %d bytes from %s to %s\n", written, srcName, dstName)
}


func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return 0, err
    }
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil {
        return 0, err
    }
    defer dst.Close()

    written, err = io.Copy(dst, src)
    if err != nil {
        return 0, err
    }

    // Ensure the content is flushed to disk
    err = dst.Sync()
    if err != nil {
        return 0, err
    }

    return written, nil
}