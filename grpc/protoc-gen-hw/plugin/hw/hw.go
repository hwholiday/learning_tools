package hw

import (
	"fmt"
	pb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"path"
	"protoc-gen-hw/generator"
	"strconv"
	"strings"
)

const (
	msgPkgPath     = "git.huoys.com/kit/gomsg/pkg"
	contextPkgPath = "context"
	serverPkgPath  = "git.huoys.com/kit/gomsg/pkg/ws/server"
)

type Hw struct {
	gen *generator.Generator
}

func (g *Hw) Name() string {
	return "hw"
}

var (
	msgPkg     string
	contextPkg string
	serverPkg  string
	pkgImports map[generator.GoPackageName]bool
)

func (g *Hw) Init(gen *generator.Generator) {
	g.gen = gen
	msgPkg = generator.RegisterUniquePackageName("pkg", nil)
	contextPkg = generator.RegisterUniquePackageName("context", nil)
	serverPkg = generator.RegisterUniquePackageName("server", nil)
}

func (g *Hw) objectNamed(name string) generator.Object {
	g.gen.RecordTypeUse(name)
	return g.gen.ObjectNamed(name)
}

func (g *Hw) typeName(str string) string {
	return g.gen.TypeName(g.objectNamed(str))
}

func (g *Hw) P(args ...interface{}) { g.gen.P(args...) }

func (g *Hw) Generate(file *generator.FileDescriptor) {
	if len(file.FileDescriptorProto.Service) == 0 {
		return
	}
	g.P("// Reference imports to suppress errors if they are not otherwise used.")
	g.P("var _ ", msgPkg, ".IRet")
	g.P("var _ ", contextPkg, ".Context")
	g.P("var _ ", serverPkg, ".Server")
	g.P()

	for i, service := range file.FileDescriptorProto.Service {
		g.generateService(file, service, i)
	}
}

func (g *Hw) GenerateImports(file *generator.FileDescriptor, imports map[generator.GoImportPath]generator.GoPackageName) {
	if len(file.FileDescriptorProto.Service) == 0 {
		return
	}
	g.P("import (")
	g.P(msgPkg, " ", strconv.Quote(path.Join(g.gen.ImportPrefix, msgPkgPath)))
	g.P(contextPkg, " ", strconv.Quote(path.Join(g.gen.ImportPrefix, contextPkgPath)))
	g.P(serverPkg, " ", strconv.Quote(path.Join(g.gen.ImportPrefix, serverPkgPath)))
	g.P(")")
	g.P()
	pkgImports = make(map[generator.GoPackageName]bool)
	for _, name := range imports {
		pkgImports[name] = true
	}
}

func (g *Hw) generateService(file *generator.FileDescriptor, service *pb.ServiceDescriptorProto, index int) {
	path := fmt.Sprintf("6,%d", index) // 6 means service.
	origServName := service.GetName()
	/*serviceName := strings.ToLower(service.GetName())
	if pkg := file.GetPackage(); pkg != "" {
		serviceName = pkg
	}*/
	servName := generator.CamelCase(origServName)
	servAlias := servName + "Server"
	servWsAlias := servName + "WsServer"

	// strip suffix
	if strings.HasSuffix(servAlias, "ServerServer") {
		servAlias = strings.TrimSuffix(servAlias, "Server")
	}
	// Client interface.
	g.P("type ", servWsAlias, " interface {")
	for i, method := range service.Method {
		g.gen.PrintComments(fmt.Sprintf("%s,2,%d", path, i)) // 2 means method in a service.
		g.P("StartNotify(ws *server.Server)")
		g.P(g.generateClientSignature(method))
	}
	g.P("}")
	g.P()
	g.P("var srv ", servAlias)
	g.P("type shandler struct {")
	g.P("closeChan chan string")
	g.P("}")
	g.P()
	g.P("func (h *shandler) OnOpen(p pkg.Session) {}")
	g.P("func (h *shandler) OnClose(p pkg.Session, b bool) {")
	g.P("h.closeChan <- fmt.Sprint(p.ID())")
	g.P("}")
	g.P()
	g.P("func (h *shandler) OnReq(pk pkg.Session, data []byte) pkg.IRet {")
	g.P("var (\n\t\tpb  Message\n\t\tctx = context.WithValue(context.Background(), \"id\", pk.ID())\n\t)\n\terr := proto.Unmarshal(data, &pb)\n\tif err != nil {\n\t\treturn nil\n\t}\n\tswitch pb.Ops {")
	for i, method := range service.Method {
		g.gen.PrintComments(fmt.Sprintf("%s,2,%d", path, i))
		g.generateWsHandler(method)
	}
	g.P("default:\n\t\treturn pkg.Error(int16(pkg.NoHandler), \"not implemented\")\n\t}")
	g.P("}")
	g.P()
	g.P("func (h *shandler) OnPush(pk pkg.Session, data []byte) pkg.IRet {")
	g.P("return pkg.Ok(nil)")
	g.P("}")
	g.P()
	g.P("func RegisterDemoWsServer(e *server.Server, addr string, closeChan chan string, s ", servAlias, ") error {")
	g.P("srv = s")
	g.P("return e.ListenAndServe(addr, &shandler{closeChan: closeChan})")
	g.P("}")
	g.P()
}
func (g *Hw) generateClientSignature(method *pb.MethodDescriptorProto) string {
	origMethName := method.GetName()
	methName := generator.CamelCase(origMethName)
	reqArg := ", req *" + g.typeName(method.GetInputType())
	respName := "*" + g.typeName(method.GetOutputType())
	return fmt.Sprintf("%s(ctx %s.Context%s) (resp %s,err error)", methName, contextPkg, reqArg, respName)
}

func (g *Hw) generateWsHandler(method *pb.MethodDescriptorProto) {
	origMethName := method.GetName()
	methName := generator.CamelCase(origMethName)
	reqArg := g.typeName(method.GetInputType())
	g.P("case int32(GameCommand_", methName, "):")
	g.P("var req = &", reqArg, "{}")
	g.P("if err = proto.Unmarshal(pb.Data, req); err != nil {\n\t\t\treturn pkg.Error(int16(pkg.ReadErrorNo), err.Error())\n\t\t}")
	g.P("resp, err := srv.", methName, "(ctx, req)")
	g.P("if err != nil {return pkg.Error(int16(pkg.Write), err.Error())}")
	g.P("res, _ := proto.Marshal(resp)")
	g.P("return pkg.Ok(res)")
}
