# Repositories events 

this program analise files with repositories events in place (in memory)

## Usage 

To use program  build it with command 
```
go build 
```

then execute program ```./adjust``` with command line arguments

-path - Path to data directory. required - example [/home/example/]

-users - Users statistics. optional

-repos - Repositories statistics. optional

-top - Amount of users or repositories statistics to show. optional - max 100

-sort - Sort by events push|pull|watch|create|issue|delete|pull-push(only for users). optional

## Tests 

To run test use command ```go test``` in directory