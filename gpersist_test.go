package gpersist

import (
	"testing"
	"golang.org/x/sys/windows/registry"
    "os"
)
const (
    NAME = "Gpersist"
    EXEC = "C:\\Windows\\System32\\cmd.exe"
)
func TestRegis(t *testing.T) {
    err := RegistryPersist(NAME, EXEC)
    if err != nil{
        t.Errorf("encountered an error when calling persist function, %s", err)
    }
	k, err2 := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, registry.QUERY_VALUE)
    if err2 != nil{
        t.Errorf("encountered an error when opening new key, %s", err2)
    }
    val, _, err3 := k.GetStringValue(NAME)
    if err3 != nil{
        t.Errorf("encountered an error when getting key's value, %s", err3)
    }
    if val != EXEC{
        t.Errorf("Wrong result, expected: %q, got: %q", EXEC, val)
    }
	key, _ := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, registry.SET_VALUE)
    key.DeleteValue(NAME)
    defer k.Close()
    defer key.Close()
}

func TestStartup(t *testing.T){
    err := StartupPersist(NAME, EXEC)
    if err != nil{
        t.Errorf("encountered an error when calling persist function, %s", err)
        return
    }
    path := os.Getenv("APPDATA") + "\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\" + NAME + ".bat"
	file, err2 := os.Open(path)
    if err2 != nil{
        t.Errorf("encountered an error when checking file's existence, %s", err2)
        return
    }
    data := make([]byte, 100)
    c, err3 := file.Read(data)
    if err3 != nil{
        t.Errorf("encountered an error when reading file, %s", err2)
        return
    }
    res := string(data[0:c])
    if res != "start " + EXEC{
        t.Errorf("Wrong result, expected: %q, got: %q", EXEC, res)
        return
    }
}