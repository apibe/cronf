ftp信息定时拉取服务

配置文件

FTP(FTP相关配置)
FTP.ADDR: FTP连接地址 host:port
FTP.USERNAME: FTP账号
FTP.PASSWORD: FTP密码

CRON_CONFIG(定时任务相关配置)
CRON_CONFIG.xxx.EXEC: 是否立即执行定时任务
CRON_CONFIG.xxx.JOIN_CRON: 是否开启定时任务
CRON_CONFIG.xxx.CRON: 定时任务表达式
CRON_CONFIG.xxx.FTP_PATH: FTP文件所在路径(目前只支持文件夹)
CRON_CONFIG.xxx.LOCAL_PATH: 本地文件存储路径
CRON_CONFIG.xxx.RETRY.INTERVAL: 重试间隔时间(分钟)
CRON_CONFIG.xxx.RETRY.TIMES: 最大重试次数(次)

CRON表达式参考网址: https://www.pppet.net/
特殊路径拼接表达式如{date:2006}:  参考 go 语言日期格式化规则 https://www.python100.com/html/76436.html

执行步骤
   进入 cron_ftp.exe 所在文件夹，双击


文件同步服务器