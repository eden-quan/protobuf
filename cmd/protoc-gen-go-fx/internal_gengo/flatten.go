package internal_gengo

import (
	"github.com/sinsay/protobuf/compiler/protogen"
	"github.com/sinsay/protobuf/proto"
	"github.com/sinsay/protobuf/reflect/protoreflect"
	"github.com/sinsay/protobuf/types/descriptorpb"

	"github.com/sinsay/protobuf/cmd/protoc-gen-go-fx/meta"
)

type FlattenInfo struct {
	Rule    *meta.FlattenRules
	Message *messageInfo
	Field   *protogen.Field
}

func extractFlattenDesc(msg protoreflect.ProtoMessage) (bool, *meta.FlattenRules) {
	flatten := proto.GetExtension(msg, meta.E_Flatten).(bool)
	ext := proto.GetExtension(msg, meta.E_FlattenRule).(*meta.FlattenRules)
	return flatten, ext
}

func extractFlatten(message *messageInfo, field *protogen.Field) (*FlattenInfo, bool) {
	flatten, ext := extractFlattenDesc(field.Desc.Options())

	if ext == nil {
		// create default ext
		m := int32(0)
		ext = &meta.FlattenRules{Reserved: &meta.Reserved{
			Min: &m, // useless now
			Max: &m, // useless now
		}}
	} else {
		flatten = true
	}

	return &FlattenInfo{
		Rule:    ext,
		Message: message,
		Field:   field,
	}, flatten
}

func (fi *FlattenInfo) GetSetterMethods(_ *protogen.GeneratedFile, _ *fileInfo) {
	// won't support week type now
}

func (fi *FlattenInfo) GetGetterMethods(g *protogen.GeneratedFile, f *fileInfo) {
	flattenMessage, sf := fi.getSf()
	if sf == nil {
		return
	}

	m := fi.Message

	for _, field := range flattenMessage.Fields {
		// Getter for message field.
		goType, pointer := fieldGoType(g, f, field)
		defaultValue := fieldDefaultValue(g, f, m, field)
		g.Annotate(m.GoIdent.GoName+".Get"+field.GoName, field.Location)
		leadingComments := appendDeprecationSuffix("",
			field.Desc.ParentFile(),
			field.Desc.Options().(*descriptorpb.FieldOptions).GetDeprecated())
		switch {
		case field.Oneof != nil && !field.Oneof.Desc.IsSynthetic():
			g.P(leadingComments, "func (x *", m.GoIdent, ") Get", field.GoName, "() ", goType, " {")
			g.P("if x, ok := x.Get", field.Oneof.GoName, "().(*", field.GoIdent, "); ok {")
			g.P("return x.", field.GoName)
			g.P("}")
			g.P("return ", defaultValue)
			g.P("}")
		default:
			g.P(leadingComments, "func (x *", m.GoIdent, ") Get", field.GoName, "() ", goType, " {")
			if !field.Desc.HasPresence() || defaultValue == "nil" {
				g.P("if x != nil {")
			} else {
				g.P("if x != nil && x.", field.GoName, " != nil {")
			}
			star := ""
			if pointer {
				star = "*"
			}
			g.P("return ", star, " x.", field.GoName)
			g.P("}")
			g.P("return ", defaultValue)
			g.P("}")
		}
		g.P()
	}
}

func (fi *FlattenInfo) getSf() (*messageInfo, *structFields) {
	flattenFile := allFiles[fi.Field.Message.Location.SourceFile]
	ff := newFileInfo(flattenFile)
	flattenMessage := newMessageInfo(ff, fi.Field.Message)
	//index := file.allMessagesByPtr[f.Field.Message]
	//flattenMessage := file.allMessages[index]
	var sf *structFields
	for k, v := range ff.allMessageFieldsByPtr {
		if k.Message == fi.Field.Message {
			sf = v
			break
		}
	}
	if sf == nil {
		return nil, nil
	}

	return flattenMessage, sf
}

func (fi *FlattenInfo) GenMessage(g *protogen.GeneratedFile, file *fileInfo) {
	// Update index before generate
	flattenMessage, sf := fi.getSf()
	if sf == nil {
		return
	}

	for _, field := range flattenMessage.Fields {
		// update field index
		genMessageField(g, file, flattenMessage, field, sf)
	}
}