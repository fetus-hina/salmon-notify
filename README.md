# salmon-notify

サーモンランのシフトが開始された時、Discord に通知します

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

## Acknowledgments

  - [Splatoon2.ink](https://splatoon2.ink/)
