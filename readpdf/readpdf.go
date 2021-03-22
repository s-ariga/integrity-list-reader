/**
 * Copyright 2021 Seiichi Ariga <seiichi.ariga@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package readpdf

import (
	"fmt"

	"github.com/ledongthuc/pdf"
)

func ReadPdf(path string) (string, error) {

	if !isPdf(path) {
		panic(fmt.Sprintf("拡張子によると、ファイル %s はPDF(.pdf)じゃないっぽい", path))
	}

	// fmt.Printf("%sを読み込み", path)
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// TODO とりあえず、1ページ目しか関係ないっぽい(2021/03)
	const coachesPage = 1
	text := ""

	for pn := 1; pn <= r.NumPage()-coachesPage; pn++ {
		p := r.Page(pn)
		if p.V.IsNull() {
			return "", nil // ページが無い？
		}
		s, err := p.GetPlainText(nil)
		if err != nil {
			return "", err
		}
		text += s
	}
	return text, nil
}

func ReadContentPdf(path string) error {
	if !isPdf(path) {
		panic("PDFじゃないっぽい from 拡張子")
	}

	f, r, err := pdf.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	const coachesPage = 1
	var content pdf.Content

	for pn := 1; pn <= r.NumPage()-coachesPage; pn++ {
		page := r.Page(pn)
		if page.V.IsNull() {
			return nil //ページにコンテンツが無い？
		}

		content = page.Content()
		for _, c := range content.Text {
			fmt.Print(c.S)
		}
	}
	return nil
}

// 拡張子でPDFか軽く調べる
func isPdf(path string) bool {
	return path[len(path)-4:] == ".pdf"
}
