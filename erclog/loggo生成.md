[TOC]

# 概述

读取ERC-20代币的事件日志

# 详细描述

用于获取交易记录等信息

# 流程

参考abi文件夹的生成.

1. 使用sol生成abi数据.

![编译为JSON ABI](./img/abi.png)

2. 编译abi文件.为go文件.
    - $ ./main --abi=erc20.abi --pkg=erclog --out=erc20.go