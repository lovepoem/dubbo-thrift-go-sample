// Autogenerated by Thrift Compiler (0.11.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
        "context"
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "git.apache.org/thrift.git/lib/go/thrift"
        "genthrift"
)


func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  bool echoBool(bool arg)")
  fmt.Fprintln(os.Stderr, "  i8 echoByte(i8 arg)")
  fmt.Fprintln(os.Stderr, "  i16 echoI16(i16 arg)")
  fmt.Fprintln(os.Stderr, "  i32 echoI32(i32 arg)")
  fmt.Fprintln(os.Stderr, "  i64 echoI64(i64 arg)")
  fmt.Fprintln(os.Stderr, "  double echoDouble(double arg)")
  fmt.Fprintln(os.Stderr, "  string echoString(string arg)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := genthrift.NewDemoServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "echoBool":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "EchoBool requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1) == "true"
    value0 := argvalue0
    fmt.Print(client.EchoBool(context.Background(), value0))
    fmt.Print("\n")
    break
  case "echoByte":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "EchoByte requires 1 args")
      flag.Usage()
    }
    tmp0, err17 := (strconv.Atoi(flag.Arg(1)))
    if err17 != nil {
      Usage()
      return
    }
    argvalue0 := int8(tmp0)
    value0 := argvalue0
    fmt.Print(client.EchoByte(context.Background(), value0))
    fmt.Print("\n")
    break
  case "echoI16":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "EchoI16 requires 1 args")
      flag.Usage()
    }
    tmp0, err18 := (strconv.Atoi(flag.Arg(1)))
    if err18 != nil {
      Usage()
      return
    }
    argvalue0 := int16(tmp0)
    value0 := argvalue0
    fmt.Print(client.EchoI16(context.Background(), value0))
    fmt.Print("\n")
    break
  case "echoI32":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "EchoI32 requires 1 args")
      flag.Usage()
    }
    tmp0, err19 := (strconv.Atoi(flag.Arg(1)))
    if err19 != nil {
      Usage()
      return
    }
    argvalue0 := int32(tmp0)
    value0 := argvalue0
    fmt.Print(client.EchoI32(context.Background(), value0))
    fmt.Print("\n")
    break
  case "echoI64":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "EchoI64 requires 1 args")
      flag.Usage()
    }
    argvalue0, err20 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err20 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.EchoI64(context.Background(), value0))
    fmt.Print("\n")
    break
  case "echoDouble":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "EchoDouble requires 1 args")
      flag.Usage()
    }
    argvalue0, err21 := (strconv.ParseFloat(flag.Arg(1), 64))
    if err21 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.EchoDouble(context.Background(), value0))
    fmt.Print("\n")
    break
  case "echoString":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "EchoString requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.EchoString(context.Background(), value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
