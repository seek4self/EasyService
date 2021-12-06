# http server configuration
http:
  ip: "127.0.0.1"
  port: 8080


# zap logger configuration
log:
  # 可使用 "debug", "info", "warn", "error", "dpanic", "panic", "fatal",
  level: "debug"
  prefix: "[{{upper .Name }}]"
  director: "logs"
  # 调用行号
  show_line: false
  # console: 控制台, json: json格式输出
  format: "console"
  # 日志等级颜色
  color: true
  # 是否开启栈追踪，默认开启，若不开启为空字符串
  stacktrace: "error"