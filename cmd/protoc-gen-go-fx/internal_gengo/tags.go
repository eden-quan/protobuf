package internal_gengo

import (
	"regexp"
	"strings"

	"github.com/eden-quan/protobuf/compiler/protogen"
)

var reg = regexp.MustCompile(`#(\S+?):"(\S+?)"`)
var commentReg = regexp.MustCompile(`/+? `)

func extractMoreTags(tags structTags, field *protogen.Field) structTags {
	if field == nil || field.Comments.Trailing == "" {
		return tags
	}

	more := reg.FindAllStringSubmatch(field.Comments.Trailing.String(), -1)
	if len(more) > 0 {
		for _, m := range more {
			tags = append(tags, [2]string{m[1], m[2]})
		}
	}

	//newComment := reg.ReplaceAllLiteralString(field.Comments.Trailing.String(), "")
	newComment := field.Comments.Trailing.String()
	newComment = strings.TrimSpace(newComment)
	field.Comments.Trailing = protogen.Comments(newComment)
	return tags
}

func cleanupTags(comment string) string {
	newComment := reg.ReplaceAllLiteralString(comment, "")
	newComment = strings.TrimSpace(newComment)
	//newComment = strings.ReplaceAll(newComment, "////", "//")
	newComment = commentReg.ReplaceAllLiteralString(newComment, " ")

	return newComment
}
