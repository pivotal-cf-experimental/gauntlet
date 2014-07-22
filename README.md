#Gauntlet

## A CI results system

### Results

###Record a test result

Request:

```
curl -X POST -d '{
  'pipeline': $GO_PIPELINE_NAME,
  'pipecount': $GO_PIPELINE_COUNTER,
  'stage': $GO_STAGE_NAME,
  'stagecount': $GO_STAGE_COUNTER,
  'jobname': $GO_JOB_NAME,
  'gitinfo': $GO_REVISION,
  'pass': $status
}' $HOST/results
```

Fields:


| field    |Required|type| description |
| ------------|--------|-----------|-----|
| pipeline|no|string| name of pipeline |
| pipecount|no|positive integer| GoCD run count|
| stage|no|string| name of pipeline stage |
| stagecount|no|positive integer| GoCD stage run count |
| jobname|yes|string| name of job |
| gitinfo|no|string; possibly json| repo and revision |
| pass|yes| 'true' or 'false'| success of run |


Response on success:

Status: 201

```
{ "resultid": 12345 }
```

Responses on failure, due to missing data:

Status: 400 Bad Request

```
{ reponse: "Missing required field 'jobname'."}
```

Responses on failure, due to incorrect datatype:

Status: 400 Bad Request

```
{ reponse: "Bad datatype on field 'pipecout': integer required."}
```


### List all results

Request:

```
curl $HOST/results
```

