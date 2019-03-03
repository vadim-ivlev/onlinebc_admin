package imgserver

import (
	"fmt"
	"io/ioutil"
	"onlinebc_admin/model/img"
	"os/exec"
	"path/filepath"
)

// bash("sshpass -p root ssh -p 222 root@localhost mkdir -p /var/www/onlinebc/uploads/foo; exit")
// bash("sshpass -p root scp -P 222 f.txt root@localhost:/var/www/onlinebc/uploads/foo/")

// MoveFileToImageServer перемещает файл  на удаленный сервер
func MoveFileToImageServer(filePath string) string {
	fileName := filepath.Base(filePath)
	destDir := params.Uploaddir + img.GenPhotoDir()

	cmdMkdir := fmt.Sprintf("sshpass -p %s ssh -q -o StrictHostKeyChecking=no -p %s %s@%s mkdir -p %s; exit", params.Password, params.Port, params.User, params.Host, destDir)
	cmdCopy := fmt.Sprintf("sshpass -p %s scp -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -P %s %s %s@%s:%s/", params.Password, params.Port, filePath, params.User, params.Host, destDir)
	cmdRemove := fmt.Sprintf("rm %s", filePath)

	bash(cmdMkdir)
	bash(cmdCopy)
	bash(cmdRemove)

	return img.GenPhotoDir() + "/" + fileName
}

func bash(cmdString string) (errMessage string) {
	cmd := exec.Command("bash")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	cmdWriter, _ := cmd.StdinPipe()
	errReader, _ := cmd.StderrPipe()
	cmd.Start()
	cmdWriter.Write([]byte(cmdString + "\n"))
	cmdWriter.Write([]byte("exit" + "\n"))
	errBytes, _ := ioutil.ReadAll(errReader)
	cmd.Wait()
	errMessage = string(errBytes)
	if errMessage != "" {
		fmt.Println(cmdString, "ERR:", errMessage)
	}
	return
}
