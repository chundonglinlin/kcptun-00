# kcptun
add config.json for https://github.com/xtaci/kcptun

（项目已废弃，请使用： https://github.com/loveuall/kcptun/tree/router 这个分支支持路由器端的透明代理，服务器端直接使用原作者项目 + socks5-server 即可！）

1. 这个仅仅是个人的便利版本，尽可能不修改原项目逻辑情况下，自带了 socks5 (简化版本的 socks5 来自 https://github.com/shadowsocks/shadowsocks-go) 以及透明代理 (类似 ss-redir 在路由器里运行，暂时不支持 ipv6，有ipv6环境者我可以帮忙提供二进制测试)
2. 实现类似 shadowsocks-go 的配置项，并且服务器端支持多用户多端口，方便很多小白用户。
3. 此项目没有直接从 原大神项目 fork，当时一方面冲突没解决，一方面大家都懂，想低调从事。
4. 小白尽管使用本项目，大神请使用原作者，也请指出项目不足的修改思路，再增加自己的代理层。
5. 路由器用户推荐使用本项目，因为给路由器的配置带来更大的便利，也是这个项目的初衷。

编译请保证 kcp-go urfave/cli 等项目处于最新代码。


