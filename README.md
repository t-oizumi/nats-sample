## 動作方法
### 1.docker-compose up
### 2.起動したインスタンスそれぞれにログインする。
###   docker exec -it ID /bin/bash
### 3.publisherからpublishする
###   ./main sub hello1
###   ./main sub hello2
###   ./main sub hello3
### 4.subscriberの起動時オプションについて以下を確認
###   ./main --all sub ←subにpublishされたメッセージを全て取得
###   ./main --seq <seqNo> ←subにpublishされたメッセージのseqNoから取得