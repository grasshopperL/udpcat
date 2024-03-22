# udpcat
catch udp packet according to giving net device and port

## 诞生背景
上游系统每天不定时发送udp数据包，要求己方系统接收并记录来自上游系统的udp数据。但是己方系统存在丢包问题，因此需要辅助工具记录udp包，本项目作为辅助工具，接收udp包并写入文件。

## 整体思路
1. 使用ticker.C作为轮询器，使用gopacket/pcap监听制定网口，使用gopacket/pcapgo将udp包写入文件。
2. 使用ticker.C每隔两秒执行轮询动作，每次轮询存在以下两种情况：a.监听端口没有udp包，直接返回，继续轮询。b.监听端口存在udp包，此时使用channel阻塞轮询，开始接收并记录udp包，直至channel中所有udp包写入文件，解除阻塞，继续轮询。



