#### Style 1
```json
{
  "cluster_property_set": [
    {
      "id": "cib-bootstrap-options",
      "nvpair": [
        {
          "name": "have-watchdog",
          "id": "cib-bootstrap-options-have-watchdog",
          "value": "false"
        },
        {
          "name": "dc-version",
          "id": "cib-bootstrap-options-dc-version",
          "value": "2.0.0+20181108.62ffcafbc-1.1-2.0.0+20181108.62ffcafbc"
        },
        {
          "name": "cluster-infrastructure",
          "id": "cib-bootstrap-options-cluster-infrastructure",
          "value": "corosync"
        },
        {
          "name": "cluster-name",
          "id": "cib-bootstrap-options-cluster-name",
          "value": "hacluster"
        },
        {
          "name": "placement-strategy",
          "id": "cib-bootstrap-options-placement-strategy",
          "value": "balanced"
        }
      ]
    }
  ]
}
```
#### Style 2
```json
{
  "cluster_property_set":[
    "cib-bootstrap-options": {
      "have-watchdog": "false",
      "dc-version": "2.0.0+20181108.62ffcafbc-1.1-2.0.0+20181108.62ffcafbc",
      "cluster-infrastructure": "corosync",
      "cluster-name": "hacluster",
      "placement-strategy": "balanced"
    }
  ]
}

```
```json
{
  "meta_attributes": [
    {
      "id": "op-options",
      "nvpair": [
        {
          "name": "timeout",
          "id": "op-options-timeout",
          "value": "600"
        },
        {
          "name": "record-pending",
          "id": "op-options-record-pending",
          "value": "true"
        }
      ]
    }
  ]
}
```
