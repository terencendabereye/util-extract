package main

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	termColor "github.com/fatih/color"
	"golang.org/x/term"
)

func extractTarGz(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		target := filepath.Join(".", header.Name)
		if header.Typeflag == tar.TypeDir {
			// Create directories
			if err := os.MkdirAll(target, os.ModePerm); err != nil {
				return err
			}
		} else {
			// Extract files
			file, err := os.Create(target)
			if err != nil {
				return err
			}
			defer file.Close()

			if _, err := io.Copy(file, tarReader); err != nil {
				return err
			}
		}

		// Print progress indicator
		terminalWidth, err := getTerminalWidth()
		if err != nil {
			terminalWidth = 0
		}
		c:= termColor.New(termColor.FgYellow)
		fmt.Printf("\r%s", strings.Repeat(" ", terminalWidth))
		fmt.Printf("\rExtracting %s...", c.Sprint(header.Name))
	}

	// Print newline after progress completion
	fmt.Println()

	return nil
}
func getTerminalWidth() (int, error) {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0, err
	}
	return width, nil
}

func extractZip(filename string) error {
	reader, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		target := filepath.Join(".", file.Name)

		if file.FileInfo().IsDir() {
			// Create directories
			if err := os.MkdirAll(target, os.ModePerm); err != nil {
				return err
			}
		} else {
			// Extract files
			rc, err := file.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			file, err := os.Create(target)
			if err != nil {
				return err
			}
			defer file.Close()

			if _, err := io.Copy(file, rc); err != nil {
				return err
			}
		}

		// Print progress indicator
		fmt.Printf("\rExtracting %s...", file.Name)
	}

	// Print newline after progress completion
	fmt.Println()

	return nil
}

func main() {
	for _, arg := range os.Args[1:] {
		switch {
		case strings.HasSuffix(arg, ".tar.gz"):
			if err := extractTarGz(arg); err != nil {
				fmt.Printf("Error extracting %s: %v\n", arg, err)
			}
		case strings.HasSuffix(arg, ".zip"):
			if err := extractZip(arg); err != nil {
				fmt.Printf("Error extracting %s: %v\n", arg, err)
			}
		default:
			// Use external command for other formats (e.g., xz, gzip)
			cmd := exec.Command("tar", "-xf", arg)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Printf("Error extracting %s: %v\n", arg, err)
			}
		}
	}
	c:= termColor.New(termColor.FgGreen)
	c.Printf("Finished\n");
}
