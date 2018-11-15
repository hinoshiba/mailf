mailf ( mail file )
===

* It is a command that can send mail from a file.
* Since the contents of stdin are transmitted, please pass the file with cat etc.

## Options
* -s
	* You can specify a mail server address. It is a required setting.
* -p
	* You can specify a mail server port.

## sample

```
[usershell]$ cat sample/sample.eml | mailf -s mta.example.com -p 25
```
