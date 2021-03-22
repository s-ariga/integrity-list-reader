// Seiichi Ariga <seiichi.ariga@gmail.com>

package main

import (
	"integrity-list-reader/readpdf"
)

func main() {
	// TODO ファイル名をコマンドラインオプションで変えられるように
	err := readpdf.ReadContentPdf("test.pdf")
	if err != nil {
		panic(err)
	}
}
