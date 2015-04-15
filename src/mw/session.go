package mw

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const SID string = "gosessionid"

var sessionMap map[string]*Session

func init() {
	sessionMap = make(map[string]*Session)
}

type SessionFunc interface {
	Set(k string, v interface{})
	Get(k string) interface{}
	Del(k string)
	ID() string
	Invalidate()
	Refresh()
}

type Session struct {
	sid          string
	maxlifetime  int64
	attrs        map[string]interface{}
	lastActivity int64
}

func (s *Session) Set(k string, v interface{}) {
	s.attrs[k] = v
}

func (s *Session) Refresh() {
	s.lastActivity = time.Now().Unix()
}

func (s *Session) Get(k string) (v interface{}) {
	v = s.attrs[k]
	return
}

func (s *Session) Del(k string) {
	delete(s.attrs, k)
}

func (s *Session) ID() string {
	return s.sid
}

func (s *Session) Invalidate() {
	delete(sessionMap, s.sid)
}

func NewSession(r *http.Request, w http.ResponseWriter) *Session {
	sid := genSessionId()
	s := &Session{sid: sid, attrs: make(map[string]interface{}), maxlifetime: 1800, lastActivity: time.Now().Unix()}
	sessionMap[sid] = s

	cookie := http.Cookie{Name: SID,
		Value:    url.QueryEscape(sid),
		Path:     "/",
		HttpOnly: true}

	http.SetCookie(w, &cookie)
	r.AddCookie(&cookie)
	return s
}

func GetSession(r *http.Request) *Session {
	cookie, errs := r.Cookie(SID)
	if errs != nil || cookie.Value == "" {
		return nil
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		log.Printf("request sid is %s", sid)
		session := sessionMap[sid]
		if session != nil {
			session.Refresh()
			return session
		} else {
			log.Printf("sid %s is expired", sid)
			return nil
		}
	}
}

func genSessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
