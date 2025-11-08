package main

import (
    "MiniSocial/Entities"
    "MiniSocial/Visitors"
)

func main() {
    post1 := Entities.Post{Title: "shit", Likes: 20, Tags: []string{"black", "programming"}, Content: "i hate ziggas"}
    post2 := Entities.Post{Title: "actual life", Likes: 50, Tags: []string{"photo", "art"}, Content: "1000-7"}
    comment := Entities.Comment{Text: "shit as always", Likes: 100}
    story := Entities.Story{Views: 100, Title: "chillin"}

    user := Entities.User{
        Name:  "a100ma",
        Posts: []Entities.Post{post1, post2},
    }

    analytics := &Visitors.AnalyticsVisitor{}
    moderation := &Visitors.ModerationVisitor{ForbiddenWords: []string{"nigga", "ugly"}}
    recommendation := &Visitors.RecommendationVisitor{TagToRecommend: "art"}

    entities := []Entities.IEntity{&user, &story, &comment, &post1, &post2}

    for _, entity := range entities {
        entity.Accept(analytics)
        entity.Accept(moderation)
        entity.Accept(recommendation)
    }

    analytics.Report()
}
