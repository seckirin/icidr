# Genereate C CIDR - Based IP List

# Installation

```bash
$ go version
  go version go1.22.2 darwin/amd64
$ go install github.com/yuukisec/gcidr
  ...
$ gcidr -h

```

# Example

```bash
$ gcidr -l ips.txt -j -sa -o ccidr.json
$ cat ccidr.json
  [
    {
      "cidr": "x.x.x.0/24",
      "count": 1,
      "ips": [
        "x.x.x.x"
      ]
    },
    {
      "cidr": "x.x.x.0/24",
      "count": 2,
      "ips": [
        "27.22.58.226",
        "27.22.58.231"
      ]
    },
    ...
    {
      "cidr": "x.x.x.0/24",
      "count": 8,
      "ips": [
        "x.x.x.x",
        "x.x.x.x"
        ...
      ]
    }
  ]
```

```bash
$ gcidr -l ips.txt -sa > ccidr.txt
$ cat ccidr.txt
  CIDR: x.x.x.x.0/24, Count: 1 IPs: [x.x.x.x]
  CIDR: 183.56.172.0/24, Count: 2, IPs: [x.x.x.x y.y.y.y]
  ...
  CIDR: 27.22.58.0/24, Count: 8, IPs: [x.x.x.x y.y.y.y z.z.z.z ...]
```