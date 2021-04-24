package grace

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	status struct {
		started int32
		stoped  int32
	}

	stopChan chan os.Signal
	hooks    = make([]func(), 0)
	hookLock sync.Mutex

	logger = logrus.WithField("component", "grace")
)

func ProcessStart(name string, usage string, version string, serve func()) {
	app := &cli.App{
		Name:    name,
		Usage:   usage,
		Version: version,

		Action: func(c *cli.Context) error {
			if atomic.AddInt32(&status.started, 1) != 1 {
				return fmt.Errorf("server Process has be started: %d", status.started)
			}

			rand.Seed(time.Now().UnixNano())

			logger.Infof("server %s version:%s process start.", name, version)

			serve()

			stopChan = make(chan os.Signal, 1)
			signal.Ignore(syscall.SIGHUP)
			signal.Notify(stopChan,
				os.Interrupt,
				os.Kill,
				syscall.SIGALRM, // 时钟定时信号
				// syscall.SIGHUP,   // 终端连接断开
				syscall.SIGINT,  // Ctrl-C
				syscall.SIGTERM, // 结束程序
				// syscall.SIGQUIT,  // Ctrl-/
			)

			select {
			case sig, _ := <-stopChan:
				logger.Infof("<<<==================>>>")
				logger.Infof("<<<stop process by:%v>>>", sig)
				logger.Infof("<<<==================>>>")
				break
			}

			if atomic.LoadInt32(&status.started) == 0 || atomic.AddInt32(&status.stoped, 1) != 1 {
				return fmt.Errorf("server stop duplication")
			}

			logger.Infof("Starting execute server shutdown hooks")
			for _, hook := range hooks {
				hook()
			}

			logger.Infof("Server shutdown hooks executed completed")
			logger.Infof("server %s version:%s process shutdown finish.", name, version)
			return nil
		},
	}
	app.Run(os.Args)
}

// OnInterrupt 进程中断退出 处理
func OnInterrupt(fn func()) {
	hookLock.Lock()
	defer hookLock.Unlock()

	// control+c,etc
	// 控制终端关闭，守护进程不退出
	hooks = append(hooks, fn)
}
