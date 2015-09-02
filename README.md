# Warehouse

Warehouse is a storage system for builds and cached libraries for all apps using nanobox in production. It automatically removes the oldest builds once the limit is reached (currently set at 5 and will be adjustable through the API int he future).

## Configuration:

you start warehouse by passing a config file (warehouse /path/to/config). If the config file is not passed a default set of options will be used.

```
# NA-Warehouse config file

# The address you want me to listen to
# ip and port are important
listenAddr 0.0.0.0:1234

# show a specific amount of logs
# default value is info
logLevel info

# the authentication token
token supersecrettoken

#the folder you want to store the files
dataDir /tmp/warehouse/
```

## Routes:

| Method | Route | Functionality |
| --- | --- | --- |
| GET | /builds/{file} | Retrieve the build by the file name |
| HEAD | /builds/{file} | Retrieve just the head of the file which includes NAME and SIZE |
| POST | /builds/{file} | Publish a New file, if the file exists replace it |
| PUT | /builds/{file} | Publish a New file, if the file exists replace it |
| Delete | /builds/{file} | Remove a existing file |
| GET | /builds" | List all the build files |
| GET | /libs" | Retrieve the Libs |
| POST | /libs" | Publish a new libs file, replace anything that already exists |
| PUT | /libs" | Publish a new libs file, replace anything that already exists |

