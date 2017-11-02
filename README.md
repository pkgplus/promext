# promext
## rules
Found it from the directory `rules`.
## API
### Metric MetaData
#### Description
The meta data of all support metrics
#### URL
`Get /api/v1/meta/metrics`

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