package main

 import (
         "compress/gzip"
         "flag"
         "fmt"
         "io"
         "os"
         "strings"
 )

 func main() {

         flag.Parse() // get the arguments from command line

         filename := flag.Arg(0)

         if filename == "" {
                 fmt.Println("Usage : gunzip sourcefile.gz")
                 os.Exit(1)
         }

         gzipfile, err := os.Open(filename)

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         reader, err := gzip.NewReader(gzipfile)
         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }
         defer reader.Close()

         newfilename := strings.TrimSuffix(filename, ".gz")

         writer, err := os.Create(newfilename)

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         defer writer.Close()

         if _, err = io.Copy(writer, reader); err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

 }
