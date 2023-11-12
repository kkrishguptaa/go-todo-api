package types

type TodoWithoutID struct {
  Item      string `json:"title" bson:"title" binding:"required"`
  Completed bool   `json:"completed" bson:"completed" binding:"required"`
}

type Todo struct {
  ID        string `json:"_id" bson:"_id" binding:"required"`
  Item      string `json:"title" bson:"title" binding:"required"`
  Completed bool   `json:"completed" bson:"completed" binding:"required"`
}
