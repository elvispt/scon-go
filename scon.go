// Command line application that helps with making SSH connections
// Makes a few assumptions:
// That the keyfiles are inside you HOME_FOLDER/.keys/
package main

import (
  // for sending output to cli
  "fmt"
  // for getting cli arguments
  "os"
  // for finding ssh binary (the ssh app)
  "os/exec"
  // for finding user home directory
  "os/user"
  // for making system calls
  "syscall"
  // for dealing with errors
  "errors"
)

// Entry point for the app.
// Expects a single param on cli that is the server identifier
func main() {
  // output app title
  fmt.Println("--- scon: server connection application ---")

  var serverId string

  if len(os.Args) > 1 {
    serverId = os.Args[1]
  } else {
    serverId = ""
  }

  binary, lookErr := exec.LookPath("ssh")

  if lookErr != nil {
    panic(lookErr)
  }

  key, host, errorServer := getServerKeyAndHost(serverId)

  args := []string{"ssh", "-i", "", ""}
  if errorServer != nil {
    fmt.Println(errorServer);
    os.Exit(1)
  }
  args[2] = key
  args[3] = host

  env := os.Environ()
  execErr := syscall.Exec(binary, args, env)
  if execErr != nil {
    fmt.Println(execErr)
    os.Exit(1)
  }
}

// Based on provided serverId it returns the key and host for the server
func getServerKeyAndHost(serverId string) (string, string, error) {
  var key string
  var host string
  var errorServer error
  keysDir := "/.keys/"
  usr, err := user.Current()

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  keysDir = usr.HomeDir + keysDir
  switch serverId {
    case "eu":
      key = keysDir + "KEYFILE"
      host = "USER@HOST"
    case "sa":
      key = keysDir + "KEYFILE"
      host = "USER@HOST"
    default:
      errorServer = errors.New("Could not find server based on provided serverId")
  }

  return key, host, errorServer
}