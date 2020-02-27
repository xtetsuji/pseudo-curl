# pseudo-curl

pseudo-curl は擬似的な curl コマンドです。

サーバの OpenSSL が古すぎて、curl が新しい TLS 対応をしていても libssl.so が古いために curl が新しい TLS をしゃべれない……というケースを想定しています。
