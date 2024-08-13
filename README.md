# パトライトを光らせたい

prometheusのwebhookを受けると、パトライト PHn-3FBにUDPパケットを送って光らせるやつ。
パトライトのIPアドレスはハードコードしている。

## 起動方法

1. git clone https://github.com/onokatio/patlite-restapi
2. docker compose build
3. docker compose up -d

これでwebサーバーが立つ

## 確認方法

以下のようなリクエストをwebサーバーに送ると、パトランプが光る。

```
curl -X POST -d '{"status": "firing", "alerts": [ {"status": "firing", "labels": {"severity": "warning"}} ]}' ?.?.?.?:8085/alert_webhook
```
