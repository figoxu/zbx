# zbx

## ztool
zabbix_agentd的工具，将普通http接口的json属性作为zabbix监控的keys使用.
配置方法如下:
* 在/etc/zabbix/zabbix_agentd.conf 配置文件里，打开 Include=/etc/zabbix/zabbix_agentd.conf.d/ 选项
* 在/etc/zabbix/zabbix_agentd.conf.d/目录下，创建xxx.conf文件
* xxx.conf里的配置项范例如下:
UserParameter=testb[*],ztool -url http://localhost:10061/zabbix/vs -prop $1

