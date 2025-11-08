package main

import (
    "MiniSocial/Entities"
    "MiniSocial/Visitors"
)

func main() {
    post1 := Entities.Post{Title: "Go Basics", Likes: 20, Tags: []string{"go", "programming"}, Content: "Learn Go step by step"}
    post2 := Entities.Post{Title: "Photography Tips", Likes: 50, Tags: []string{"photo", "art"}, Content: "Avoid bad lighting!"}
    comment := Entities.Comment{Text: "Great post!", Likes: 5}
    story := Entities.Story{Views: 100, Title: "Weekend Vibes"}

    user := Entities.User{
        Name:  "Tamerlan",
        Posts: []Entities.Post{post1, post2},
    }

    analytics := &Visitors.AnalyticsVisitor{}
    moderation := &Visitors.ModerationVisitor{ForbiddenWords: []string{"bad", "ugly"}}
    recommendation := &Visitors.RecommendationVisitor{TagToRecommend: "go"}

    entities := []Entities.IEntity{&user, &story, &comment, &post1, &post2}

    for _, entity := range entities {
        entity.Accept(analytics)
        entity.Accept(moderation)
        entity.Accept(recommendation)
    }

    analytics.Report()
}
