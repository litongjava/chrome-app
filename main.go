package main

import (
  "os"
  "os/exec"
  "runtime"
  "syscall"
)

func main() {
  // 获取命令行参数，如果没有传递参数则默认打开 Google 首页
  url := "https://www.baidu.com"
  if len(os.Args) > 1 {
    url = os.Args[1]
  }

  switch runtime.GOOS {
  case "linux":
    exec.Command("xdg-open", "--app", "chrome", "--args", "--app="+url).Start() // Linux
  case "windows":
    cmd := exec.Command("cmd", "/C", "start", "chrome", "--app="+url) // Windows
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
    cmd.Start()
  case "darwin":
    exec.Command("open", "-a", "Google Chrome", "--args", "--app="+url).Start() // Mac OS
  default:
    panic("unsupported platform")
  }
}
