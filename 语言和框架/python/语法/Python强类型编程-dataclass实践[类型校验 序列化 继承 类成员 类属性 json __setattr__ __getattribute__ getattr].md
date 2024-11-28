python: 3.9.20

```py
import json

from typing import ClassVar
from dataclasses import dataclass, field, fields, _FIELDS

@dataclass
class NodeMetricsBase():
    def __setattr__(self, name:str, value) -> None:
        '''
        类型校验, 非法类型直接忽略, 只打印日志, 不抛异常.
        每次设置成员属性时就会调用到该方法
        '''
        value_type = type(value)
        ## 允许 None 值, 不报错, 但是会忽略
        if value_type == type(None): return
        ## _FIELDS 包含所有 dataclass 的字段
        field_type = getattr(self, _FIELDS).get(name).type
        if value_type != field_type:
            print(f'invalid value type of field {name}, except {field_type}, got {value_type}')
            return
        super().__setattr__(name, value)
    def metadata(self):
        result = []
        for field in fields(self):
            result.append(dict(field.metadata))
        return result
    def toJson(self):
        '''
        json 序列化
        dataclass 仍然不支持直接使用 json.dumps(实例对象), 只通过这种方式找补一下
        使用时可以 json.dumps(实例对象.toJson())
        '''
        result = {}
        for field in fields(self):
            key = field.metadata.get('name')
            ## 返回的 dict 对象需要使用 'MEM.IDLE' 这个字符串为 key,
            ## 需要通过 metadata 信息进行转换
            result[key] = self.__getattribute__(field.name)
        return result

## dataclass 继承
@dataclass
class NodeMetricsBasic(NodeMetricsBase):
    ## 返回的 dict 对象需要使用 'MEM.IDLE' 这个字符串为 key, 需要通过 metadata 信息进行转换
    ## TypeDict 也不支持(类似于 golang 中的 json tag)
    cpu:float      = field(default=None, metadata={'name':'CPU',      'desc':'CPU使用率'})
    mem_idle:float = field(default=None, metadata={'name':'MEM.IDLE', 'desc':'空闲内存占比'})
@dataclass
class NodeMetricsExtDisk(NodeMetricsBase):
    disk_root:float = field(default=None, metadata={'name':'DISK.ROOT', 'desc':'根文件系统分区空间利用率'})
    disk_boot:float = field(default=None, metadata={'name':'DISK.BOOT', 'desc':'启动分区空间利用率'})
    disk_data:float = field(default=None, metadata={'name':'DISK.DATA', 'desc':'数据分区空间利用率'})
    ## 类属性, 不会被 fields() 遍历到
    mountpoint_map: ClassVar = {
        '/':            'disk_root',
        '/boot':        'disk_boot',
        '/data':        'disk_data',
    }
    def set_attr(self, mountpoint, val):
        attr = NodeMetricsExtDisk.mountpoint_map.get(mountpoint)
        if attr == None: return
        self.__setattr__(attr, val)

## dataclass 会自动的生成构造函数和默认值, 更贴近其他强类型语言的使用方式.
## 要求: 必须提供默认值, 如果无值则需要展示为 None.
@dataclass
class NodeMetrics:
    ## 自定义类型成员默认值, 需要通过 default_factory 指定.
    basic:NodeMetricsBasic      = field(default_factory=NodeMetricsBasic)
    external:NodeMetricsExtDisk = field(default_factory=NodeMetricsExtDisk)

    def metadata(self):
        '''
        自定义 json key 结构, 包含 key 的描述信息
        '''
        result = {
            'basic': self.basic.metadata(),
            'external': {
                'disk': self.external.metadata(),
            }
        }
        return result
    def toJson(self) -> dict:
        '''
        json 结构扁平化, 不想要像 class 成员结构那样, 而是所有 key 都放在同一层级.
        '''
        result = {}
        result.update(self.basic.toJson())
        result.update(self.external.toJson())
        return result

if __name__ == '__main__':
    nodeInfo = NodeMetrics()
    ## nodeInfo.basic.cpu = 12 ## 日志报错
    nodeInfo.basic.cpu = 12.34
    nodeInfo.external.set_attr('/boot', 12.34)
    print(nodeInfo.metadata())
    print(nodeInfo.toJson())

'''
{'basic': [{'name': 'CPU', 'desc': 'CPU使用率'}, {'name': 'MEM.IDLE', 'desc': '空闲内存占比'}], 'external': {'disk': [{'name': 'DISK.ROOT', 'desc': '根文件系统分区空间利用率'}, {'name': 'DISK.BOOT', 'desc': '启动分区空间利用率'}, {'name': 'DISK.DATA', 'desc': '数据分区空间利用率'}]}}
{'CPU': 12.34, 'MEM.IDLE': None, 'DISK.ROOT': None, 'DISK.BOOT': 12.34, 'DISK.DATA': None}
'''

```
