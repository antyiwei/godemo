

### go语言net包rpc远程调用的使用


- go对RPC的支持，支持三个级别：TCP、HTTP、JSONRPC
- go的RPC只支持GO开发的服务器与客户端之间的交互，因为采用了gob编码
