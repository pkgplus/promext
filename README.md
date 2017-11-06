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
      "type": "scalar"
    },
    {
      "name": "cpuUtilizationMax",
      "type": "scalar"
    },
    {
      "name": "cpuUtilizationMin",
      "type": "scalar"
    },
    {
      "name": "memoryUtilization",
      "type": "scalar"
    },
    {
      "name": "memoryUtilizationMax",
      "type": "scalar"
    },
    {
      "name": "memoryUtilizationMin",
      "type": "scalar"
    },
    {
      "name": "diskUtilization",
      "type": "vector"
    }
  ]
}
```

### All Metric Console Query
#### Description
The console query api, can query the total metric current values
#### URL
`GET /console/node/:host/metrics`

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
  "data": [{
    "metric" : {
        "name" : "cpuUtilization",
        "type" : "scalar"
    },
    "value": 0.2384
  },{
    "metric" : {
        "name" : "cpuUtilization",
        "type" : "scalar"
    },
    "value": 0.2384
  },{
    "metric" : {
        "name" : "cpuUtilizationMax",
        "type" : "scalar"
    },
    "value": 0.2384
  },{
    "metric" : {
        "name" : "cpuUtilizationMin",
        "type" : "scalar"
    },
    "value": 0.2384
  },{
    "metric" : {
        "name" : "memoryUtilization",
        "type" : "scalar"
    },
    "value": 0.2384
  },{
    "metric" : {
        "name" : "memoryUtilizationMax",
        "type" : "scalar"
    },
    "value": 0.2384
  },{
    "metric" : {
        "name" : "memoryUtilizationMin",
        "type" : "scalar"
    },
    "value": 0.2384
  },{
    "metric" : {
        "name" : "diskUtilization",
        "type" : "vector"
    },
    "value": {
        "/": 0.2384,
        "/data": 0.2384
    }
  }]
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
// if metric type is "scalar", for example cpuUtilization
{
  "status": "success",
  "data": {
    "metric" : {
       "name" : "cpuUtilization",
       "type" : "scalar"
    },
    "value": 0.2384
  }
}

// if metric type is "vector", for example diskUtilization metric
{
  "status": "success",
  "data": {
    "metric" : {
       "name" : "diskUtilization",
       "type" : "vector"
    },
    "values": {
        "/": 0.2384,
        "/data": 0.2384
    }
  }
}
```

### Single Metric Graph Query
#### Description
The graph query api, can query the single metric current value
#### URL
`GET /graph/node/:host/metrics/:metric`

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
// if metric type is "scalar", for example cpuUtilization
{
  "status": "success",
  "data": [{
    "metric" : {
       "name" : "cpuUtilization",
       "type" : "scalar"
    },
    "values" : [
       [ 1435781430.781, "1" ],
       [ 1435781445.781, "1" ],
       [ 1435781460.781, "1" ]
    ]
  }]
}

// if metric type is "vector", for example diskUtilization metric
{
  "status": "success",
  "data": [{
    "metric" : {
       "name" : "diskUtilization",
       "type" : "vector",
       "param" : "/",
    },
    "values": [
        [ 1435781430.781, "1" ],
        [ 1435781445.781, "1" ],
        [ 1435781460.781, "1" ]
    ]
  },{
    "metric" : {
       "name" : "diskUtilization",
       "type" : "vector",
       "param" : "/data",
    },
    "values": [
        [ 1435781430.781, "1" ],
        [ 1435781445.781, "1" ],
        [ 1435781460.781, "1" ]
    ]
  }]
}
```
