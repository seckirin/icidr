# Genereate C CIDR - Based IP List

# Installation

```bash
go install github.com/yuukisec/gcidr@latest
```

# Options

```text
gcidr -h
  -json       Output the result as JSON
  -l string   The path to the IP list file
  -sa-cidr    Sort the result in ascending order by CIDR
  -sa-count   Sort the result in ascending order by count
  -sd-cidr    Sort the result in descending order by CIDR
  -sd-count   Sort the result in descending order by count
```

# Example

**示例一：根据 ips.txt 生成 C 段并根据 CIDR 所包含的 IP 数量降序排序**

```bash
gcidr -l ips.txt -sd-count
```

outpt

```text
CIDR: 111.1.2.0/24, Count: 7, IPs: [111.1.2.1 111.1.2.2 111.1.2.3 111.1.2.7 111.1.2.8 111.1.2.9 111.1.2.10]
CIDR: 1.1.1.0/24, Count: 5, IPs: [1.1.1.1 1.1.1.2 1.1.1.3 1.1.1.6 1.1.1.8]
CIDR: 11.1.2.0/24, Count: 1, IPs: [11.1.2.90]
```

**示例二：根据 ips.txt 生成 C 段并根据 CIDR 升序排序**

```bash
gcidr -l ips.txt -sa-cidr
```

outpt

```text
CIDR: 1.1.1.0/24, Count: 5, IPs: [1.1.1.1 1.1.1.2 1.1.1.3 1.1.1.6 1.1.1.8]
CIDR: 11.1.2.0/24, Count: 1, IPs: [11.1.2.90]
CIDR: 111.1.2.0/24, Count: 7, IPs: [111.1.2.1 111.1.2.2 111.1.2.3 111.1.2.7 111.1.2.8 111.1.2.9 111.1.2.10]
```

示例三：将结果输出为 json 并根据 CIDR 所包含的 IP 数量进行降序排序

```bash
gcidr -l ips.txt -sd-count -json
```

outptu

```text
[
  {
    "cidr": "111.1.2.0/24",
    "count": 7,
    "ips": [
      "111.1.2.1",
      "111.1.2.2",
      "111.1.2.3",
      "111.1.2.7",
      "111.1.2.8",
      "111.1.2.9",
      "111.1.2.10"
    ]
  },
  {
    "cidr": "1.1.1.0/24",
    "count": 5,
    "ips": [
      "1.1.1.1",
      "1.1.1.2",
      "1.1.1.3",
      "1.1.1.6",
      "1.1.1.8"
    ]
  },
  {
    "cidr": "11.1.2.0/24",
    "count": 1,
    "ips": [
      "11.1.2.90"
    ]
  }
]
```

## STDIN

程序支持 STDIN，可以使用 cat 或其他命令将结果作为管道符输入进 GCIDR

```bash
cat ips.txt | gcidr [OPTION]
```