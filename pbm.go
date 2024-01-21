package netpbm

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PBM struct {
	Data          [][]bool
	Width, Height int
	MagicNumber   string
}

func ReadPBM(filename string) (*PBM, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return nil, fmt.Errorf("failed to read magic number")
	}
	magicNumber := scanner.Text()

	if magicNumber != "P1" && magicNumber != "P4" {
		return nil, fmt.Errorf("unsupported PBM format: %s", magicNumber)
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 && line[0] != '#' {
			break
		}
		fmt.Println("Skipped line:", line)
	}

	if scanner.Err() != nil {
		return nil, fmt.Errorf("error reading dimensions line: %v", scanner.Err())
	}
	dimensions := strings.Fields(scanner.Text())
	if len(dimensions) != 2 {
		return nil, fmt.Errorf("invalid dimensions line")
	}

	width, err := strconv.Atoi(dimensions[0])
	if err != nil {
		return nil, fmt.Errorf("failed to parse width: %v", err)
	}

	height, err := strconv.Atoi(dimensions[1])
	if err != nil {
		return nil, fmt.Errorf("failed to parse height: %v", err)
	}

	var data [][]bool
	if magicNumber == "P1" {
		for scanner.Scan() {
			line := scanner.Text()
			tokens := strings.Fields(line)
			row := make([]bool, width)
			for i, token := range tokens {
				if i >= width {
					break
				}
				if token == "1" {
					row[i] = true
				} else if token == "0" {
					row[i] = false
				} else {
					return nil, fmt.Errorf("invalid character in data: %s", token)
				}
			}
			data = append(data, row)
		}
	} else if magicNumber == "P4" {
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return &PBM{
		Data:        data,
		Width:       width,
		Height:      height,
		MagicNumber: magicNumber,
	}, nil
}

func (pbm *PBM) Size() (int, int) {
	return pbm.Height, pbm.Width
}
func (pbm *PBM) At(x, y int) bool {
if x < 0 || y < 0 || x >= pbm.Width || y >= pbm.Height {
    return false
}


    return pbm.Data[y][x]
}
func (pbm *PBM) Set(x, y int, valeur bool) {

if x < 0 || y < 0 || x >= pbm.Width || y >= pbm.Height {
    return
}

pbm.Data[y][x] = valeur
}
func (pbm *PBM) Save(filename string) error{
    file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%s\n", pbm.MagicNumber)
	if err != nil {
		return fmt.Errorf("error writing magic number: %v", err)
	}


	_, err = fmt.Fprintf(file, "%d %d\n", pbm.Width, pbm.Height)
	if err != nil {
		return fmt.Errorf("error writing dimensions: %v", err)
	}


	for _, row := range pbm.Data {
		for _, pixel := range row {
			if pbm.MagicNumber == "P1" {
				
				if pixel {
					_, err = fmt.Fprint(file, "1 ")
				} else {
					_, err = fmt.Fprint(file, "0 ")
				}
			} else if pbm.MagicNumber == "P4" {
			
				if pixel {
					_, err = fmt.Fprint(file, "1")
				} else {
					_, err = fmt.Fprint(file, "0")
				}
			}
			if err != nil {
				return fmt.Errorf("error writing data: %v", err)
			}
		}
		if pbm.MagicNumber == "P1" {
	
			_, err = fmt.Fprintln(file)
			if err != nil {
				return fmt.Errorf("error writing data: %v", err)
			}
		}
	}

	fmt.Printf("Image sauvegardée avec succès dans %s\n", filename)

	return nil
}

func (pbm *PBM) Invert() {
 
    for i := 0; i < pbm.Height; i++ {
        for j := 0; j < pbm.Width; j++ {
            pbm.Set(j, i, !pbm.At(j, i))
        }
    }
}
func (pbm *PBM) Flip() {
   
    for i := 0; i < pbm.Height; i++ {
        for j := 0; j < pbm.Width/2; j++ {
            temp := pbm.At(j, i)
            pbm.Set(j, i, pbm.At(pbm.Width-1-j, i))
            pbm.Set(pbm.Width-1-j, i, temp)
        }
    }
}
func (pbm *PBM) SetMagicNumber(magicNumber string) {
    pbm.MagicNumber = magicNumber
}