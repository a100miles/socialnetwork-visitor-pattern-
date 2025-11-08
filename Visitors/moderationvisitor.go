package Visitors

import (
    "MiniSocial/Entities"
    "fmt"
    "strings"
)

type ModerationVisitor struct {
    ForbiddenWords []string
}

func (m *ModerationVisitor) VisitUser(user interface{}) {}
func (m *ModerationVisitor) VisitStory(story interface{}) {}

func (m *ModerationVisitor) VisitPost(post interface{}) {
    if p, ok := post.(*Entities.Post); ok {
        for _, word := range m.ForbiddenWords {
            if strings.Contains(strings.ToLower(p.Content), word) {
                fmt.Printf("[Moderation] Post \"%s\" contains forbidden word: %s\n", p.Title, word)
            }
        }
    }
}

func (m *ModerationVisitor) VisitComment(comment interface{}) {
    if c, ok := comment.(*Entities.Comment); ok {
        for _, word := range m.ForbiddenWords {
            if strings.Contains(strings.ToLower(c.Text), word) {
                fmt.Printf("[Moderation] Comment \"%s\" contains forbidden word: %s\n", c.Text, word)
            }
        }
    }
}
