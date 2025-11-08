package Visitors

import (
    "MiniSocial/Entities"
    "fmt"
    "strings"
)

type ModerationVisitor struct {
    ForbiddenWords []string
}

func (m *ModerationVisitor) VisitUser(user *Entities.User) {}
func (m *ModerationVisitor) VisitStory(story *Entities.Story) {}

func (m *ModerationVisitor) VisitPost(post *Entities.Post) {
    for _, word := range m.ForbiddenWords {
        if strings.Contains(strings.ToLower(post.Content), word) {
            fmt.Printf("[Moderation] Post \"%s\" contains forbidden word: %s\n", post.Title, word)
        }
    }
}

func (m *ModerationVisitor) VisitComment(comment *Entities.Comment) {
    for _, word := range m.ForbiddenWords {
        if strings.Contains(strings.ToLower(comment.Text), word) {
            fmt.Printf("[Moderation] Comment \"%s\" contains forbidden word: %s\n", comment.Text, word)
        }
    }
}
