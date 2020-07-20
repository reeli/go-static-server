1. 处理 HTTP 请求，读出对应文件，返回 content, serve 本地文件（assets(png,jpg), html, js）
2. 静态服务（提供前端需要的 assets） + 代理服务（转发 API）
2. Handle redirect
3. 服务的优雅关闭
4. 动态替换 HTML 环境变量
5. 了解 makefile

1. 读配置文件，自动 serve static folder 下的文件，配置 fallback 以及静态文件的路由（找对应的文件 js, jpg 或者 html，找不到则返 404）
2. 其他路由，则 fallback 到 index.html
3. 设置 Cache 的规则（expire, cache-control, etag）
4. proxy
