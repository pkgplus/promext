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
    "cpuUtilization",
    "cpuUtilizationMax",
    "cpuUtilizationMin",
    "memoryUtilization",
    "memoryUtilizationMax",
    "memoryUtilizationMin",
    "diskUtilization"
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
  "data": {
    "cpuUtilization": 0.2384,
    "cpuUtilizationMax": 0.2384,
    "cpuUtilizationMin": 0.2384,
    "memoryUtilization": 0.2384,
    "memoryUtilizationMax": 0.2384,
    "memoryUtilizationMin": 0.2384,
    "diskUtilization": 0.2384
  }
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
  "data": 0.2384
}
```
