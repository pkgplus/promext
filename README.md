# promext
## rules
Found it from the directory `rules`.
## API
### Metric MetaData
#### Description
The meta data of all support metrics
#### URL
`GET /api/v1/meta/metrics`

#### Response
```
{
  "status": "success" | "error",
  "data": <data>, // the list of support metric

  // Only set if status is "error". The data field may still hold
  // additional data.
  "errorType": "<string>",
  "error": "<string>"
}
```

#### Response Example
```json
{
  "status": "success",
  "data": [
    {
      "name": "cpuUtilization",
      "type": "current",
      "dsl": "node_cpu_not_idle_rate{$filter}*100"
    },
    {
      "name": "cpuUtilizationAvg",
      "type": "range",
      "dsl": "avg_over_time(node_cpu_not_idle_rate{$filter}[$duration])*100"
    },
    {
      "name": "cpuUtilizationMax",
      "type": "range",
      "dsl": "max_over_time(node_cpu_not_idle_rate{$filter}[$duration])*100"
    },
    {
      "name": "cpuUtilizationMin",
      "type": "range",
      "dsl": "min_over_time(node_cpu_not_idle_rate{$filter}[$duration])*100"
    },
    {
      "name": "memoryUtilization",
      "type": "current",
      "dsl": "100 - ( node_memory_MemAvailable{$filter} OR node_memory_MemAvailable_ext{$filter})/ node_memory_MemTotal{$filter} * 100"
    },
    {
      "name": "memoryUtilizationAvg",
      "type": "range",
      "dsl": "100 - (avg_over_time(node_memory_MemAvailable{$filter}[$duration]) OR avg_over_time(node_memory_MemAvailable_ext{$filter}[$duration])) / node_memory_MemTotal{$filter} * 100"
    },
    {
      "name": "memoryUtilizationMax",
      "type": "range",
      "dsl": "100 - (max_over_time(node_memory_MemAvailable{$filter}[$duration]) OR max_over_time(node_memory_MemAvailable_ext{$filter}[$duration])) / node_memory_MemTotal{$filter} * 100"
    },
    {
      "name": "memoryUtilizationMin",
      "type": "range",
      "dsl": "100 - (min_over_time(node_memory_MemAvailable{$filter}[$duration]) OR min_over_time(node_memory_MemAvailable_ext{$filter}[$duration])) / node_memory_MemTotal{$filter} * 100"
    },
    {
      "name": "diskUtilization",
      "type": "current",
      "dsl": "100 - node_filesystem_free{$filter, fstype!~\"rootfs|selinuxfs|autofs|rpc_pipefs|tmpfs|iso.+\"} / node_filesystem_size * 100"
    },
    {
      "name": "diskUtilizationMax",
      "type": "range",
      "dsl": "100 - (max_over_time(node_filesystem_free{$filter, fstype!~\"rootfs|selinuxfs|autofs|rpc_pipefs|tmpfs|iso.+\"}[$duration])) / node_filesystem_size * 100"
    }
  ]
}
```

### All Metric Current Query
#### Description
The console query api, can query the total metric current values
#### URL
`GET /current/node/:host/metrics`

### URL PARAMTER

#### Response
```
{
  "status": "success" | "error",
  "data": {
    "<metric>": "<metric-value>",
    "<metric2>": <metric-value2>,
    ...
  },

  // Only set if status is "error". The data field may still hold
  // additional data.
  "errorType": "<string>",
  "error": "<string>"
}
```

#### Response Example
```json
{
  "status": "success",
  "data": [
    {
      "metric": {
        "_name_": "cpuUtilization",
        "cluster": "QD",
        "instance": "10.138.16.188",
        "job": "linux",
        "project": "MONITOR-PROD"
      },
      "value": [
        1511256351.109,
        "3.012760416665604"
      ]
    },
    {
      "metric": {
        "_name_": "memoryUtilization",
        "cluster": "QD",
        "instance": "10.138.16.188",
        "job": "linux",
        "project": "MONITOR-PROD"
      },
      "value": [
        1511256445.479,
        "46.9969964779287"
      ]
    },
    {
      "metric": {
        "_name_": "diskUtilization",
        "cluster": "QD",
        "device": "/dev/sda1",
        "fstype": "xfs",
        "instance": "10.138.16.188",
        "job": "linux",
        "mountpoint": "/boot",
        "project": "MONITOR-PROD"
      },
      "value": [
        1511256470.75,
        "29.472296537551856"
      ]
    },
    {
      "metric": {
        "_name_": "diskUtilization",
        "cluster": "QD",
        "device": "/dev/mapper/centos-root",
        "fstype": "xfs",
        "instance": "10.138.16.188",
        "job": "linux",
        "mountpoint": "/",
        "project": "MONITOR-PROD"
      },
      "value": [
        1511256470.75,
        "12.086414651922965"
      ]
    },
    {
      "metric": {
        "_name_": "diskUtilization",
        "cluster": "QD",
        "device": "/dev/mapper/centos-var",
        "fstype": "xfs",
        "instance": "10.138.16.188",
        "job": "linux",
        "mountpoint": "/var",
        "project": "MONITOR-PROD"
      },
      "value": [
        1511256470.75,
        "1.5972357882398853"
      ]
    },
    {
      "metric": {
        "_name_": "diskUtilization",
        "cluster": "QD",
        "device": "/dev/mapper/centos-data",
        "fstype": "xfs",
        "instance": "10.138.16.188",
        "job": "linux",
        "mountpoint": "/data",
        "project": "MONITOR-PROD"
      },
      "value": [
        1511256470.75,
        "6.748608893397858"
      ]
    },
    {
      "metric": {
        "_name_": "diskUtilization",
        "cluster": "QD",
        "device": "/dev/mapper/centos-home",
        "fstype": "xfs",
        "instance": "10.138.16.188",
        "job": "linux",
        "mountpoint": "/home",
        "project": "MONITOR-PROD"
      },
      "value": [
        1511256470.75,
        "0.1753263251379451"
      ]
    },
    {
      "metric": {
        "_name_": "diskUtilization",
        "cluster": "QD",
        "device": "/dev/mapper/centos-tmp",
        "fstype": "xfs",
        "instance": "10.138.16.188",
        "job": "linux",
        "mountpoint": "/tmp",
        "project": "MONITOR-PROD"
      },
      "value": [
        1511256470.75,
        "0.517217142577735"
      ]
    }
  ]
}
```

### Single Metric Console Query
#### Description
The console query api, can query the single metric current value
#### URL
`GET /console/node/:host/metrics/:metric`

#### Response
```
{
  "status": "success" | "error",
  "data":  "<metric-value>",

  // Only set if status is "error". The data field may still hold
  // additional data.
  "errorType": "<string>",
  "error": "<string>"
}
```

#### Response Example
```json
{
  "status": "success",
  "data": [
    {
      "metric": {
        "_name_": "cpuUtilization",
        "cluster": "QD",
        "instance": "10.138.16.188",
        "job": "linux",
        "project": "MONITOR-PROD"
      },
      "value": [
        1511256351.109,
        "3.012760416665604"
      ]
    }
  ]
}
```

### Single Metric Graph Query
#### Description
The graph query api, can query the single metric current value
#### URL
`GET /range/node/:host/metrics/:metric`

#### Url Parameters
| Field    | Require | DataType      |  Name       | Description |
| ---------| ------- |:-------------:| :---------- |:--------|
| start    |   Y     | string        | Start timestamp   | unix_timestamp, eg:1509959076 |
| end      |   Y     | string        | End timestamp    | unix_timestamp, eg:1509959076 |
| step     |   Y     | string        | Query resolution step width.  | eg: 5m |

#### Response
```
{
  "status": "success" | "error",
  "data":  "<metric-value>",

  // Only set if status is "error". The data field may still hold
  // additional data.
  "errorType": "<string>",
  "error": "<string>"
}
```

#### Response Example
```json
{
  "status": "success",
  "data": [
    {
      "metric": {
        "_name_": "cpuUtilizationMax",
        "cluster": "QD",
        "instance": "10.138.16.188",
        "job": "linux",
        "project": "MONITOR-PROD"
      },
      "values": [
        [
          1511252715,
          "3.517968750001186"
        ],
        [
          1511256315,
          "4.196614583331282"
        ]
      ]
    }
  ]
}
```
