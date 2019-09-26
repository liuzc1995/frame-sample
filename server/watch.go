package main

import (
	"bytes"
	"flag"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/542213314/frame-sample/config"
	"golanger.com/fsnotify"
	"golanger.com/log"
)

var (
	autoWatch   = flag.String("auto", "false", `when u modify some dir to auto restart this app`)
	deferToAuto = flag.Duration("defer", 1500*time.Millisecond, `defer to when u modify some file`)
	appName     = strings.TrimRight(filepath.Base(os.Args[0]), ".auto")
)

func copyApp() error {
	src := appName
	dst := appName + ".auto"
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	tmp, err := ioutil.TempFile(filepath.Dir(dst), "autofile") //创建临时目录
	if err != nil {
		return err
	}
	_, err = io.Copy(tmp, in)
	if err != nil {
		tmp.Close()
		os.Remove(tmp.Name())
		return err
	}
	if err := tmp.Close(); err != nil {
		os.Remove(tmp.Name())
		return err
	}
	const perm = 0744
	if err := os.Chmod(tmp.Name(), perm); err != nil {
		os.Remove(tmp.Name())
		return err
	}
	if err := os.Rename(tmp.Name(), dst); err != nil {
		os.Remove(tmp.Name())
		return err
	}

	return nil
}

func installApp() error {
	//go build
	cmd := exec.Command("go", "build")
	outp, err := cmd.CombinedOutput()
	log.Debug("Rebuild " + appName)
	log.Debug("Error: ", err)
	log.Debug("OutPut: " + string(outp))

	return err
}

func startApp() *exec.Cmd {
	cmd := exec.Command("./" + appName)
	var b bytes.Buffer
	//标准输出
	cmd.Stdout = &b
	//错误输出
	cmd.Stderr = &b
	err := cmd.Start()
	log.Debug("Error: ", err)
	log.Debug("OutPut: " + b.String())
	if err == nil {
		//保存进程，写入文件
		pid := cmd.Process.Pid
		ioutil.WriteFile(filepath.Clean(config.GetPid()), []byte(strconv.Itoa(pid)), 0700)
		log.Debug("Start "+appName+" - pid:", pid)
	}

	return cmd
}

func watchApp() {
	if *autoWatch == "true" {
		autoAppName := appName + ".auto"
		//copy server执行文件，生成server.auto执行文件
		if err := copyApp(); err != nil {
			log.Fatal(err)
		}
		//cmd执行 ./server.auto -auto watch
		//-auto 设置不为`false`的值 这里用watch表示为监控
		if err := exec.Command("./"+autoAppName, "-auto", "watch").Run(); err != nil {
			log.Fatal(err)
		}
	} else {
		//开启文件监听
		//open watch
		//添加监控对象
		watch, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}

		defer watch.Close()

		for _, dir := range config.GetWatchPath() {
			if err := watch.Add(dir); err != nil {
				log.Warn("watch dir:", dir, " error:", err)
			}
		}

		cmd := startApp()      //启动新进程
		execTime := time.Now() //执行时间

		for {
			select {
			case ev, ok := <-watch.Events:
				if ok {
					log.Debug("watch event:", ev)
					//监听创建、删除、重命名、写入
					if ev.Op&fsnotify.Chmod != fsnotify.Chmod {
						//设置延迟
						if time.Since(execTime) > *deferToAuto {
							time.AfterFunc((*deferToAuto + 2500*time.Millisecond), func() {
								if time.Since(execTime) > (*deferToAuto + 2500*time.Millisecond) {
									//go build
									if installApp() == nil {
										//kill process
										if err := cmd.Process.Kill(); err == nil {
											cmd = startApp()
										} else {
											log.Debug("cmd.Process.Kill error:", err)
										}
									}
								}
							})
							execTime = time.Now()
						}
					}
				}
			}
		}

		//Watch End

	}
}
