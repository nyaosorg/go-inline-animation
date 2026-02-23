Release notes
=============

- 終了処理の際、パターンの消去が非同期で行われて、消去位置が狂う問題を修正 (#6)

v0.3.0
------
Jan 22, 2026

- 構造体 Animation のフィールド Width を省略できるようにした (#1)
  - go-runewidth で自動計測するようにした
- フィールド: Interval を追加 (#2)
- パターン Dots, Bars を用意 (#3)
- 関数 Progress のデフォルトパターンとして Dots を使うようにした (#3)
- 定数 CursorOn と CursorOff を追加 (#4)

v0.2.0
------
Sep 14, 2021

- Go Modules をサポート
- レポジトリ所有者を`zetamatta` から `nyaosorg` へ変更

v0.1.0
------
Jul 2, 2019

- 初版
