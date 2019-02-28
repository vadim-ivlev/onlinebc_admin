package ftp

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Save сохраняет байты bytes в  файл filename в директории dirname.
// Возвращает URI сохраненного файла или пустую строку в случае неудачи.
func Save(dirname string, filename string, bytes []byte) (uri string) {
	err := os.MkdirAll(dirname, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	path := dirname + "/" + filename
	err = ioutil.WriteFile(path, bytes, 0644)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return dirname + "/" + filename
}

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
