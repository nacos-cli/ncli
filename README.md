# ncli
ncli is a Nacos cli client.

## Installation
Download from [the release page](https://github.com/nacos-cli/ncli/releases), or you can build it by yourself via go build.

## Usage

### Subcommand

`ncli`

>  ncli is a nacos cli client  
> 
>Usage:  
>  
>ncli [command]  
>Available Commands:  
>  completion  Generate the autocompletion script for the specified shell  
>  config      manage config  
>  help        Help about any command  
>  namespace   manage namespace  
>  version     Show the current version  
> 
>Flags:
>-h, --help      help for ncli  
>-v, --verbose   verbose output  
>  
> Use "ncli [command] --help" for more information about a command.  

### Common Options

All the server info shorthand options are in Uppercase.  
All the other options are usually in lowercase and the first letter(s) of long options.  

>Server-related Flags:  
>-S, --schema string         nacos server schema (default "http")  
>-H, --host string           nacos server ip (default "127.0.0.1")  
>-P, --port uint16           nacos server port (default 8848)  
>-C, --context string        nacos server context path (default "/nacos")
>
>-u, --username string      nacos server auth username (default "nacos")  
>-p, --password string      nacos server auth password (default "nacos")  
>-n, --namespaceId string   nacos namespace id (default "public")  
>
>Global Flags:  
>-v, --verbose   verbose output

### Nacos Namespace

`ncli namespace -h`

>Manage Nacos namespace.  
>  
>Usage:  
>ncli namespace [command]  
>  
>Available Commands:  
>add         add Nacos namespace  
>exist       check namespace  
>  
 
>Flags:  
> check [common options](#common-options)  
>  
>Use "ncli namespace [command] --help" for more information about a command.  

#### Nacos Add Namespace

`ncli namespace add -h`

> Add Nacos namespace.  
>   
> Usage:  
> ncli namespace add [flags]  
>   
> Flags:  
> -D, --desc string   namespace description  
> -h, --help          help for add  
> -N, --name string   namespace name  
>   
 
> Global Flags:  
> check [common options](#common-options)  
 
#### Nacos Exist Namespace

`ncli namespace exist -h`

>   
> Check Nacos namespace for existence.  
>   
> Usage:  
> ncli namespace exist [flags]  
>   
> Flags:  
> -h, --help   help for exist  


> Global Flags:  
> check [common options](#common-options)  


### Nacos Config

`ncli config -h`

> Manage Nacos config.  
>   
> Usage:  
> ncli config [command]  
>   
> Available Commands:  
> add         add config  
> get         get config  
   
> Flags:  
> -g, --group string         config group (default "DEFAULT_GROUP")  
> -d, --dataId string        config data id  
>
> check [common options](#common-options)  
>   
> Use "ncli config [command] --help" for more information about a command.  

#### Nacos Add Config

`ncli config add -h`

>   Add Nacos config.  
>   
> Usage:  
> ncli config add [flags]  
>   
> Flags:  
> -a, --app string       optional, config application  
> -c, --content string   config content, optional if 'from' is specified  
> -D, --desc string      optional, config description (optional)  
> -f, --from string      post config from file, automatically override config content/type  
> -h, --help             help for add  
> -T, --tags string      optional, comma-delimited config tags  
> -t, --type string      config type, optional if 'from' is specified  

> Global Flags:  
> -g, --group string         config group (default "DEFAULT_GROUP")  
> -d, --dataId string        config data id  
>
> check [common options](#common-options)  

#### Nacos Get Config

`ncli config get -h`

> Get Nacos config.  
>   
> Usage:  
> ncli config get [flags]  
>   
> Flags:  
> -h, --help   help for get  

> Global Flags:  
> -g, --group string         config group (default "DEFAULT_GROUP")  
> -d, --dataId string        config data id  
>
> check [common options](#common-options)  


### Shell Completion

`ncli completion`

>  Generate the autocompletion script for ncli for the specified shell.  
>  See each sub-command's help for details on how to use the generated script.  
>    
>  Usage:  
>  ncli completion [command]  
>    
>  Available Commands:  
>  bash        Generate the autocompletion script for bash  
>  fish        Generate the autocompletion script for fish  
>  powershell  Generate the autocompletion script for powershell  
>  zsh         Generate the autocompletion script for zsh  
>    
>  Flags:  
>  -h, --help   help for completion  
>    
>  Use "ncli completion [command] --help" for more information about a command.  

#### bash

```
$ ncli completion bash > ~/.ncli_completion.bash
$ echo "source ${HOME}/.ncli_completion.bash" >> ~/.bashrc
```

#### fish

```
$ ncli completion fish > ~/.config/fish/completions/ncli.fish
```

#### zsh

```
$ ncli completion zsh > ~/.ncli_completion.zsh
$ echo "source ${HOME}/.ncli_completion.zsh" >> ~/.zshrc
```

#### powershell

```
PS > ncli completion powershell | Out-String | Invoke-Expression  
```

`ncli completion powershell -h`

> Generate the autocompletion script for powershell.  
>   
> To load completions in your current shell session:  
>   
>     ncli completion powershell | Out-String | Invoke-Expression  
>   
> To load completions for every new session, add the output of the above command  
> to your powershell profile.  
>   
> Usage:  
> ncli completion powershell [flags]  
>   
> Flags:  
> -h, --help              help for powershell  
> --no-descriptions   disable completion descriptions  
