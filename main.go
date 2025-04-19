package main

import (
	"context"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/emersion/go-smtp"
)

type Backend struct{}

func (bkd *Backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	// Here you can implement your authentication logic
	// For this example, we will just log the connection
	return &Session{}, nil
}

func (bkd *Backend) AnonymousLogin(state *smtp.Conn) (smtp.Session, error) {
	log.Printf("Anonymous login from %s", state.Hostname())
	return &Session{}, nil
}

type Session struct {
	from string
	to   []string
}

// Ensure Session implements smtp.Session interface
var _ smtp.Session = &Session{}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	log.Println("Mail from:", from)
	s.from = from
	return nil
}

func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error {
	log.Println("Rcpt to:", to)
	s.to = append(s.to, to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	buf, err := io.ReadAll(r)
	if err != nil {
		log.Printf("Error reading message: %v", err)
		return err
	}
	log.Printf("Message: From=%s, To=%v, Data=%s", s.from, s.to, string(buf))
	return nil
}

func (s *Session) Reset() {
	s.from = ""
	s.to = nil
}

func (s *Session) Logout() error {
	return nil
}

func main() {
	be := &Backend{}
	s := smtp.NewServer(be)
	s.Addr = ":1025"
	s.Domain = "localhost"
	s.AllowInsecureAuth = true
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50

	log.Println("Starting SMTP server at", s.Addr)

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Setup graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server gracefully stopped")
}
