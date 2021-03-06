# ycp传输协议



- ycp是一个包含可靠和不可靠传输的协议,参考了kcp(https://github.com/skywind3000/kcp).
- ycp为了方便编码的实现,直接采用了protobuf序列化协议数据
- ycp提供udp一样的功能,并不是所有的数据都需要完全可靠
- ycp除了udp传输外还可以在相同的端口进行tcp连接,以防止udp断流,默认不连接,当udp长时间传输不成功后会自动开启tcp连接



## ycp交互说明

- 服务端监听指定的端口(udp+tcp)
- client建立udp socket后,需要先发送握手信息,可用于计算rto,也可以防止意外的udp干扰(攻击),握手时Pktno=0
- client可发送多种类型的数据,udp,ycp,tcp等
  - udp类型,只是发送,能不能收到就靠网络了
  - ycp类型,类似kcp,基本为可靠传输,但是也允许超时失败,同时ycp可以对每个发送的包设置是否需要有序
    - 有序可以保证应用收到的包是按顺序收到的,如果发送了1,2,3,4,5个包,1,3,4,5收到,但是2还没收到,如果是有序,则只要没收到2,则3,4,5不会送给应用,如果不是有序,则3,4,5包可以在收到后马上通知应用
    - ycp是可靠的,但是有些数据如果没在指定的时间内传输,则后面重传输是没有意义的,如实时语音,所以ycp也允许定义超时,如果超时了,则数据丢弃,不会无限重试
  - tcp类型,在tcp socket上传输



ycp 协议命令说明:

| cmd  | 方向 | 描述                  |      |
| :--: | :--: | --------------------- | ---- |
|      |      |                       |      |
|  0   | 双向 | 发送单个ycp单独数据包 |      |
|  1   | 双向 | 确认收到数据包        |      |
|  5   | 双向 | 发送udp数据包         |      |
|  15  | 双向 | 握手探测              |      |
|      |      |                       |      |
|      |      |                       |      |





