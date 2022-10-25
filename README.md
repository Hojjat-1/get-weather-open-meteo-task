# Get weather with Open Meteo task

An example of how create a task for DoTenX. (with golang)

# Notes
- You should haven't `main.go` file in your source code
- Your module name in `go.mod` file should be `main`
- You should haven't `main` function in all of your files

# An example of deploy command with dotenx cli
`
dotenx deploy -l go -f HandleLambdaEvent -p {YOUR_ABSOLUTE_PATH_OF_SOURCE_CODE_FILES} -d {YOUR_ABSOLUTE_PATH_OF_TASK_DEFINITION_YAML_FILE} -t task
`

# Deploy command options
- language, l 
    - language name ('go', 'node')
-	function, f 
    - function name
- path, p
    - the path to a directory where the source codes are stored
- definition_path, d
    - the path to a file where the definition of task/trigger stored (in yaml format)
- type, t
    - type of your function ('task' or 'trigger')
