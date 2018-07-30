# CEPH 分布式存储服务器

Ceph 存储集群始于部署一个个 Ceph 节点、网络和 Ceph 存储集群。Ceph 存储集群需要一个 Ceph Monitor 两个 OSD 守护进程。而运行 Ceph 文件系统客户端时，则必须要有元数据服务器（ Metadata Server ）。

- Ceph OSDs: Ceph OSD 守护进程的功能是存储数据，处理数据的复制、恢复、回填、再均衡，并通过检查其他 OSD 守护进程的心跳向 Ceph Monitors 提供一些监控信息。当 Ceph 存储集群设定为有 2 个 OSD 守护进程，集群才能达到 *active+clean* 状态（ Ceph 默认有3个副本，但你可以调整副本数 ）。
- Monitors: Ceph Monitor 维护着展示集群状态的各种图表，包括监视器图、OSD 图、归置组（ PG ）图和 CRUSH 图。Ceph 保存着发生在 Monitors、OSD 和 PG 上的每一次状态变更的历史信息（ 称为 epoch ）。
- MDSs: Ceph 元数据服务器（ MDS ）为 Ceph 文件系统存储元数据（ 也就是说，Ceph 块设备和 Ceph 对象存储不使用 MDS ）。元数据服务器使得 POSIX 文件系统的用户们可以在不对 Ceph 存储集群造成负担的前提下，执行诸如 *ls*、*find* 等 基本命令。

Ceph 把客户端数据保存为存储池内的对象。通过使用 CRUSH 算法，Ceph 可以计算出哪个归置组（ PG ）应该持有指定的对象（ Object ），然后进一步计算出哪个 OSD 守护进程持有该归置组。CRUSH 算法使得 Ceph 存储集群能够动态地伸缩、再均衡和修复。
