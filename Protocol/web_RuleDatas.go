package Protocol

type RuleData struct {
	Name string
	Type string
	Rule string
}

var RuleDatas = []RuleData{
	{"bootstrap", "body", "(bootstrap)"},
	{"Nextcloud", "body", "(Nextcloud)"},
	{"兰空图床", "body", "(兰空图床|Lsky Pro)"},
	{"Exchange", "body", "(Outlook)|/owa/"},
	{"APACHE-ActiveMQ", "body", "(Apache ActiveMQ)"},
	{"Jetty", "body", "(Powered by Jetty)"},
	{"华夏ERP", "body", "(jshERP-boo)"},
	{"Lightdash", "body", "(Lightdash)"},
	{"Apache-storm", "body", "(Storm UI)"},
	{"HiveServer", "body", "(HiveServer)"},
	{"D-Link-Route", "server", "HTTPD_ac 1.0"},
	{"Kibana", "body", "(kibanaLegacy|.kbnLoader)"},
	{"Node-Exporter", "body", "(Node Exporter)"},
	{"Prometheus", "body", "(Prometheus)"},
	{"Ambari", "body", "(Ambari uses and their respective authors)"},
	{"docker-registry", "body", "(docker-registry-frontend)"},
	{"ThinkPHP", "body", "(ApiAdmin开发维护团队|ThinkPHP)"},
	{"亿邮电子邮件系统", "body", "(亿邮电子邮件系统)"},
	{"用友NC", "body", "(url=nccloud|YONYOU NC)"},
	{"小米路由器", "headers", "(Micgi-Preload)"},
	{"小米路由器", "body", "(<title>小米路由器</title>)"},
	{"Everything", "body", "Everything"},
	{"Jenkins", "headers", "Jenkins"},
	{"Kafka-Manager", "headers", "Kafka-Manager"},
	{"vhost-Panel", "body", "This is the default server vhost"},
	{"Baidu-ECharts", "body", "(echarts\\.js|echarts\\.min\\.js)"},
	{"Kafka-Manager[未授权访问]", "body", "(<title>Kafka Manager</title>)"},
	{"Jumpserver堡垒机", "body", "(Jumpserver|全球首款完全开源的堡垒机)"},
	{"天融信-入侵检测系统TopSentry", "body", "(<title>天融信入侵检测系统TopSentry</title>|TopSentry)"},
	{"天融信-入侵防御系统TopIDP", "body", "(<title>天融信入侵防御系统TopIDP</title>|TopIDP)"},
	{"天融信-TOPSEC", "body", "(/site/js/xblib.js)"},
	{"绿盟科技认证系统", "body", "用户认证 - NSFOCUS NF|绿盟科技认证系统"},
	{"HIKVISION-视频监控系统", "body", "(/doc/page/login.asp)"},
	{"Coremail邮件系统", "body", "(Coremail邮件系统)"},
	{"阿里云OSS", "body", "(<Code>AccessDenied</Code>)"},
	{"frp", "body", "(Faithfully yours, frp)"},
	{"Spark", "body", "(serverSparkVersion)"},
	{"Apache-Spark", "body", "(Spark Worker at)"},
	{"数据库|CouchDB[未授权访问]", "body", "(couchdb.*?uuid)"},
	{"Hadoop-Administration", "body", "(DataNode Information|Hadoop Administration)"},
	{"go-pprof", "body", "(Node Exporter)"},
	{"Django", "body", "(DisallowedHost)"},
	{"Grafana", "body", "(Grafana)"},
	{"深信服-上网行为管理系统", "body", "(Internet Authentication System|<title>上网认证系统</title>)"},
	{"深信服-防火墙", "body", "(function\\(str, key\\))"},
	{"任子行-防火墙", "body", "(任子行下一代防火墙)"},
	{"NPS内网穿透", "body", "(ehang.io/nps|<title>nps error</title>)"},
	{"华为路由器", "body", "(/css/cat_public.css.cgz|/api/system/routerstatus)"},
	{"Shiro", "cookie", "(rememberMe)"},
	{"Struts2", "body", "(\\.action|\\.do|Struts2 Showcase|org.apache.struts2|Struts Problem Report|struts.devMode|struts-tags|There is no Action mapped for namespace)"},
	{"qBittorrent-Web-UI", "body", "(qBittorrent Web UI)"},
	{"织梦内容管理系统", "body", "(织梦内容管理系统)"},
	{"宝塔", "body", "(app.bt.cn/static/app.png|安全入口校验失败|<title>入口校验失败</title>|href=\"http://www.bt.cn/bbs|恭喜, 站点创建成功！)"},
	{"启明防火墙", "body", "(/cgi-bin/webui?op=get_product_model)"},
	{"数据库|ElasticSearch[未授权访问]", "body", `(?s)"name"\s*:\s*"[^"]*".*?"cluster_name"\s*:\s*"[^"]*".*?"cluster_uuid"\s*:\s*"[^"]*".*?"number"\s*:\s*"[^"]*"`},
	{"数据库|ElasticSearch", "body", "security_exception"},
	{"AList", "body", "(由 AList 驱动|alist_pic.js)"},
	{"数据库「MongoDB」", "body", `(MongoDB)`},
	{"ZABBIX-监控系统", "body", "(Zabbix SIA|<title>omni: Zabbix</title>|images/general/zabbix.ico|Zabbix SIA|zabbix-server: Zabbix)"},
	{"phpinfo", "body", "(<title>(.*phpinfo.*)</title>)"},
	{"Tomcat", "body", "(<title>(.*Tomcat.*)</title>|Manager App|Apache Tomcat)"},
	{"天融信VPN", "body", "(/portal_default/index.html)"},
	{"SonarQube-代码管理", "body", "(SonarQube)"},
	{"TamronOS-IPTV", "body", "(TamronOS IPTV系统|IPTV/VOD系统)"},
	{"安恒信息-明御运维审计与风险控制系统", "body", "(<title>(.*明御运维审计.*)</title>)"},
	{"知道创宇-WEBSOC", "body", "(<title>\\s*登录\\s*-\\s*WebSOC知道网站立体监控系统\\s*-\\s*知道创宇\\s*</title>)"},
	{"联软网络智能准入系统", "body", "(下载助手安装包|href=.*portal/resources/css/auth_login.css?)"},
	{"联软IT安全运维管理系统", "body", "(/manager/Resource/Js/ajax.js)"},
	{"phpPgAdmin", "body", "(<title>phpPgAdmin</title>)"},
	{"ThinkPHP", "headers", "(ThinkPHP)"},
	{"ActiveMQ", "headers", "(ActiveMQRealm)"},
	{"VisualSVN", "headers", "(VisualSVN Server)"},
	{"ThinkPHP", "cookie", "(think_var)"},
	{"vmware-ESX", "cookie", "(vmware_esx_host|vmware_client)"},
	{"vmware-ESX", "body", "(URL='/ui')"},
	{"RouterOS", "body", "(mikrotik_logo.png)"},
	{"jQuery", "body", "(jquery.*?js)"},
	{"MinIO-Console", "body", "(MinIO Console)"},
	{"Spring", "body", "(Whitelabel Error Page|No message available|令牌不能为空|timestamp.*?404.*?Not Found|timestamp.*?500.*?Internal Server Error)"},
	{"H3C-Switch", "body", "(../images/Cnlink.jpg|Web user login)"},
	{"H3C-安全管理平台", "body", "(安全产品管理平台|<title>Web managerment Home</title>)"},
	{"360网站卫士", "body", "(webscan.360.cn/status/pai/hash|wzws-waf-cgi|zhuji.360.cn/guard/firewall/stopattack.html)"},
	{"绿盟防火墙", "body", "(NSFOCUS NF)"},
	{"Safedog", "body", "(404.safedog.cn/images/safedogsite/broswer_logo.jpg)"},
	{"阿里云", "body", "(errors.aliyun.com)"},
	{"Portainer(Docker管理)", "body", "(portainer.updatePassword|portainer.init.admin)"},
	{"Nexus", "body", "(Nexus Repository Manager)"},
	{"Harbor", "body", "(<title>Harbor</title>)"},
	{"禅道", "body", "(/theme/default/images/main/zt-logo.png)"},
	{"协众OA", "body", "(Powered by 协众OA)"},
	{"xxl-job", "body", "(分布式任务调度平台XXL-JOB)"},
	{"atmail-WebMail", "body", "(/index.php/mail/auth/processlogin|Powered by Atmail)"},
	{"weblogic", "body", "(/console/framework/skins/wlsconsole/images/login_WebLogic_branding.png|Welcome to Weblogic Application Server|<i>Hypertext Transfer Protocol -- HTTP/1.1</i>)"},
	{"致远OA", "body", "(/seeyon/common/|/seeyon/USER-DATA/IMAGES/LOGIN/login.gif)"},
	{"discuz", "body", "(content=\"Discuz! X\")"},
	{"Typecho", "body", "(Typecho</a>)"},
	{"金蝶EAS", "body", "(easSessionId)"},
	{"phpMyAdmin", "body", "(/themes/pmahomme/img/logo_right.png)"},
	{"H3C-AM8000", "body", "(AM8000)"},
	{"360企业版", "body", "(360EntWebAdminMD5Secret)"},
	{"H3C公司产品", "body", "(service@h3c.com)"},
	{"H3C ICG 1000", "body", "(ICG 1000系统管理)"},
	{"Citrix-Metaframe", "body", "(window.location=\"/Citrix/MetaFrame)"},
	{"H3C ER5100", "body", "(ER5100系统管理)"},
	{"阿里云CDN", "body", "(cdn.aliyuncs.com)"},
	{"CISCO_EPC3925", "body", "(Docsis_system)"},
	{"CISCO ASR", "body", "(CISCO ASR)"},
	{"H3C ER3200", "body", "(ER3200系统管理)"},
	{"万户OA", "body", "(/defaultroot/templates/template_system/common/css/|/defaultroot/scripts/|css/css_whir.css)"},
	{"Spark_Master", "body", "(Spark Master at)"},
	{"nginxWebUI", "body", "(nginxWebUI)"},
	{"华为_HUAWEI_SRG2220", "body", "(HUAWEI SRG2220)"},
	{"蓝凌OA", "body", "(/scripts/jquery.landray.common.js)"},
	{"深信服ssl-vpn", "body", "(login_psw.csp)"},
	{"华为 NetOpen", "body", "(/netopen/theme/css/inFrame.css)"},
	{"Citrix-Web-PN-Server", "body", "(Citrix Web PN Server)"},
	{"juniper_vpn", "body", "(welcome.cgi?p=logo|/images/logo_juniper_reversed.gif)"},
	{"H3C ER8300", "body", "(ER8300系统管理)"},
	{"Citrix-Access-Gateway", "body", "(Citrix Access Gateway)"},
	{"华为 MCU", "body", "(McuR5-min.js)"},
	{"TP-LINK Wireless WDR3600", "body", "(TP-LINK Wireless WDR3600)"},
	{"泛微OA", "body", "(/spa/portal/public/index.js)"},
	{"华为_HUAWEI_ASG2050", "body", "(HUAWEI ASG2050)"},
	{"360网站卫士", "body", "(360wzb)"},
	{"Citrix-XenServer", "body", "(Citrix Systems, Inc. XenServer)"},
	{"H3C ER2100V2", "body", "(ER2100V2系统管理)"},
	{"360站长平台", "body", "(360-site-verification)"},
	{"H3C ER3108GW", "body", "(ER3108GW系统管理)"},
	{"H3C ER3260G2", "body", "(ER3260G2系统管理)"},
	{"H3C ICG1000", "body", "(ICG1000系统管理)"},
	{"CISCO-CX20", "body", "(CISCO-CX20)"},
	{"H3C ER5200", "body", "(ER5200系统管理)"},
	{"linksys-vpn-bragap14-parintins", "body", "(linksys-vpn-bragap14-parintins)"},
	{"360网站卫士常用前端公共库", "body", "(libs.useso.com)"},
	{"H3C ER3100", "body", "(ER3100系统管理)"},
	{"360webfacil_360WebManager", "body", "(publico/template/)"},
	{"Citrix_Netscaler", "body", "(ns_af)"},
	{"H3C ER6300G2", "body", "(ER6300G2系统管理)"},
	{"H3C ER3260", "body", "(ER3260系统管理)"},
	{"华为_HUAWEI_SRG3250", "body", "(HUAWEI SRG3250)"},
	{"exchange", "body", "(/owa/auth.owa|Exchange Admin Center)"},
	{"H3C ER3108G", "body", "(ER3108G系统管理)"},
	{"Citrix-ConfProxy", "body", "(confproxy)"},
	{"360网站安全检测", "body", "(webscan.360.cn/status/pai/hash)"},
	{"H3C ER5200G2", "body", "(ER5200G2系统管理)"},
	{"华为（HUAWEI）安全设备", "body", "(sweb-lib/resource/)"},
	{"华为（HUAWEI）USG", "body", "(UI_component/commonDefine/UI_regex_define.js)"},
	{"H3C ER6300", "body", "(ER6300系统管理)"},
	{"华为_HUAWEI_ASG2100", "body", "(HUAWEI ASG2100)"},
	{"TP-Link 3600 DD-WRT", "body", "(TP-Link 3600 DD-WRT)"},
	{"NETGEAR WNDR3600", "body", "(NETGEAR WNDR3600)"},
	{"H3C ER2100", "body", "(ER2100系统管理)"},
	{"jira", "body", "(jira.webresources)"},
	{"金和协同管理平台", "body", "(金和协同管理平台)"},
	{"Citrix-NetScaler", "body", "(NS-CACHE)"},
	{"通达OA", "body", "(/static/images/tongda.ico|http://www.tongda2000.com|通达OA移动版|Office Anywhere)"},
	{"华为（HUAWEI）Secoway设备", "body", "(Secoway)"},
	{"华为_HUAWEI_SRG1220", "body", "(HUAWEI SRG1220)"},
	{"H3C ER2100n", "body", "(ER2100n系统管理)"},
	{"H3C ER8300G2", "body", "(ER8300G2系统管理)"},
	{"金蝶政务GSiS", "body", "(/kdgs/script/kdgs.js)"},
	{"Jboss", "body", "(Welcome to JBoss|jboss.css)"},
	{"泛微E-mobile", "body", "(Weaver E-mobile|weaver,e-mobile)"},
	{"齐治堡垒机", "body", "(logo-icon-ico72.png|resources/themes/images/logo-login.png)"},
	{"ThinkPHP", "body", "(/Public/static/js/)"},
	{"weaver-ebridge", "body", "(e-Bridge,http://wx.weaver)"},
	{"DWR", "body", "(dwr/engine.js)"},
	{"swagger_ui", "body", "(swagger-ui/css|\"swagger\":|swagger-ui.min.js)"},
	{"大汉版通发布系统", "body", "(大汉版通发布系统|大汉网络)"},
	{"druid", "body", "(druid.index|DruidDrivers|DruidVersion|Druid Stat Index)"},
	{"红帆OA", "body", "(iOffice)"},
	{"VMware vSphere", "body", "(VMware vSphere)"},
	{"打印机", "body", "(更换硒鼓|media/canon.gif|<title>*Brother*</title>|耗材商店|碳粉盒|耗材量严重不足时)"},
	{"finereport", "body", "(isSupportForgetPwd|FineReport,Web Reporting Tool)"},
	{"蓝凌OA", "body", "(蓝凌软件|StylePath:\"/resource/style/default/\"|/resource/customization|sys/ui/extend/theme/default/style/profile.css|sys/ui/extend/theme/default/style/icon.css)"},
	{"GitLab", "body", "(href=\"https://about.gitlab.com/)"},
	{"JQuery-1.7.2", "body", "(/webui/js/jquerylib/jquery-1.7.2.min.js)"},
	{"Hadoop Applications", "body", "(/cluster/app/application)"},
	{"海昌OA", "body", "(/loginmain4/js/jquery.min.js)"},
	{"帆软报表", "body", "(WebReport/login.html|ReportServer)"},
	{"久其财务报表", "body", "(netrep/login.jsp|/netrep/intf)"},
	{"若依管理系统", "body", "(ruoyi/login.js|ruoyi/js/ry-ui.js)"},
	{"启莱OA", "body", "(js/jQselect.js|js/jquery-1.4.2.min.js)"},
	{"智慧校园管理系统", "body", "(DC_Login/QYSignUp)"},
	{"浪潮 ClusterEngineV4.0", "body", "(0;url=module/login/login.html)"},
	{"会捷通云视讯平台", "body", "(him/api/rest/v1.0/node/role|him.app)"},
	{"源码泄露账号密码 F12查看", "body", "(get_dkey_passwd)"},
	{"Smartbi Insight", "body", "(smartbi.gcf.gcfutil)"},
	{"汉王人脸考勤管理系统", "body", "(汉王人脸考勤管理系统|/Content/image/hanvan.png|/Content/image/hvicon.ico)"},
	{"亿赛通-电子文档安全管理系统", "body", "(电子文档安全管理系统|/CDGServer3/index.jsp|/CDGServer3/SysConfig.jsp|/CDGServer3/help/getEditionInfo.jsp)"},
	{"天融信 TopApp-LB 负载均衡系统", "body", "(TopApp-LB 负载均衡系统)"},
	{"中新金盾信息安全管理系统", "body", "(中新金盾信息安全管理系统|中新网络信息安全股份有限公司)"},
	{"好视通", "body", "(深圳银澎云计算有限公司|itunes.apple.com/us/app/id549407870|hao-shi-tong-yun-hui-yi-yuan)"},
	{"蓝海卓越计费管理系统", "body", "(蓝海卓越计费管理系统|星锐蓝海网络科技有限公司)"},
	{"和信创天云桌面系统", "body", "(和信下一代云桌面VENGD|/vesystem/index.php)"},
	{"金山", "body", "(北京猎鹰安全科技有限公司|金山终端安全系统V9.0Web控制台|北京金山安全管理系统技术有限公司|金山V8)"},
	{"WIFISKY-7层流控路由器", "body", "(深圳市领空技术有限公司|WIFISKY 7层流控路由器)"},
	{"MetInfo-米拓建站", "body", "(MetInfo|/skin/style/metinfo.css|/skin/style/metinfo-v2.css)"},
	{"IBM-Lotus-Domino", "body", "(/mailjump.nsf|/domcfg.nsf|/names.nsf|/homepage.nsf)"},
	{"APACHE-kylin", "body", "(url=kylin)"},
	{"C-Lodop打印服务系统", "body", "(/CLodopfuncs.js|www.c-lodop.com)"},
	{"ATLASSIAN-Confluence", "body", "(Atlassian Confluence)"},
	{"HFS", "body", "(href=\"http://www.rejetto.com/hfs/)"},
	{"Jellyfin", "body", "(content=\"http://jellyfin.org\")"},
	{"FIT2CLOUD-JumpServer-堡垒机", "body", "(<title>JumpServer</title>)"},
	{"Nacos", "body", "(<title>Nacos</title>)"},
	{"Pulse Connect Secure", "body", "(/dana-na/imgs/space.gif)"},
	{"h5ai", "body", "(powered by h5ai)"},
	{"天融信脆弱性扫描与管理系统", "body", "(/js/report/horizontalReportPanel.js)"},
	{"天融信网络审计系统", "body", "(onclick=dlg_download())"},
	{"天融信日志收集与分析系统", "body", "(天融信日志收集与分析系统)"},
	{"URP教务系统", "body", "(北京清元优软科技有限公司)"},
	{"科来RAS", "body", "(科来软件 版权所有|i18ninit.min.js)"},
	{"正方OA", "body", "(zfoausername)"},
	{"希尔OA", "body", "(/heeroa/login.do)"},
	{"泛普建筑工程施工OA", "body", "(/dwr/interface/LoginService.js)"},
	{"中望OA", "body", "(/IMAGES/default/first/xtoa_logo.png|/app_qjuserinfo/qjuserinfoadd.jsp)"},
	{"海天OA", "body", "(HTVOS.js)"},
	{"信达OA", "body", "(http://www.xdoa.cn</a>)"},
	{"任我行CRM", "body", "(CRM_LASTLOGINUSERKEY)"},
	{"Spammark邮件信息安全网关", "body", "(/cgi-bin/spammark?empty=1)"},
	{"winwebmail", "body", "(WinWebMail Server|images/owin.css)"},
	{"浪潮政务系统", "body", "(LangChao.ECGAP.OutPortal|OnlineQuery/QueryList.aspx)"},
	{"网神防火墙", "body", "(css/lsec/login.css)"},
	{"帕拉迪统一安全管理和综合审计系统", "body", "(module/image/pldsec.css)"},
	{"蓝盾BDWebGuard", "body", "(BACKGROUND: url(images/loginbg.jpg) #e5f1fc)"},
	{"Huawei SMC", "body", "(Script/SmcScript.js?version=)"},
	{"coremail", "body", "(/coremail/bundle/|contextRoot: \"/coremail\"|coremail/common)"},
	{"activemq", "body", "(activemq_logo|Manage ActiveMQ broker)"},
	{"锐捷数据库安全审计系统", "body", "(锐捷数据库安全审计系统)"},
	{"锐捷网络", "body", "(static/img/title.ico|support.ruijie.com.cn|Ruijie - NBR|eg.login.loginBtn|锐捷网络股份有限公司|/pub/img5/public/no600x140.png)"},
	{"禅道", "body", "(/theme/default/images/main/zt-logo.png|zentaosid|欢迎使用禅道集成运行环境|禅道软件旗下产品)"},
	{"weblogic", "body", "(/console/framework/skins/wlsconsole/images/login_WebLogic_branding.png|Welcome to Weblogic Application Server|<i>Hypertext Transfer Protocol -- HTTP/1.1</i>|<TITLE>Error 404--Not Found</TITLE>|Welcome to Weblogic Application Server|<title>Oracle WebLogic Server 管理控制台</title>)"},
	{"致远OA", "body", "(/seeyon/USER-DATA/IMAGES/LOGIN/login.gif|/seeyon/common/)"},
	{"蓝凌EIS智慧协同平台", "body", "(/scripts/jquery.landray.common.js)"},
	{"深信服ssl-vpn", "body", "(login_psw.csp|loginPageSP/loginPrivacy.js|/por/login_psw.csp)"},
	{"泛微OA", "body", "(/spa/portal/public/index.js|wui/theme/ecology8/page/images/login/username_wev8.png|/wui/index.html#/?logintype=1)"},
	{"Swagger UI", "body", "(/swagger-ui.css|swagger-ui-bundle.js|swagger-ui-standalone-preset.js)"},
	{"金蝶政务GSiS", "body", "(/kdgs/script/kdgs.js|HTML5/content/themes/kdcss.min.css|/ClientBin/Kingdee.BOS.XPF.App.xap)"},
	{"蓝凌OA", "body", "(蓝凌软件|StylePath:\"/resource/style/default/\"|/resource/customization|sys/ui/extend/theme/default/style/icon.css|sys/ui/extend/theme/default/style/profile.css)"},
	{"用友NC", "body", "(Yonyou UAP|YONYOU NC|/Client/Uclient/UClient.dmg|logo/images/ufida_nc.png|iufo/web/css/menu.css|/System/Login/Login.asp?AppID=|/nc/servlet/nc.ui.iufo.login.Index)"},
	{"用友IUFO", "body", "(iufo/web/css/menu.css)"},
	{"TELEPORT堡垒机", "body", "(/static/plugins/blur/background-blur.js)"},
	{"JEECMS", "body", "(/r/cms/www/red/js/common.js|/r/cms/www/red/js/indexshow.js|Powered by JEECMS|JEECMS|/jeeadmin/jeecms/index.do)"},
	{"CMS", "body", "(Powered by .*CMS)"},
	{"目录遍历", "body", "(Directory listing for /)"},
	{"向日葵", "body", "({\"success\":false,\"msg\":\"Verification failure\"})"},
	{"Kubernetes", "body", "(Kubernetes Dashboard</title>|Kubernetes Enterprise Manager|Mirantis Kubernetes Engine|Kubernetes Resource Report)"},
	{"WordPress", "body", "(/wp-login.php?action=lostpassword|WordPress</title>|WordPress站点)"},
	{"RabbitMQ", "body", "(RabbitMQ Management)"},
	{"Spring env", "body", "(logback)"},
	{"ueditor", "body", "(ueditor.all.js|UE.getEditor)"},
	{"亿邮电子邮件系统", "body", "(亿邮电子邮件系统|亿邮邮件整体解决方案)"},
	{"ZABBIX-监控系统", "cookie", "(zbx_session)"},
	{"万户OA", "cookie", "(LocLan)"},
	{"360网站卫士", "headers", "(360wzws|CWAP-waf|zhuji.360.cn|X-Safe-Firewall)"},
	{"绿盟防火墙", "headers", "(NSFocus)"},
	{"Anquanbao", "headers", "(Anquanbao)"},
	{"BaiduYunjiasu", "headers", "(yunjiasu)"},
	{"BigIP", "headers", "(BigIP|BIGipServer)"},
	{"BinarySEC", "headers", "(binarysec)"},
	{"BlockDoS", "headers", "(BlockDos.net)"},
	{"CloudFlare", "headers", "(cloudflare)"},
	{"Cloudfront", "headers", "(cloudfront)"},
	{"Comodo", "headers", "(Protected by COMODO)"},
	{"IBM-DataPower", "headers", "(X-Backside-Transport)"},
	{"DenyAll", "headers", "(sessioncookie=)"},
	{"dotDefender", "headers", "(dotDefender)"},
	{"Incapsula", "headers", "(X-CDN|Incapsula)"},
	{"Jiasule", "headers", "(jsluid=)"},
	{"KONA", "headers", "(AkamaiGHost)"},
	{"ModSecurity", "headers", "(Mod_Security|NOYB)"},
	{"NetContinuum", "headers", "(Cneonction|nnCoection|citrix_ns_id)"},
	{"Newdefend", "headers", "(newdefend)"},
	{"Safe3", "headers", "(Safe3WAF|Safe3 Web Firewall)"},
	{"Safedog", "headers", "(Safedog|WAF/2.0)"},
	{"SonicWALL", "headers", "(SonicWALL)"},
	{"Stingray", "headers", "(X-Mapping-)"},
	{"Sucuri", "headers", "(Sucuri/Cloudproxy)"},
	{"Usp-Sec", "headers", "(Secure Entry Server)"},
	{"Varnish", "headers", "(varnish)"},
	{"Wallarm", "headers", "(wallarm)"},
	{"WebKnight", "headers", "(WebKnight)"},
	{"Yundun", "headers", "(YUNDUN)"},
	{"Yunsuo", "headers", "(yunsuo)"},
	{"Coding pages", "header", "(Coding Pages)"},
	{"Shiro", "headers", "(=deleteMe|rememberMe=)"},
	{"360主机卫士", "headers", "(zhuji.360.cn)"},
	{"Nagios", "headers", "(Nagios Access)"},
	{"泛微OA", "headers", "(ecology_JSessionid)"},
	{"CISCO_VPN", "headers", "(webvpn)"},
	{"o2security_vpn", "headers", "(client_param=install_active)"},
	{"linksys-vpn", "headers", "(linksys-vpn)"},
	{"Jboss", "headers", "(JBoss)"},
	{"泛微E-mobile", "headers", "(EMobileServer)"},
	{"ThinkPHP", "headers", "(ThinkPHP)"},
	{"Laravel", "headers", "(laravel_session)"},
	{"帆软报表", "headers", "(数据决策系统)"},
	{"华夏ERP", "headers", "(华夏ERP)"},
	{"Nagios", "headers", "(nagios admin)"},
	{"weblogic", "headers", "(WebLogic)"},
	{"dubbo", "headers", "(Basic realm=\"dubbo\")"},
	{"DenyAll", "headers", "(sessioncookie=)"},
	{"Gogs简易Git服务", "cookie", "(i_like_gogs)"},
	{"Gitea简易Git服务", "cookie", "(i_like_gitea)"},
	{"Nexus", "cookie", "(NX-ANTI-CSRF-TOKEN)"},
	{"Harbor", "cookie", "(harbor-lang)"},
	{"禅道", "cookie", "(zentaosid)"},
	{"协众OA", "cookie", "(CNOAOASESSID)"},
	{"atmail-WebMail", "cookie", "(atmail6)"},
	{"phpMyAdmin", "cookie", "(pma_lang|phpMyAdmin)"},
	{"金和OA", "cookie", "(ASPSESSIONIDSSCDTDBS)"},
	{"jeesite", "cookie", "(jeesite.session.id)"},
	{"拓尔思SSO", "cookie", "(trsidsssosessionid)"},
	{"拓尔思WCMv7/6", "cookie", "(com.trs.idm.coSessionId)"},
}