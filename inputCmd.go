package bintb

import (
	"io"
	"os/exec"
	"syscall"

	"github.com/moisespsena-go/logging"
	"github.com/moisespsena-go/task"
)

// hora do GPS (GMT)
// temp do ar
// temp interna da caixa
// umidade do ar
// temp do ponto de orvalho
// pressao atm
// precipitacao (prp em 10min)
// prp (em 30min)
// prp em 1h
// prp em 24h
// lat
// long
// alt da estacao m
// potencia do sinal de radio

type InputCmd struct {
	Name   string
	Args   []string
	Env    []string
	Log    logging.Logger
	OnData io.Writer
}

func (this *InputCmd) Start(done func()) (stop task.Stoper, err error) {
	cmd := exec.Command(this.Name, this.Args...)
	if len(this.Env) > 0 {
		cmd.Env = this.Env
	}

	cmd.Stdout = this.OnData
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	errLog := logging.WithPrefix(this.Log, "ERR >")
	cmd.Stderr = IOWriter(func(data []byte) (n int, err error) {
		errLog.Error(string(data))
		return len(data), nil
	})

	if err = cmd.Start(); err != nil {
		return
	}

	return &InputCmdStoper{this, cmd}, nil
}

type InputCmdStoper struct {
	Input *InputCmd
	cmd   *exec.Cmd
}

func (this *InputCmdStoper) Stop() {
	if err := this.cmd.Process.Kill(); err != nil {
		this.Input.Log.Error(err.Error())
	} else {
		this.Input.Log.Info("done")
	}
}

func (this *InputCmdStoper) IsRunning() bool {
	return this.cmd.ProcessState == nil
}
