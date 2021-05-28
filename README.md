# Building Go with Docker

## API
```
GET /status - HTTP/200
```
```
POST /
Body: 
{ "NAME": "YOUR NAME"}
```

## Testing

```
make test
```

## Building for Target

### Build Local (Match OS)
```
make
```

### Build for Target
```
make PLATFORM=windows/amd64
```
```
make PLATFORM=darwin/amd64
```
...

## Cleaning
```
make clean
```

## Running
Execute build target in `bin/...`