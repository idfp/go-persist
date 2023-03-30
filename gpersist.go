package gpersist

import (
	"golang.org/x/sys/windows/registry"
	"os"
	"os/exec"
	"fmt"
)

func cmdOut(command string) (string, error) {
	cmd := exec.Command("cmd", "/C", command)
	output, err := cmd.CombinedOutput()
	out := string(output)
	return out, err
}

func RegistryPersist(name string, path string) error{
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, registry.SET_VALUE)
	if err != nil {
		return err
	}
	err2 := k.SetStringValue(name, path)
	if err2!= nil{
		return err2
	}
	defer k.Close()
	return nil
}

func StartupPersist(name string, execPath string) error{
	path := os.Getenv("APPDATA") + "\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\" + name + ".bat"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil{
		if os.IsNotExist(err){
			file, err = os.Create(path)
			if err != nil{
				return err
			}
		}else{
			return err
		}	
	}
	_, err2 := file.Write([]byte("start " + execPath))
	if err2 != nil{
		return err2
	}
	defer file.Close()
	return nil
}

func SchTaskPersist(name string, path string) error{
	_, err := cmdOut(fmt.Sprintf("schtasks /create /st 00:00 /tn %q /tr %s", name, path))
	return err
}