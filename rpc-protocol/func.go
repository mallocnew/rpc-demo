package protocol

// Arith 远程对象，必须首字母大写
type Arith struct{}

// Args 参数结构体，字段首字母也必须大写
type Args struct {
  A, B int
}

// Multiply 乘法方法，需要满足RPC规范
func (t *Arith) Multiply(args *Args, reply *int) error {
  *reply = args.A * args.B
  return nil
}
