# Windows Persistence Techniques in go
Do a partial actions of sharpersist, just rewritten in go so I can use it as a library in my own go-based payload.

## Installation 
```
go get github.com/idfp/go-persist
```
## API Docs
for myself, just in case i forgot how to exist.
### Functions
#### Function RegistryPersist
```go
func RegistryPersist(name string, path string) error
```
Create new registry key in `HKEY_CURRENT_USER\SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, the given path will be executed on windows logon.

E.g. 
```go
err := RegistryPersist("gpersist", "C:\\Windows\\System32\\cmd.exe")
```
#### Function StartupPersist
```go
func StartupPersist(name string, execPath string) error
```
Create a new windows Batch file `(name).bat` that will run `start (execPath)` which responsible for creating new process, and the path parameter is where ur exe exists. Then place the said batch file to `%APPDATA%\Microsoft\Windows\Start Menu\Programs\Startup\`

E.g. 
```go
err := StartupPersist("gpersist", "C:\\Windows\\System32\\cmd.exe")
```
#### Function SchTaskPersist
```go
func SchTaskPersist(name string, path string) error
```
A proxy function to call `schtasks` command in cmd, literally just call `schtasks` with certain arguments. The called command is more or less like this: 
```
schtasks /create /st 00:00 /tn "(name)" /tr (path)
```
which means "create a scheduled task that will run everyday at 00 AM, with name (name) and executes (path)"


E.g. 
```go
err := SchTaskPersist("gpersist", "C:\\Windows\\System32\\cmd.exe")
```