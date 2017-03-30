package log

import (
	"bufio"
	"fmt"
	"github.com/labstack/echo"
	"io"
	"os/exec"
)

func testLogPipeline() {
	cmd := exec.Command("ping www.baidu.com")
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	var reader = bufio.NewReader(stdout)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}

}

// Pineline is a api to return logs.
func Pineline(ctx echo.Context) error {
	logID := ctx.Get("id")
	fmt.Println(logID)
	go testLogPipeline()
	return nil
}
