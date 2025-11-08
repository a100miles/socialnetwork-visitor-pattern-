package Visitors

import (
    "MiniSocial/Entities"
    "fmt"
)

type AnalyticsVisitor struct {
    TotalLikes int
    TotalViews int
}

func (a *AnalyticsVisitor) VisitUser(user *Entities.User) {
    for _, post := range user.Posts {
        post.Accept(a)
    }
}

func (a *AnalyticsVisitor) VisitPost(post *Entities.Post) {
    a.TotalLikes += post.Likes
}

func (a *AnalyticsVisitor) VisitStory(story *Entities.Story) {
    a.TotalViews += story.Views
}

func (a *AnalyticsVisitor) VisitComment(comment *Entities.Comment) {
    a.TotalLikes += comment.Likes
}

func (a *AnalyticsVisitor) Report() {
    fmt.Printf("Total Likes: %d | Total Views: %d\n", a.TotalLikes, a.TotalViews)
}
