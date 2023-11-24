/*
Copyright © 2023 TERENCE NDABEREYE ndabereye@gmail.com
*/
package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// compressCmd represents the compress command
var compressCmd = &cobra.Command{
	Use:   "compress filename...",
	Short: "Compress a given file",
	Run: func(cmd *cobra.Command, args []string) {
		zipCompress(args...)
	},
}

func init() {
	rootCmd.AddCommand(compressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func zipCompress(filename ...string) {
	for _, v := range filename {
		if strings.HasSuffix(v, "/") {
			// is dir
			outputFile := fmt.Sprintf("%s.zip", strings.TrimRight(v, "/"))

			archive, err := os.Create(outputFile)
			if err != nil {
				log.Fatalf("Failed to create/open file: %s", err)
			}
			defer archive.Close()
			zipWriter := zip.NewWriter(archive)

			w, err := zipWriter.Create(outputFile)
			if err != nil {
				log.Fatalf("Failed to create zipWriter: %s", err)
			}
			fs.WalkDir(os.DirFS("."), strings.TrimRight(v, "/"), func(path string, d fs.DirEntry, err error) error {
				dirFile := d.Name()
				fmt.Printf("d.Name()=%q path=%q\n", d.Name(), path)
				fmt.Printf("d.IsDir=%t\n", d.IsDir())
				if d.IsDir() {
					fmt.Printf("is dir, skipping...\n")
					return nil
				}
				f, err := os.Open(path)
				if err != nil {
					log.Fatalf("Failed to open %q: %s\n", dirFile, err)
				}
				defer f.Close()
				fmt.Printf("from: %s, to: %s\n", path, outputFile)
				if _, err = io.Copy(w, f); err != nil {
					log.Fatalf("Failed to write dir: %s", err)
				}
				return nil
			})
			zipWriter.Close()
		} else {
			outputFile := fmt.Sprintf("%s.zip", v)
			archive, err := os.Create(outputFile)
			if err != nil {
				log.Fatalf("Failed to create/open file: %s", err)
			}
			defer archive.Close()

			zipWriter := zip.NewWriter(archive)

			f, err := os.Open(v)
			if err != nil {
				log.Fatalf("Failed to open %q: %s", v, err)
			}
			defer f.Close()
			w, err := zipWriter.Create(v)
			if err != nil {
				log.Fatalf("Failed to create zipWriter: %s", err)
			}
			if _, err = io.Copy(w, f); err != nil {
				log.Fatalf("Failed to write: %s", err)
			}
			zipWriter.Close()
		}

	}
}
