<!--
 Copyright 2021 Seiichi Ariga <seiichi.ariga@gmail.com>
 
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
 
     http://www.apache.org/licenses/LICENSE-2.0
 
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
-->

# integrity-list-reader

NRAJ(日ラ)インテグリティ教育受講者リストを文字列に変換します。
試合運営のため、PDFのままだと使いにくいので、とりあえず文字へ。

## インテグリティ教育受講者リスト公開サイト
https://www.riflesports.jp/member/for_member/member_info/

## ファイルについて

正式なフォーマットは分からないので、現状のファイルを見た理解

* PDF
* 内容の構成
  * 前半 受講者の日ラID(8桁の数字)リスト <- 必要なのはこちら
  * 後半 コーチ名リスト
  * 現在は```ページ数-1```が前半
  * 前半の最初に```ＮＲＡＪインテグリティ教育　　２０１９-２０２０認定コーチも含む受講報告済みＩＤリスト　　　{日付}報告到着分```

## 仕様

上記前半部分の日ラIDリストだけを、文字列として1行1IDで表示
