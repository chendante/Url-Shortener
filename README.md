# Url-Shortener
## 项目结构
- model/shortenUrl.go: 与MySql、Redis、kafka等组件交互的方法
- model/shortenUrlGenertor.go: 短链接生成算法，来自[short-url](https://github.com/by-zhang/short-url)
- controller：处理请求
- listener：一个单独的kafka消费者
