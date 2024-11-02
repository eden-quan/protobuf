package internal_gengo

import (
	"regexp"
	"strings"

	"gitlab.lainuoniao.cn/rhinobird/backend/protobuf.git/compiler/protogen"
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

			newTags := make(structTags, 0)
			// append 之前检查是否存在重复的，实现替换默认 tag 的能力
			for _, t := range tags {
				if t[0] == m[1] {
					continue
				}

				newTags = append(newTags, t)
			}

			tags = newTags
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
