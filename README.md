# Version Compliance Checker

A Golang web-service to validate mobile app version compliance

## Requirements

Two environments variables are required for this application to work.
The values of the variables stipulate the required versions for both iPhone and Android.

```
$ IOS_REQUIRED_VERSION
$ ANDROID_REQUIRED_VERSION
```

## Endpoints

### /status

The `/status` endpoint is there to be used a validate whether or not the version is up.

### /validate

The `/validate` endpoint is a `POST` method with the following **mandatory** body fields.

```
device_type             # Possible options are ios or android
current_version         # Must be a semantic version
```

The request must be sent using the structure below

```
$ curl -XPOST -d '{"device_type":"android", "current_version": "1.1.0"}' http://localhost:8080/devicevalidation/v1/validate
```