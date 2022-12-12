package models

type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"UsuarioId"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"UsuarioRelacionId"`
}
