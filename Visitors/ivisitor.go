package Visitors

import "MiniSocial/Entities"

type IVisitor interface {
    VisitUser(user *Entities.User)
    VisitPost(post *Entities.Post)
    VisitStory(story *Entities.Story)
    VisitComment(comment *Entities.Comment)
}
