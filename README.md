# Version Compliance Checker

A Golang web-service to validate mobile app version compliance

[![Docker Repository on Quay](https://quay.io/repository/swade1987/version-compliance-checker/status "Docker Repository on Quay")](https://quay.io/repository/swade1987/version-compliance-checker)

## Requirements

Two environments variables are required for this application to work.
The values of the variables stipulate the required versions for both iPhone and Android.

```
$ IOS_REQUIRED_VERSION
$ ANDROID_REQUIRED_VERSION
```

## Endpoints

### /metrics

This endpoint provides prometheus metrics about the running service.

### /validate

This endpoint is a `POST` method with the following **mandatory** body fields.

```
device_type             # Possible options are ios or android
current_version         # Must be a semantic version
```

The request must be sent using the structure below

```
$ curl -XPOST -d '{"device_type":"android", "current_version": "1.1.0"}' http://localhost:8080/validate
```

Possible responses are:

```
{"compliant":true}

{"compliant":false,"required_version":"2.2.2"}

{"error":"The current version provided (xyz) is not a valid semantic version."}

{"error":"device mapping not found"}
```

## Performance

The application has been load tested with 700 current users accessing it using [Vegeta](https://github.com/tsenart/vegeta) (see below).

![Alt text](img/vegeta-plot.png?raw=true "Load Test")