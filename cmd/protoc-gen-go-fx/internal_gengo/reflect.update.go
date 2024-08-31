package internal_gengo

import (
	"github.com/eden-quan/protobuf/compiler/protogen"
	"github.com/eden-quan/protobuf/internal/filedesc"
	"github.com/eden-quan/protobuf/proto"
	"github.com/eden-quan/protobuf/reflect/protoreflect"
	"github.com/eden-quan/protobuf/types/descriptorpb"
)

// processFlattenExtension 检查 descProto 中的消息类型，为每个消息调用 updateFlattenMessage 进行更新
func processFlattenExtension(gen *protogen.Plugin, descProto *descriptorpb.FileDescriptorProto, f *fileInfo) *descriptorpb.FileDescriptorProto {
	descProto = proto.Clone(descProto).(*descriptorpb.FileDescriptorProto)
	stripSourceRetentionFieldsFromMessage(descProto.ProtoReflect())

	for i, m := range descProto.MessageType {
		descProto.MessageType[i] = updateFlattenMessage(gen, m, i, descProto, f)
	}

	return descProto
}

// updateFlattenMessage 将 Flatten 字段添加到上级数据结构中
// 同时负责处理复杂对象，如果需要打平的字段中存在非字面量类型，则同时会将复杂对象的类型信息注册到元数据
func updateFlattenMessage(gen *protogen.Plugin, msg *descriptorpb.DescriptorProto, msgIndex int, descProto *descriptorpb.FileDescriptorProto, fileInfo *fileInfo) *descriptorpb.DescriptorProto {
	currFile := gen.FilesByPath[descProto.GetName()]
	currFileMsg := currFile.Messages[msgIndex]
	currFileMsgFields := make([]*protogen.Field, 0, len(currFileMsg.Fields))
	fields := make([]*descriptorpb.FieldDescriptorProto, 0, len(msg.Field))
	for i, f := range msg.GetField() {
		isFlatten, _ := extractFlattenDesc(f.Options)
		if !isFlatten {
			fields = append(fields, f)
			currFileMsgFields = append(currFileMsgFields, currFileMsg.Fields[i])
			continue
		}

		currField := currFileMsg.Fields[i]

		// find the message need flatten
		targetFile := gen.FilesByPath[currField.Message.Location.SourceFile]
		targetDescProto := interface{}(targetFile.Proto).(*descriptorpb.FileDescriptorProto)
		targetMsg := targetDescProto.MessageType[currField.Message.Desc.Index()]

		// recursive check
		targetMsg = updateFlattenMessage(gen, targetMsg, currField.Message.Desc.Index(), targetDescProto, fileInfo)

		for i, f := range targetMsg.GetField() {
			if currField.Message.Fields[i].Message != nil { // nested type
				newName := string(currFileMsg.Desc.FullName()) + "." + string(currField.Message.Fields[i].Message.Desc.Name())
				m1 := currField.Message.Fields[i].Message
				currField.Message.Fields[i].Message = m1
				if m2, ok := interface{}(m1.Desc).(*filedesc.Message); ok {
					m22 := *m2
					m22.L0.FullName = protoreflect.FullName(newName)
					m22.L1.IsMapEntry = true
					m1.Desc = &m22

					//checkMsg := fileInfo.File.Desc.Messages().(*filedesc.Messages)
					//checkMsg.List = append(checkMsg.List, *m2)

					// add extra message to parent
					//currFileMsg.Messages = append(currFileMsg.Messages, m1)
					//checkMsg2 := currFileMsg.Desc.Messages().(*filedesc.Messages)
					//checkMsg2.List[0].L1.IsMapEntry = true
					//checkMsg2.List = append(checkMsg.List, *m2)

					//fileInfo.Proto.MessageType = append(fileInfo.Proto.MessageType)

				}

				m2 := *m1
				m := newMessageInfo(newFileInfo(targetFile), &m2)
				newName = string(currFileMsg.Desc.Name()) + "." + string(m.Desc.Name())
				fileInfo.allMessages = append(fileInfo.allMessages, m)
				fileInfo.allMessagesByPtr[m] = len(fileInfo.allMessages) - 1
				fileInfo.allMessageFieldsByPtr[m] = new(structFields)
				//fileInfo.Messages = append(fileInfo.Messages, m.Message)
				fileInfo.File.Messages = append(fileInfo.File.Messages, m1)
				//fileInfo.Proto.MessageType = append(fileInfo.Proto.MessageType, )

				file := gen.FilesByPath[m.Message.Location.SourceFile].Proto
				// ger parent first
				index := currField.Message.Desc.Index()
				if currField.Message.Fields[i].Message != nil && currField.Message.Fields[i].Message.GoIdent.GoName == "Timestamp" {
					index = currField.Message.Fields[i].Message.Desc.Index()
					parentMsg := file.MessageType[index]
					msgProto := parentMsg
					//parentMsg.NestedType = append(parentMsg.NestedType, msgProto)
					descProto.MessageType = append(descProto.MessageType, msgProto)

				} else {
					parentMsg := file.MessageType[index]
					msgProto := parentMsg.NestedType[m.Message.Desc.Index()]
					// change name
					newName = string(currFileMsg.Desc.Name()) + "." + msgProto.GetName()
					msgProto = (proto.Clone(msgProto)).(*descriptorpb.DescriptorProto)
					msgProto.Name = &newName

					parentMsg.NestedType = append(parentMsg.NestedType, msgProto)
					descProto.MessageType = append(descProto.MessageType, msgProto)
				}

				//fileInfo.File.Proto.MessageType = append(fileInfo.File.Proto.MessageType, msgProto)
				//checkMsg.List = append(checkMsg.List, m.Message.Desc.)
				//if checkMsg.Len() == 1 {
				//	checkMsg = checkMsg
				//}
			}
			currFileMsgFields = append(currFileMsgFields, currField.Message.Fields[i])
			fields = append(fields, f)
		}
	}

	msg.Field = fields
	currFileMsg.Fields = currFileMsgFields
	return msg
}
