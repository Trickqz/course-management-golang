package models

type Curso struct {
    ID          int    `json:"id"`
    Nome        string `json:"nome"`
    Descricao   string `json:"descricao"`
}

type Usuario struct {
    ID    int    `json:"id" gorm:"primaryKey;autoIncrement"`
    Nome  string `json:"nome"`
    Email string `json:"email"`
}

type Progresso struct {
    UsuarioID int     `json:"usuario_id"`
    CursoID   int     `json:"curso_id"`
    Progresso float64  `json:"progresso"`
}
