package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Eko", "Kurniawan", "Khannedy"})
	_ = writer.Write([]string{"Budi", "Pratama", "Nugraha"})
	_ = writer.Write([]string{"joko", "moro", "diah"})

	writer.Flush()
}