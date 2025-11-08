package Core

type IVisitor interface {
    VisitUser(user interface{})
    VisitPost(post interface{})
    VisitStory(story interface{})
    VisitComment(comment interface{})
}
