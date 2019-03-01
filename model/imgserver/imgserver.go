package imgserver

import (
	"fmt"
	"onlinebc_admin/model/img"
	"os"
	"os/exec"
	"path/filepath"
)

// sshpass -p root scp -P 222 TODO.md  root@localhost:/var/www/onlinebc/uploads/TODO.md
// rsync -r local-dir remote-machine:path
// ssh remote-host 'mkdir -p foo/bar/qux'
// ssh user@host "mkdir -p /target/path/" && scp /path/to/source user@host:/target/path/
// if err := exec.Command("go", "doc", "flag", "Flag.Name").Run(); err == nil {

///usr/bin/rsync -ratlz --rsh="/usr/bin/sshpass -p password ssh -o StrictHostKeyChecking=no -l username" src_path  dest_path

/// sshpass -p "password" rsync root@1.2.3.4:/abc /def
// Note the space at the start of the command, in the bash shell
// this will stop the command (and the password) from being stored in the history.

// rsync -rvz -e 'ssh -p 2222' --progress --remove-sent-files ./dir user@host:/path
// - just note the SSH command itself is enclosed in quotes.

// MoveFile перемещает файл  на удаленный сервер
func MoveFile(filePath string) string {
	fileName := filepath.Base(filePath)
	userhost := params.User + "@" + params.Host
	destDir := img.GenPhotoDir()
	destPath := img.GenPhotoPath(fileName)
	destination := params.User + "@" + params.Host + ":" + destPath

	println(fileName)
	println(userhost)
	println(destDir)
	println(destPath)
	println(destination)

	if err := createRemoteDir(userhost, destDir); err != nil {
		fmt.Println(err.Error())
		return ""
	}

	if err := copyFileToRemoteServer(filePath, destination); err != nil {
		fmt.Println(err.Error())
		return ""
	}

	if err := os.Remove(filePath); err != nil {
		fmt.Println(err.Error())
	}

	return destPath
}

func createRemoteDir(userhost string, dir string) error {
	return exec.Command("sshpass", "-p", params.Password, "ssh", "-p", params.Port, userhost, "'mkdir -p foo/bar/qux'").Run()
}

func copyFileToRemoteServer(filePath string, destination string) error {
	return exec.Command("sshpass", "-p", params.Password, "scp", "-P", params.Port, filePath, destination).Run()
}
