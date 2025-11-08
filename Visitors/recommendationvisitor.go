package Visitors

import (
    "MiniSocial/Entities"
    "fmt"
)

type RecommendationVisitor struct {
    TagToRecommend string
}

func (r *RecommendationVisitor) VisitUser(user interface{}) {}
func (r *RecommendationVisitor) VisitStory(story interface{}) {}
func (r *RecommendationVisitor) VisitComment(comment interface{}) {}

func (r *RecommendationVisitor) VisitPost(post interface{}) {
    if p, ok := post.(*Entities.Post); ok {
        for _, tag := range p.Tags {
            if tag == r.TagToRecommend {
                fmt.Printf("[Recommendation] You might like post: \"%s\"\n", p.Title)
            }
        }
    }
}
