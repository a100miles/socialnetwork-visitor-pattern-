package Visitors

import (
    "MiniSocial/Core"
    "MiniSocial/Entities"
    "fmt"
)

type AnalyticsVisitor struct {
    TotalLikes int
    TotalViews int
}

func (a *AnalyticsVisitor) VisitUser(user interface{}) {
    if u, ok := user.(*Entities.User); ok {
        for _, post := range u.Posts {
            post.Accept(a)
        }
    }
}

func (a *AnalyticsVisitor) VisitPost(post interface{}) {
    if p, ok := post.(*Entities.Post); ok {
        a.TotalLikes += p.Likes
    }
}

func (a *AnalyticsVisitor) VisitStory(story interface{}) {
    if s, ok := story.(*Entities.Story); ok {
        a.TotalViews += s.Views
    }
}

func (a *AnalyticsVisitor) VisitComment(comment interface{}) {
    if c, ok := comment.(*Entities.Comment); ok {
        a.TotalLikes += c.Likes
    }
}

func (a *AnalyticsVisitor) Report() {
    fmt.Printf("Total Likes: %d | Total Views: %d\n", a.TotalLikes, a.TotalViews)
}

var _ Core.IVisitor = (*AnalyticsVisitor)(nil)
