# relpath

```sh
$ ./relpath.exe 
Usage: relpath.exe <path-to>
       relpath.exe <path-base> <path-to>

Description:
    'relpath.exe' calculates and outputs the relative path between specified paths.

Arguments:
    <path-to>      The destination path for which the relative path is calculated.
                   This path is output as relative from the current directory or from <path-base>.
    <path-base>    The base directory path (optional).
                   If specified, the relative path from this path to <path-to> is calculated.

Examples:
    1. Output the relative path from the current directory to a specified path:
       $ relpath.exe /home/user/docs

    2. Output the relative path from one directory to another:
       $ relpath.exe /var/log /var/log/syslog

Note:
    If the path does not exist or cannot be read, an error message will be displayed.

```

