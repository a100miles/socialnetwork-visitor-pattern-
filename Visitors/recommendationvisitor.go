package Visitors

import (
    "MiniSocial/Entities"
    "fmt"
)

type RecommendationVisitor struct {
    TagToRecommend string
}

func (r *RecommendationVisitor) VisitUser(user *Entities.User) {}
func (r *RecommendationVisitor) VisitStory(story *Entities.Story) {}
func (r *RecommendationVisitor) VisitComment(comment *Entities.Comment) {}

func (r *RecommendationVisitor) VisitPost(post *Entities.Post) {
    for _, tag := range post.Tags {
        if tag == r.TagToRecommend {
            fmt.Printf("[Recommendation] You might like post: \"%s\"\n", post.Title)
        }
    }
}
