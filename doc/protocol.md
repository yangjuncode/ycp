# ycp传输协议



- ycp是一个包含可靠和不可靠传输的协议,参考了kcp(https://github.com/skywind3000/kcp).

- ycp为了方便编码的实现,直接采用了protobuf序列化协议数据
- ycp提供udp一样的功能,并不是所有的数据都需要完全可靠
- ycp除了udp传输外还可以在相同的端口进行tcp连接,以防止udp断流,默认不连接,当udp长时间传输不成功后会自动开启tcp连接



