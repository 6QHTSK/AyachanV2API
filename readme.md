# Ayachan Bandori谱面难度分析器

[![codebeat badge](https://codebeat.co/badges/3482bd1e-45d7-4e83-af70-3f1ccb874fcd)](https://codebeat.co/projects/github-com-6qhtsk-ayachan-development)
[![Go Report Card](https://goreportcard.com/badge/github.com/6QHTSK/ayachan)](https://goreportcard.com/report/github.com/6QHTSK/ayachan)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/6QHTSK/ayachan)
![GitHub](https://img.shields.io/github/license/6QHTSK/ayachan)
![Libraries.io dependency status for GitHub repo](https://img.shields.io/librariesio/github/6QHTSK/ayachan)

可对Bandori谱面进行特征提取，并拟合难度

## 使用方法

```bash
docker pull ghcr.io/6qhtsk/ayachan:latest
```

[API 文档](https://api.ayachan.fun/v2/doc/index.html)

## 名词与机理解释

### 标准谱面

是指不需要玩家跨手/出张（左手跨过右手所在轨道）处理，且谱面的多压数不超过2的谱面，以及不存在重叠音符（两个音符在完全相同的轨道，并要求在完全相同的时间击打）


### 基础信息

是否非标准谱面、Note数、谱面时长、每秒平均谱面物件数（NPS）、是否有SP键、BPM情况、每秒平均击打数（HPS）、单位时间最大NPS（MaxScreenNPS）、谱面物件类型分布、谱面物件时间分布。

### 标准谱面额外可计算的信息：

左右手占比、左右手最大移动速度、单指最高每秒击打次数、单手粉键-音符平均间隔、单手音符-粉键平均间隔

### 难度计算：

基于统计回归的原理，通过拟合各个信息在各自难度所处于的位置，对比较的方式为每个信息进行难度标定。

在当前版本，只会对基础信息部分的NPS、HPS、MaxScreenNPS三个维度进行难度回归计算，并对这三个信息的回归计算难度进行加权平均，给出拟合的**谱面总体难度**。

额外可计算信息部分只会给出其相对于当前难度谱面的难度比较情况。例如，是否相对偏高/偏低。