# salmon-notify

サーモンランのシフトが開始された時、Discord に通知します

<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">べんり <a href="https://t.co/mrkOlNSYS3">pic.twitter.com/mrkOlNSYS3</a></p>&mdash; 相沢陽菜 (@fetus_hina) <a href="https://twitter.com/fetus_hina/status/1191687022542372865?ref_src=twsrc%5Etfw">November 5, 2019</a></blockquote><!--script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script-->

## Getting Started

### Prerequisites

  - Go 1.13

### Installing

  - `$ go get -u github.com/fetus-hina/salmon-notify`

### Running

  - `$ ./salmon-notify <webhookurl>`

実行時に、30分以内にシフトが開始されていた場合、`webhookurl` に通知が行われます。  
2時間おきなどで自動実行するツール(cronなど)と組み合わせることを前提にしています。

Splatoon2.ink の API を利用しています。  
過度な呼び出しは決して行ってはいけません。

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

Copyright (C) 2019-2022 AIZAWA Hina

## Acknowledgments

  - [Splatoon2.ink](https://splatoon2.ink/)
